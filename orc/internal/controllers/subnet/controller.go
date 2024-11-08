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
	"fmt"
	"slices"
	"strings"
	"time"

	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/tools/record"
	"k8s.io/client-go/util/workqueue"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/builder"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	orcv1alpha1 "github.com/k-orc/openstack-resource-controller/api/v1alpha1"
	"github.com/k-orc/openstack-resource-controller/pkg/predicates"

	ctrlcommon "github.com/k-orc/openstack-resource-controller/internal/controllers/common"
	ctrlexport "github.com/k-orc/openstack-resource-controller/internal/controllers/export"
	"github.com/k-orc/openstack-resource-controller/internal/scope"
)

const (
	Finalizer = "openstack.k-orc.cloud/subnet"

	FieldOwner = "openstack.k-orc.cloud/subnetcontroller"
	// Field owner of the object finalizer.
	SSAFinalizerTxn = "finalizer"
	// Field owner of transient status.
	SSAStatusTxn = "status"
	// Field owner of persistent id field.
	SSAIDTxn = "id"
	// Field owner of persistent network id field.
	SSANetworkIDTxn = "networkID"
)

// ssaFieldOwner returns the field owner for a specific named SSA transaction.
func ssaFieldOwner(txn string) client.FieldOwner {
	return client.FieldOwner(FieldOwner + "/" + txn)
}

const (
	// The time to wait before reconciling again when we are expecting glance to finish some task and update status.
	externalUpdatePollingPeriod = 15 * time.Second
)

// orcSubnetReconciler reconciles an ORC Subnet.
type orcSubnetReconciler struct {
	client           client.Client
	recorder         record.EventRecorder
	watchFilterValue string
	scopeFactory     scope.Factory
}

func New(client client.Client, recorder record.EventRecorder, watchFilterValue string, scopeFactory scope.Factory) ctrlexport.SetupWithManager {
	return &orcSubnetReconciler{
		client:           client,
		recorder:         recorder,
		watchFilterValue: watchFilterValue,
		scopeFactory:     scopeFactory,
	}
}

// SetupWithManager sets up the controller with the Manager.
func (r *orcSubnetReconciler) SetupWithManager(ctx context.Context, mgr ctrl.Manager, options controller.Options) error {
	const networkRefPath = "spec.resource.networkRef"

	log := mgr.GetLogger()

	getNetworkRefsForSubnet := func(obj client.Object) []string {
		subnet, ok := obj.(*orcv1alpha1.Subnet)
		if !ok {
			return nil
		}
		return []string{string(subnet.Spec.NetworkRef)}
	}

	// Index subnets by referenced network
	if err := mgr.GetFieldIndexer().IndexField(ctx, &orcv1alpha1.Subnet{}, networkRefPath, func(obj client.Object) []string {
		return getNetworkRefsForSubnet(obj)
	}); err != nil {
		return fmt.Errorf("adding subnets by network index: %w", err)
	}

	getSubnetsForNetwork := func(obj *orcv1alpha1.Network) ([]orcv1alpha1.Subnet, error) {
		k8sClient := mgr.GetClient()
		subnetList := &orcv1alpha1.SubnetList{}
		if err := k8sClient.List(ctx, subnetList, client.InNamespace(obj.Namespace), client.MatchingFields{networkRefPath: obj.Name}); err != nil {
			return nil, err
		}

		return subnetList.Items, nil
	}

	const networkFinalizer = "subnet"

	err := addDeletionGuard(mgr, networkFinalizer, FieldOwner, getNetworkRefsForSubnet, getSubnetsForNetwork)
	if err != nil {
		return err
	}

	return ctrl.NewControllerManagedBy(mgr).
		For(&orcv1alpha1.Subnet{}).
		Watches(&orcv1alpha1.Network{},
			handler.EnqueueRequestsFromMapFunc(func(ctx context.Context, obj client.Object) []reconcile.Request {
				log = log.WithValues("watch", "Network")

				network, ok := obj.(*orcv1alpha1.Network)
				if !ok {
					return nil
				}

				subnets, err := getSubnetsForNetwork(network)
				if err != nil {
					log.Error(err, "listing Subnets")
				}
				requests := make([]reconcile.Request, len(subnets))
				for i := range subnets {
					subnet := &subnets[i]
					request := &requests[i]

					request.Name = subnet.Name
					request.Namespace = subnet.Namespace
				}
				return requests
			}),
			builder.WithPredicates(predicates.NewBecameAvailable(log, &orcv1alpha1.Network{})),
		).
		WithOptions(options).
		WithEventFilter(ctrlcommon.NeedsReconcilePredicate(log)).
		Complete(r)
}

type pointerToObject[T any] interface {
	*T
	client.Object
}

func addDeletionGuard[guardedP pointerToObject[guarded], dependencyP pointerToObject[dependency], guarded, dependency any](
	mgr ctrl.Manager, finalizer string, fieldOwner client.FieldOwner,
	getGuarded func(client.Object) []string,
	getDependencies func(guardedP) ([]dependency, error),
) error {
	deletionGuard := reconcile.Func(func(ctx context.Context, req reconcile.Request) (reconcile.Result, error) {
		k8sClient := mgr.GetClient()

		var guarded guardedP = new(guarded)
		err := k8sClient.Get(ctx, req.NamespacedName, guarded)
		if err != nil {
			if apierrors.IsNotFound(err) {
				return ctrl.Result{}, nil
			}
			return ctrl.Result{}, err
		}

		// If the object hasn't been deleted, we simply check that it has our finalizer
		if !guarded.GetDeletionTimestamp().IsZero() {
			if !slices.Contains(guarded.GetFinalizers(), finalizer) {
				patch := ctrlcommon.GetFinalizerPatch(guarded, finalizer)
				return ctrl.Result{}, k8sClient.Patch(ctx, guarded, patch, client.ForceOwnership, fieldOwner)
			}

			return ctrl.Result{}, nil
		}

		dependencies, err := getDependencies(guarded)
		if err != nil {
			return reconcile.Result{}, nil
		}
		if len(dependencies) == 0 {
			patch := ctrlcommon.RemoveFinalizerPatch(guarded)
			return ctrl.Result{}, k8sClient.Patch(ctx, guarded, patch, client.ForceOwnership, fieldOwner)
		}
		return ctrl.Result{}, nil
	})

	var guardedSpecimen guarded
	var guardedSpecimenP guardedP = &guardedSpecimen
	var dependencySpecimen dependency
	var dependencySpecimenP dependencyP = &dependencySpecimen

	lowerType := func(v any) string {
		return strings.ToLower(fmt.Sprintf("%T", v))
	}

	controllerName := lowerType(guardedSpecimen) + "_deletion_guard_for_" + lowerType(dependencySpecimen)

	err := builder.ControllerManagedBy(mgr).
		For(guardedSpecimenP).
		Watches(dependencySpecimenP,
			handler.Funcs{
				DeleteFunc: func(ctx context.Context, evt event.TypedDeleteEvent[client.Object], q workqueue.TypedRateLimitingInterface[reconcile.Request]) {
					for _, guarded := range getGuarded(evt.Object) {
						q.Add(reconcile.Request{
							NamespacedName: types.NamespacedName{
								Namespace: evt.Object.GetNamespace(),
								Name:      guarded,
							},
						})
					}
				},
			},
		).
		Named(controllerName).
		Complete(deletionGuard)

	if err != nil {
		return fmt.Errorf("failed to construct %s deletion guard for %s controller", lowerType(guardedSpecimen), lowerType(dependencySpecimen))
	}

	return nil
}
