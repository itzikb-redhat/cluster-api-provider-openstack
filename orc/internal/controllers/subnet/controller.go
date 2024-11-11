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
	"time"

	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/builder"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
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

	// The time to wait between polling for subnet deletion
	deletePollingPeriod = 1 * time.Second
)

// orcSubnetReconciler reconciles an ORC Subnet.
type orcSubnetReconciler struct {
	client       client.Client
	recorder     record.EventRecorder
	scopeFactory scope.Factory
}

func New(client client.Client, recorder record.EventRecorder, scopeFactory scope.Factory) ctrlexport.SetupWithManager {
	return &orcSubnetReconciler{
		client:       client,
		recorder:     recorder,
		scopeFactory: scopeFactory,
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

	err := ctrlcommon.AddDeletionGuard(mgr, networkFinalizer, FieldOwner, getNetworkRefsForSubnet, getSubnetsForNetwork)
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
