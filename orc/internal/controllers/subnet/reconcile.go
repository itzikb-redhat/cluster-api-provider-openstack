/*
Copyright 2024 The Kubernetes Authors.

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

package subnet

import (
	"context"
	"errors"
	"fmt"
	"slices"
	"time"

	"github.com/go-logr/logr"
	"github.com/gophercloud/gophercloud/v2/openstack/networking/v2/subnets"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/utils/ptr"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"

	orcv1alpha1 "github.com/k-orc/openstack-resource-controller/api/v1alpha1"
	osclients "github.com/k-orc/openstack-resource-controller/internal/osclients"
	orcerrors "github.com/k-orc/openstack-resource-controller/internal/util/errors"
	"github.com/k-orc/openstack-resource-controller/internal/util/neutrontags"
	"github.com/k-orc/openstack-resource-controller/internal/util/ssa"
	orcapplyconfigv1alpha1 "github.com/k-orc/openstack-resource-controller/pkg/clients/applyconfiguration/api/v1alpha1"
)

// +kubebuilder:rbac:groups=openstack.k-orc.cloud,resources=subnets,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=openstack.k-orc.cloud,resources=subnets/status,verbs=get;update;patch

func (r *orcSubnetReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	orcObject := &orcv1alpha1.Subnet{}
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

func (r *orcSubnetReconciler) getNetworkClient(ctx context.Context, orcSubnet *orcv1alpha1.Subnet) (osclients.NetworkClient, error) {
	log := ctrl.LoggerFrom(ctx)

	clientScope, err := r.scopeFactory.NewClientScopeFromObject(ctx, r.client, r.caCertificates, log, orcSubnet)
	if err != nil {
		return nil, err
	}
	return clientScope.NewNetworkClient()
}

func (r *orcSubnetReconciler) reconcileNormal(ctx context.Context, orcObject *orcv1alpha1.Subnet) (_ ctrl.Result, err error) {
	log := ctrl.LoggerFrom(ctx)
	log.V(3).Info("Reconciling subnet")

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

	orcNetwork := &orcv1alpha1.Network{}
	if err := r.client.Get(ctx, client.ObjectKey{Name: string(orcObject.Spec.NetworkRef), Namespace: orcObject.Namespace}, orcNetwork); err != nil {
		if apierrors.IsNotFound(err) {
			addStatus(withProgressMessage(fmt.Sprintf("waiting for network object %s to be created", orcObject.Spec.NetworkRef)))
			return ctrl.Result{}, nil
		}
		return ctrl.Result{}, err
	}

	if !orcv1alpha1.IsAvailable(orcNetwork) {
		addStatus(withProgressMessage(fmt.Sprintf("waiting for network object %s to be available", orcObject.Spec.NetworkRef)))
		return ctrl.Result{}, nil
	}

	if orcNetwork.Status.ID == nil {
		return ctrl.Result{}, fmt.Errorf("network %s is available but status.ID is not set", orcNetwork.Name)
	}
	networkID := orcv1alpha1.UUID(*orcNetwork.Status.ID)

	if orcObject.Status.NetworkID == nil {
		return ctrl.Result{}, r.setStatusNetworkID(ctx, orcObject, networkID)
	}

	if *orcObject.Status.NetworkID != networkID {
		return ctrl.Result{}, orcerrors.Terminal(orcv1alpha1.OpenStackConditionReasonUnrecoverableError, "Parent network ID has changed")
	}

	// Don't add finalizer until parent network is available to avoid unnecessary reconcile on delete
	if !controllerutil.ContainsFinalizer(orcObject, Finalizer) {
		return ctrl.Result{}, r.setFinalizer(ctx, orcObject)
	}

	networkClient, err := r.getNetworkClient(ctx, orcObject)
	if err != nil {
		return ctrl.Result{}, err
	}

	osResource, waitingOnExternal, err := getOSResourceFromObject(ctx, log, orcObject, networkID, networkClient)
	if err != nil {
		return ctrl.Result{}, err
	}
	if waitingOnExternal {
		log.V(3).Info("OpenStack resource does not yet exist")
		addStatus(withProgressMessage("Waiting for OpenStack resource to be created externally"))
		return ctrl.Result{RequeueAfter: externalUpdatePollingPeriod}, err
	}

	if osResource == nil {
		if orcObject.Spec.ManagementPolicy == orcv1alpha1.ManagementPolicyManaged {
			osResource, err = createResource(ctx, orcObject, networkID, networkClient)
			if err != nil {
				return ctrl.Result{}, nil
			}
		} else {
			// Programming error
			return ctrl.Result{}, fmt.Errorf("unmanaged object does not exist and not waiting on dependency")
		}
	}

	addStatus(withResource(osResource))

	if orcObject.Status.ID == nil {
		if err := r.setStatusID(ctx, orcObject, osResource.ID); err != nil {
			return ctrl.Result{}, err
		}
	}

	log = log.WithValues("ID", osResource.ID)
	log.V(4).Info("Got subnet")
	ctx = ctrl.LoggerInto(ctx, log)

	return ctrl.Result{}, nil
}

func getOSResourceFromObject(ctx context.Context, log logr.Logger, orcObject *orcv1alpha1.Subnet, networkID orcv1alpha1.UUID, networkClient osclients.NetworkClient) (*subnets.Subnet, bool, error) {
	switch {
	case orcObject.Status.ID != nil:
		log.V(4).Info("Fetching existing OpenStack resource", "ID", *orcObject.Status.ID)
		osResource, err := networkClient.GetSubnet(*orcObject.Status.ID)
		if err != nil {
			if orcerrors.IsNotFound(err) {
				// An OpenStack resource we previously referenced has been deleted unexpectedly. We can't recover from this.
				return nil, false, orcerrors.Terminal(orcv1alpha1.OpenStackConditionReasonUnrecoverableError, "resource has been deleted from OpenStack")
			}
			return nil, false, err
		}
		return osResource, false, nil

	case orcObject.Spec.Import != nil && orcObject.Spec.Import.ID != nil:
		log.V(4).Info("Importing existing OpenStack resource by ID")
		osResource, err := networkClient.GetSubnet(*orcObject.Spec.Import.ID)
		if err != nil {
			if orcerrors.IsNotFound(err) {
				// We assume that an image imported by ID must already exist. It's a terminal error if it doesn't.
				return nil, false, orcerrors.Terminal(orcv1alpha1.OpenStackConditionReasonUnrecoverableError, "referenced resource does not exist in OpenStack")
			}
			return nil, false, err
		}
		return osResource, false, nil

	case orcObject.Spec.Import != nil && orcObject.Spec.Import.Filter != nil:
		log.V(4).Info("Importing existing OpenStack resource by filter")
		listOpts := listOptsFromImportFilter(orcObject.Spec.Import.Filter, networkID)
		osResource, err := getResourceFromList(ctx, listOpts, networkClient)
		if err != nil {
			return nil, false, err
		}
		if osResource == nil {
			return nil, true, nil
		}
		return osResource, false, nil

	default:
		log.V(4).Info("Checking for previously created OpenStack resource")
		listOpts := listOptsFromCreation(orcObject)
		osResource, err := getResourceFromList(ctx, listOpts, networkClient)
		if err != nil {
			return nil, false, nil
		}
		return osResource, false, nil
	}
}

func (r *orcSubnetReconciler) reconcileDelete(ctx context.Context, orcObject *orcv1alpha1.Subnet) (_ ctrl.Result, err error) {
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
		networkClient, err := r.getNetworkClient(ctx, orcObject)
		if err != nil {
			return ctrl.Result{}, err
		}

		// XXX: Don't do this. getOSResourceFromObject refactor was probably unnecessary. Consider undoing it for simplicity. Here we should:
		// * fetch by ID if set
		// * fetch by creation opts if ID is not set
		osResource, _, err := getOSResourceFromObject(ctx, log, orcObject, "", networkClient)
		if err != nil && !orcerrors.IsNotFound(err) {
			return ctrl.Result{}, err
		}
		addStatus(withResource(osResource))

		// Delete any returned OpenStack resource, but don't clear the finalizer until a fetch returns nothing
		if osResource != nil {
			log.V(4).Info("Deleting OpenStack resource", "id", osResource.ID)
			err := networkClient.DeleteSubnet(osResource.ID)
			if err != nil {
				return ctrl.Result{}, err
			}
			return ctrl.Result{RequeueAfter: 5 * time.Second}, nil
		}

		log.V(4).Info("OpenStack resource is deleted")
	}

	deleted = true

	// Clear the finalizer
	applyConfig := orcapplyconfigv1alpha1.Image(orcObject.Name, orcObject.Namespace).WithUID(orcObject.UID)
	return ctrl.Result{}, r.client.Patch(ctx, orcObject, ssa.ApplyConfigPatch(applyConfig), client.ForceOwnership, ssaFieldOwner(SSAFinalizerTxn))
}

// getResourceName returns the name of the OpenStack resource we should use.
func getResourceName(orcObject *orcv1alpha1.Subnet) orcv1alpha1.OpenStackName {
	if orcObject.Spec.Resource.Name != nil {
		return *orcObject.Spec.Resource.Name
	}
	return orcv1alpha1.OpenStackName(orcObject.Name)
}

func listOptsFromImportFilter(filter *orcv1alpha1.SubnetFilter, networkID orcv1alpha1.UUID) subnets.ListOptsBuilder {
	listOpts := subnets.ListOpts{
		Name:        string(ptr.Deref(filter.Name, "")),
		Description: string(ptr.Deref(filter.Description, "")),
		NetworkID:   string(networkID),
		IPVersion:   int(ptr.Deref(filter.IPVersion, 0)),
		GatewayIP:   string(ptr.Deref(filter.GatewayIP, "")),
		CIDR:        string(ptr.Deref(filter.CIDR, "")),
		Tags:        neutrontags.Join(filter.FilterByNeutronTags.Tags),
		TagsAny:     neutrontags.Join(filter.FilterByNeutronTags.TagsAny),
		NotTags:     neutrontags.Join(filter.FilterByNeutronTags.NotTags),
		NotTagsAny:  neutrontags.Join(filter.FilterByNeutronTags.NotTagsAny),
	}
	if filter.IPv6 != nil {
		listOpts.IPv6AddressMode = string(ptr.Deref(filter.IPv6.AddressMode, ""))
		listOpts.IPv6RAMode = string(ptr.Deref(filter.IPv6.RAMode, ""))
	}

	return &listOpts
}

// listOptsFromCreation returns a listOpts which will return the OpenStack
// resource which would have been created from the current spec and hopefully no
// other. Its purpose is to automatically adopt a resource that we created but
// failed to write to status.id.
func listOptsFromCreation(osResource *orcv1alpha1.Subnet) subnets.ListOptsBuilder {
	return subnets.ListOpts{Name: string(getResourceName(osResource))}
}

func getResourceFromList(_ context.Context, listOpts subnets.ListOptsBuilder, networkClient osclients.NetworkClient) (*subnets.Subnet, error) {
	osResources, err := networkClient.ListSubnet(listOpts)
	if err != nil {
		return nil, err
	}

	if len(osResources) == 1 {
		return &osResources[0], nil
	}

	// No image found
	if len(osResources) == 0 {
		return nil, nil
	}

	// Multiple resources found
	return nil, orcerrors.Terminal(orcv1alpha1.OpenStackConditionReasonInvalidConfiguration, fmt.Sprintf("Expected to find exactly one OpenStack resource to import. Found %d", len(osResources)))
}

// createResource creates a Glance image for an ORC Image.
func createResource(ctx context.Context, orcObject *orcv1alpha1.Subnet, networkID orcv1alpha1.UUID, networkClient osclients.NetworkClient) (*subnets.Subnet, error) {
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

	// TODO: Write this
	createOpts := subnets.CreateOpts{}

	tags := make([]string, len(resource.Tags))
	for i := range resource.Tags {
		tags[i] = string(resource.Tags[i])
	}
	// Sort tags before creation to simplify comparisons
	slices.Sort(tags)

	osResource, err := networkClient.CreateSubnet(&createOpts)

	// We should require the spec to be updated before retrying a create which returned a conflict
	if orcerrors.IsConflict(err) {
		err = orcerrors.Terminal(orcv1alpha1.OpenStackConditionReasonInvalidConfiguration, "invalid configuration creating image: "+err.Error(), err)
	}

	return osResource, err
}
