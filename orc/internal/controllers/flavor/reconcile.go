/*
Copyright 2024 The ORC Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package flavor

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/go-logr/logr"
	"github.com/gophercloud/gophercloud/v2/openstack/compute/v2/flavors"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/utils/ptr"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"

	orcv1alpha1 "github.com/k-orc/openstack-resource-controller/api/v1alpha1"
	"github.com/k-orc/openstack-resource-controller/internal/controllers/common"
	"github.com/k-orc/openstack-resource-controller/internal/controllers/generic"
	osclients "github.com/k-orc/openstack-resource-controller/internal/osclients"
	"github.com/k-orc/openstack-resource-controller/internal/util/applyconfigs"
	orcerrors "github.com/k-orc/openstack-resource-controller/internal/util/errors"
	orcapplyconfigv1alpha1 "github.com/k-orc/openstack-resource-controller/pkg/clients/applyconfiguration/api/v1alpha1"
)

// +kubebuilder:rbac:groups=openstack.k-orc.cloud,resources=flavors,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=openstack.k-orc.cloud,resources=flavors/status,verbs=get;update;patch

func (r *orcFlavorReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	orcObject := &orcv1alpha1.Flavor{}
	err := r.client.Get(ctx, req.NamespacedName, orcObject)
	if err != nil {
		if apierrors.IsNotFound(err) {
			return ctrl.Result{}, nil
		}
		return ctrl.Result{}, err
	}

	if !orcObject.DeletionTimestamp.IsZero() {
		return r.reconcileDelete(ctx, orcObject)
	}

	return r.reconcileNormal(ctx, orcObject)
}

func (r *orcFlavorReconciler) getOpenStackClient(ctx context.Context, orcObject *orcv1alpha1.Flavor) (osclients.ComputeClient, error) {
	log := ctrl.LoggerFrom(ctx)

	clientScope, err := r.scopeFactory.NewClientScopeFromObject(ctx, r.client, log, orcObject)
	if err != nil {
		return nil, err
	}
	return clientScope.NewComputeClient()
}

func (r *orcFlavorReconciler) reconcileNormal(ctx context.Context, orcObject *orcv1alpha1.Flavor) (_ ctrl.Result, err error) {
	log := ctrl.LoggerFrom(ctx)
	log.V(3).Info("Reconciling resource")

	var statusOpts []updateStatusOpt
	addStatus := func(opt updateStatusOpt) {
		statusOpts = append(statusOpts, opt)
	}

	// Ensure we always update status
	defer func() {
		if err != nil {
			addStatus(withError(err))
		}

		err = errors.Join(err, r.updateStatus(ctx, orcObject, statusOpts...))

		var terminalError *orcerrors.TerminalError
		if errors.As(err, &terminalError) {
			log.Error(err, "not scheduling further reconciles for terminal error")
			err = nil
		}
	}()

	if !controllerutil.ContainsFinalizer(orcObject, Finalizer) {
		patch := common.SetFinalizerPatch(orcObject, Finalizer)
		return ctrl.Result{}, r.client.Patch(ctx, orcObject, patch, client.ForceOwnership, ssaFieldOwner(SSAFinalizerTxn))
	}

	osClient, err := r.getOpenStackClient(ctx, orcObject)
	if err != nil {
		return ctrl.Result{}, err
	}

	adapter := flavorActuator{
		Flavor:   orcObject,
		osClient: osClient,
	}

	osResource, err := generic.GetOSResourceFromObject(ctx, log, adapter)
	if err != nil {
		return ctrl.Result{}, err
	}

	if osResource == nil {
		log.V(3).Info("OpenStack resource does not yet exist")
		addStatus(withProgressMessage("Waiting for OpenStack resource to be created externally"))
		return ctrl.Result{RequeueAfter: externalUpdatePollingPeriod}, err
	}

	addStatus(withResource(osResource))

	if orcObject.Status.ID == nil {
		if err := r.setStatusID(ctx, orcObject, osResource.ID); err != nil {
			return ctrl.Result{}, err
		}
	}

	log = log.WithValues("ID", osResource.ID)
	log.V(4).Info("Got resource")
	ctx = ctrl.LoggerInto(ctx, log)

	if orcObject.Spec.ManagementPolicy == orcv1alpha1.ManagementPolicyManaged {
		for _, updateFunc := range needsUpdate(osClient, orcObject, osResource) {
			if err := updateFunc(ctx); err != nil {
				addStatus(withProgressMessage("Updating the OpenStack resource"))
				return ctrl.Result{}, fmt.Errorf("failed to update the OpenStack resource: %w", err)
			}
		}
	}

	return ctrl.Result{}, nil
}

func (r *orcFlavorReconciler) reconcileDelete(ctx context.Context, orcObject *orcv1alpha1.Flavor) (_ ctrl.Result, err error) {
	log := ctrl.LoggerFrom(ctx)
	log.V(3).Info("Reconciling OpenStack resource delete")

	var statusOpts []updateStatusOpt
	addStatus := func(opt updateStatusOpt) {
		statusOpts = append(statusOpts, opt)
	}

	deleted := false
	defer func() {
		// No point updating status after removing the finalizer
		if !deleted {
			if err != nil {
				addStatus(withError(err))
			}
			err = errors.Join(err, r.updateStatus(ctx, orcObject, statusOpts...))
		}
	}()

	// We won't delete the resource for an unmanaged object, or if onDelete is detach
	if orcObject.Spec.ManagementPolicy == orcv1alpha1.ManagementPolicyUnmanaged || orcObject.Spec.ManagedOptions.GetOnDelete() == orcv1alpha1.OnDeleteDetach {
		logPolicy := []any{"managementPolicy", orcObject.Spec.ManagementPolicy}
		if orcObject.Spec.ManagementPolicy == orcv1alpha1.ManagementPolicyManaged {
			logPolicy = append(logPolicy, "onDelete", orcObject.Spec.ManagedOptions.GetOnDelete())
		}
		log.V(4).Info("Not deleting OpenStack resource due to policy", logPolicy...)
	} else {
		deleted, requeue, err := r.deleteResource(ctx, log, orcObject, addStatus)
		if err != nil {
			return ctrl.Result{}, err
		}

		if !deleted {
			return ctrl.Result{RequeueAfter: requeue}, nil
		}
		log.V(4).Info("OpenStack resource is deleted")
	}

	deleted = true

	// Clear the finalizer
	applyConfig := orcapplyconfigv1alpha1.Flavor(orcObject.Name, orcObject.Namespace).WithUID(orcObject.UID)
	return ctrl.Result{}, r.client.Patch(ctx, orcObject, applyconfigs.Patch(types.ApplyPatchType, applyConfig), client.ForceOwnership, ssaFieldOwner(SSAFinalizerTxn))
}

func (r *orcFlavorReconciler) deleteResource(ctx context.Context, log logr.Logger, orcObject *orcv1alpha1.Flavor, addStatus func(updateStatusOpt)) (bool, time.Duration, error) {
	osClient, err := r.getOpenStackClient(ctx, orcObject)
	if err != nil {
		return false, 0, err
	}

	if orcObject.Status.ID != nil {
		// This GET is technically redundant because we could just check the
		// result from DELETE, but it's necessary if we want to report
		// status while deleting
		osResource, err := osClient.GetFlavor(ctx, *orcObject.Status.ID)

		switch {
		case orcerrors.IsNotFound(err):
			// Success!
			return true, 0, nil

		case err != nil:
			return false, 0, err

		default:
			addStatus(withResource(osResource))

			if len(orcObject.GetFinalizers()) > 1 {
				log.V(4).Info("Deferring resource cleanup due to remaining external finalizers")
				return false, 0, nil
			}

			if err := osClient.DeleteFlavor(ctx, *orcObject.Status.ID); err != nil {
				return false, 0, err
			}
			return false, deletePollingPeriod, nil
		}
	}

	// If status.ID is not set we need to check for an orphaned
	// resource. If we don't find one, assume success and continue,
	// otherwise set status.ID and let the controller delete by ID.

	osResource, err := GetByFilter(ctx, osClient, specToFilter(*orcObject.Spec.Resource))
	if err != nil {
		return false, 0, err
	}

	if osResource != nil {
		addStatus(withResource(osResource))
		return false, deletePollingPeriod, r.setStatusID(ctx, orcObject, osResource.ID)
	}

	// Didn't find an orphaned resource. Assume success.
	return true, 0, nil
}

// getResourceName returns the name of the OpenStack resource we should use.
func getResourceName(orcObject *orcv1alpha1.Flavor) orcv1alpha1.OpenStackName {
	if orcObject.Spec.Resource.Name != nil {
		return *orcObject.Spec.Resource.Name
	}
	return orcv1alpha1.OpenStackName(orcObject.Name)
}

// createResource creates an OpenStack resource for an ORC object.
func createResource(ctx context.Context, orcObject *orcv1alpha1.Flavor, osClient osclients.ComputeClient) (*flavors.Flavor, error) {
	if orcObject.Spec.ManagementPolicy == orcv1alpha1.ManagementPolicyUnmanaged {
		// Should have been caught by API validation
		return nil, orcerrors.Terminal(orcv1alpha1.OpenStackConditionReasonInvalidConfiguration, "Not creating unmanaged resource")
	}

	log := ctrl.LoggerFrom(ctx)
	log.V(3).Info("Creating OpenStack resource")

	resource := orcObject.Spec.Resource

	if resource == nil {
		// Should have been caught by API validation
		return nil, orcerrors.Terminal(orcv1alpha1.OpenStackConditionReasonInvalidConfiguration, "Creation requested, but spec.resource is not set")
	}

	createOpts := flavors.CreateOpts{
		Name:        string(getResourceName(orcObject)),
		RAM:         int(resource.RAM),
		VCPUs:       int(resource.Vcpus),
		Disk:        ptr.To(int(resource.Disk)),
		Swap:        ptr.To(int(resource.Swap)),
		IsPublic:    resource.IsPublic,
		Ephemeral:   ptr.To(int(resource.Ephemeral)),
		Description: string(ptr.Deref(resource.Description, "")),
	}

	osResource, err := osClient.CreateFlavor(ctx, createOpts)
	if err != nil {
		// We should require the spec to be updated before retrying a create which returned a conflict
		if orcerrors.IsConflict(err) {
			err = orcerrors.Terminal(orcv1alpha1.OpenStackConditionReasonInvalidConfiguration, "invalid configuration creating resource: "+err.Error(), err)
		}
		return nil, err
	}

	return osResource, nil
}

// needsUpdate returns a slice of functions that call the OpenStack API to
// align the OpenStack resoruce to its representation in the ORC spec object.
// Flavor does not support update yet.
func needsUpdate(_ osclients.ComputeClient, _ *orcv1alpha1.Flavor, _ *flavors.Flavor) (updateFuncs []func(context.Context) error) {
	return nil
}
