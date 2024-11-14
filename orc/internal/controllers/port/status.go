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

package port

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/gophercloud/gophercloud/v2/openstack/networking/v2/ports"
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

// setStatusID sets status.ID in its own SSA transaction.
func (r *orcPortReconciler) setStatusID(ctx context.Context, obj client.Object, id string) error {
	applyConfig := orcapplyconfigv1alpha1.Port(obj.GetName(), obj.GetNamespace()).
		WithUID(obj.GetUID()).
		WithStatus(orcapplyconfigv1alpha1.PortStatus().
			WithID(id))

	return r.client.Status().Patch(ctx, obj, ssa.ApplyConfigPatch(applyConfig), client.ForceOwnership, ssaFieldOwner(SSAIDTxn))
}

// setStatusNetworkID sets status.NetworkID in its own SSA transaction.
func (r *orcPortReconciler) setStatusNetworkID(ctx context.Context, obj client.Object, id orcv1alpha1.UUID) error {
	applyConfig := orcapplyconfigv1alpha1.Port(obj.GetName(), obj.GetNamespace()).
		WithUID(obj.GetUID()).
		WithStatus(orcapplyconfigv1alpha1.PortStatus().
			WithNetworkID(id))

	return r.client.Status().Patch(ctx, obj, ssa.ApplyConfigPatch(applyConfig), client.ForceOwnership, ssaFieldOwner(SSANetworkIDTxn))
}

type updateStatusOpts struct {
	resource        *ports.Port
	progressMessage *string
	err             error
}

type updateStatusOpt func(*updateStatusOpts)

func withResource(resource *ports.Port) updateStatusOpt {
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

func getOSResourceStatus(osResource *ports.Port) *orcapplyconfigv1alpha1.PortResourceStatusApplyConfiguration {
	status := orcapplyconfigv1alpha1.PortResourceStatus().
		WithName(osResource.Name).
		WithDescription(osResource.Description).
		WithAdminStateUp(osResource.AdminStateUp).
		WithStatus(osResource.Status).
		WithProjectID(osResource.ProjectID).
		WithTags(osResource.Tags...).
		WithRevisionNumber(int64(osResource.RevisionNumber)).
		WithCreatedAt(metav1.NewTime(osResource.CreatedAt)).
		WithUpdatedAt(metav1.NewTime(osResource.UpdatedAt))

	return status
}

// createStatusUpdate computes a complete status update based on the given
// observed state. This is separated from updateStatus to facilitate unit
// testing, as the version of k8s we currently import does not support patch
// apply in the fake client.
// Needs: https://github.com/kubernetes/kubernetes/pull/125560
func createStatusUpdate(ctx context.Context, orcPort *orcv1alpha1.Port, now metav1.Time, opts ...updateStatusOpt) *orcapplyconfigv1alpha1.PortApplyConfiguration {
	log := ctrl.LoggerFrom(ctx)

	statusOpts := updateStatusOpts{}
	for i := range opts {
		opts[i](&statusOpts)
	}

	osResource := statusOpts.resource
	err := statusOpts.err

	applyConfigStatus := orcapplyconfigv1alpha1.PortStatus()
	applyConfig := orcapplyconfigv1alpha1.Port(orcPort.Name, orcPort.Namespace).WithStatus(applyConfigStatus)

	if osResource != nil {
		resourceStatus := getOSResourceStatus(osResource)
		applyConfigStatus.WithResource(resourceStatus)
	}

	availableCondition := applyconfigv1.Condition().
		WithType(orcv1alpha1.OpenStackConditionAvailable).
		WithObservedGeneration(orcPort.Generation)
	progressingCondition := applyconfigv1.Condition().
		WithType(orcv1alpha1.OpenStackConditionProgressing).
		WithObservedGeneration(orcPort.Generation)

	// A port is available as soon as it exists
	available := osResource != nil
	if available {
		availableCondition.
			WithStatus(metav1.ConditionTrue).
			WithReason(orcv1alpha1.OpenStackConditionReasonSuccess).
			WithMessage("OpenStack resource is available")
	} else {
		// Port is not available. Reason and message will be copied from Progressing
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
		previous := meta.FindStatusCondition(orcPort.Status.Conditions, *condition.Type)
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
		log.V(4).Info("Setting resource status", logValues...)
	}

	return applyConfig
}

// updateStatus computes a complete status based on the given observed state and writes it to status.
func (r *orcPortReconciler) updateStatus(ctx context.Context, orcObject *orcv1alpha1.Port, opts ...updateStatusOpt) error {
	now := metav1.NewTime(time.Now())

	statusUpdate := createStatusUpdate(ctx, orcObject, now, opts...)

	return r.client.Status().Patch(ctx, orcObject, ssa.ApplyConfigPatch(statusUpdate), client.ForceOwnership, ssaFieldOwner(SSAStatusTxn))
}
