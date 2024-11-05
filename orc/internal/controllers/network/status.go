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

package network

import (
	"context"
	"encoding/json"
	"errors"
	"strconv"
	"time"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	applyconfigv1 "k8s.io/client-go/applyconfigurations/meta/v1"
	"k8s.io/utils/ptr"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	orcv1alpha1 "github.com/k-orc/openstack-resource-controller/api/v1alpha1"
	orcerrors "github.com/k-orc/openstack-resource-controller/internal/util/errors"
	"github.com/k-orc/openstack-resource-controller/internal/util/ssa"
	orcapplyconfigv1alpha1 "github.com/k-orc/openstack-resource-controller/pkg/clients/applyconfiguration/api/v1alpha1"
)

// setFinalizer sets a finalizer on the object in its own SSA transaction.
func (r *orcNetworkReconciler) setFinalizer(ctx context.Context, obj client.Object) error {
	gvk := obj.GetObjectKind().GroupVersionKind()

	applyConfig := struct {
		applyconfigv1.TypeMetaApplyConfiguration   `json:",inline"`
		applyconfigv1.ObjectMetaApplyConfiguration `json:"metadata,omitempty"`
	}{}

	// Type meta
	applyConfig.
		WithAPIVersion(gvk.GroupVersion().String()).
		WithKind(gvk.Kind)

	// Object meta
	applyConfig.
		WithName(obj.GetName()).
		WithNamespace(obj.GetNamespace()).
		WithUID(obj.GetUID()). // For safety: ensure we don't accidentally create a new object if we race with delete
		WithFinalizers(Finalizer)

	return r.client.Patch(ctx, obj, ssa.ApplyConfigPatch(applyConfig), client.ForceOwnership, ssaFieldOwner(SSAFinalizerTxn))
}

// setStatusID sets status.ID in its own SSA transaction.
func (r *orcNetworkReconciler) setStatusID(ctx context.Context, obj client.Object, id string) error {
	applyConfig := orcapplyconfigv1alpha1.Subnet(obj.GetName(), obj.GetNamespace()).
		WithUID(obj.GetUID()).
		WithStatus(orcapplyconfigv1alpha1.SubnetStatus().
			WithID(id))

	return r.client.Status().Patch(ctx, obj, ssa.ApplyConfigPatch(applyConfig), client.ForceOwnership, ssaFieldOwner(SSAIDTxn))
}

type updateStatusOpts struct {
	resource        *networkExt
	progressMessage *string
	err             error
}

type updateStatusOpt func(*updateStatusOpts)

func withResource(resource *networkExt) updateStatusOpt {
	return func(opts *updateStatusOpts) {
		opts.resource = resource
	}
}

func withError(err error) updateStatusOpt {
	return func(opts *updateStatusOpts) {
		opts.err = err
	}
}

// withProgressMessage sets a custom progressing message if and only if the reconcile is progressing.
func withProgressMessage(message string) updateStatusOpt {
	return func(opts *updateStatusOpts) {
		opts.progressMessage = &message
	}
}

func getOSResourceStatus(log logr.Logger, osResource *networkExt) *orcapplyconfigv1alpha1.NetworkResourceStatusApplyConfiguration {
	networkResourceStatus := (&orcapplyconfigv1alpha1.NetworkResourceStatusApplyConfiguration{}).
		WithName(osResource.Name).
		WithDescription(osResource.Description).
		WithAdminStateUp(osResource.AdminStateUp).
		WithAvailabilityZoneHints(osResource.AvailabilityZoneHints...).
		WithStatus(osResource.Status).
		WithProjectID(osResource.ProjectID).
		WithTags(osResource.Tags...).
		WithDNSDomain(osResource.DNSDomain).
		WithRevisionNumber(orcv1alpha1.NeutronRevisionNumber(osResource.RevisionNumber)).
		WithExternal(osResource.External).
		WithSubnets(osResource.Subnets...).
		WithMTU(int32(osResource.MTU)).
		WithPortSecurityEnabled(osResource.PortSecurityEnabled).
		WithShared(osResource.Shared).
		WithCreatedAt(metav1.NewTime(osResource.CreatedAt)).
		WithUpdatedAt(metav1.NewTime(osResource.UpdatedAt))

	if osResource.NetworkType != "" {
		providerProperties := orcapplyconfigv1alpha1.ProviderProperties().
			WithNetworkType(orcv1alpha1.ProviderNetworkType(osResource.NetworkType)).
			WithPhysicalNetwork(orcv1alpha1.PhysicalNetwork(osResource.PhysicalNetwork))

		if osResource.SegmentationID != "" {
			segmentationID, err := strconv.ParseInt(osResource.SegmentationID, 10, 32)
			if err != nil {
				log.V(3).Error(err, "Invalid segmentation ID", "segmentationID", osResource.SegmentationID)
			} else {
				providerProperties.WithSegmentationID(int32(segmentationID))
			}
		}
		networkResourceStatus.WithProvider(providerProperties)
	}

	return networkResourceStatus
}

const NetworkStatusActive = "ACTIVE"

// createStatusUpdate computes a complete status update based on the given
// observed state. This is separated from updateStatus to facilitate unit
// testing, as the version of k8s we currently import does not support patch
// apply in the fake client.
// Needs: https://github.com/kubernetes/kubernetes/pull/125560
func createStatusUpdate(ctx context.Context, orcNetwork *orcv1alpha1.Network, now metav1.Time, opts ...updateStatusOpt) *orcapplyconfigv1alpha1.NetworkApplyConfiguration {
	log := ctrl.LoggerFrom(ctx)

	statusOpts := updateStatusOpts{}
	for i := range opts {
		opts[i](&statusOpts)
	}

	osResource := statusOpts.resource
	err := statusOpts.err

	applyConfigStatus := orcapplyconfigv1alpha1.NetworkStatus()
	applyConfig := orcapplyconfigv1alpha1.Network(orcNetwork.Name, orcNetwork.Namespace).WithStatus(applyConfigStatus)

	if osResource != nil {
		resourceStatus := getOSResourceStatus(log, osResource)
		applyConfigStatus.WithResource(resourceStatus)
	}

	availableCondition := applyconfigv1.Condition().
		WithType(orcv1alpha1.OpenStackConditionAvailable).
		WithObservedGeneration(orcNetwork.Generation)
	progressingCondition := applyconfigv1.Condition().
		WithType(orcv1alpha1.OpenStackConditionProgressing).
		WithObservedGeneration(orcNetwork.Generation)

	// A network is available when it has a status of ACTIVE
	available := osResource != nil && osResource.Status == NetworkStatusActive
	if available {
		availableCondition.
			WithStatus(metav1.ConditionTrue).
			WithReason(orcv1alpha1.OpenStackConditionReasonSuccess).
			WithMessage("OpenStack resource is available")
	} else {
		// Image is not available. Reason and message will be copied from Progressing
		availableCondition.WithStatus(metav1.ConditionFalse)
	}

	// We are progressing until the OpenStack resource is available or there was an error
	if err == nil {
		if available {
			progressingCondition.
				WithStatus(metav1.ConditionFalse).
				WithReason(orcv1alpha1.OpenStackConditionReasonSuccess).
				WithMessage(*availableCondition.Message)
		} else {
			progressingCondition.
				WithStatus(metav1.ConditionTrue).
				WithReason(orcv1alpha1.OpenStackConditionReasonProgressing)

			if statusOpts.progressMessage == nil {
				progressingCondition.WithMessage("Reconciliation is progressing")
			} else {
				progressingCondition.WithMessage(*statusOpts.progressMessage)
			}
		}
	} else {
		progressingCondition.WithStatus(metav1.ConditionFalse)

		var terminalError *orcerrors.TerminalError
		if errors.As(err, &terminalError) {
			progressingCondition.
				WithReason(terminalError.Reason).
				WithMessage(terminalError.Message)
		} else {
			progressingCondition.
				WithReason(orcv1alpha1.OpenStackConditionReasonTransientError).
				WithMessage(err.Error())
		}
	}

	// Copy available status from progressing if it's not available yet
	if !available {
		availableCondition.
			WithReason(*progressingCondition.Reason).
			WithMessage(*progressingCondition.Message)
	}

	// Maintain condition timestamps if they haven't changed
	// This also ensures that we don't generate an update event if nothing has changed
	for _, condition := range []*applyconfigv1.ConditionApplyConfiguration{availableCondition, progressingCondition} {
		previous := meta.FindStatusCondition(orcNetwork.Status.Conditions, *condition.Type)
		if previous != nil && ssa.ConditionsEqual(previous, condition) {
			condition.WithLastTransitionTime(previous.LastTransitionTime)
		} else {
			condition.WithLastTransitionTime(now)
		}
	}

	applyConfigStatus.WithConditions(availableCondition, progressingCondition)

	if log.V(4).Enabled() {
		logValues := make([]any, 0, 12)
		addConditionValues := func(condition *applyconfigv1.ConditionApplyConfiguration) {
			if condition.Type == nil {
				bytes, _ := json.Marshal(condition)
				log.V(0).Info("Attempting to set condition with no type", "condition", string(bytes))
				return
			}

			for _, v := range []struct {
				name  string
				value *string
			}{
				{"status", (*string)(condition.Status)},
				{"reason", condition.Reason},
				{"message", condition.Message},
			} {
				logValues = append(logValues, *condition.Type+"."+v.name, ptr.Deref(v.value, ""))
			}
		}
		addConditionValues(availableCondition)
		addConditionValues(progressingCondition)
		log.V(4).Info("Setting image status", logValues...)
	}

	return applyConfig
}

// updateStatus computes a complete status based on the given observed state and writes it to status.
func (r *orcNetworkReconciler) updateStatus(ctx context.Context, orcObject *orcv1alpha1.Network, opts ...updateStatusOpt) error {
	now := metav1.NewTime(time.Now())

	statusUpdate := createStatusUpdate(ctx, orcObject, now, opts...)

	return r.client.Status().Patch(ctx, orcObject, ssa.ApplyConfigPatch(statusUpdate), client.ForceOwnership, ssaFieldOwner(SSAStatusTxn))
}
