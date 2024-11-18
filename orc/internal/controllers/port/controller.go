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
	Finalizer = "openstack.k-orc.cloud/port"

	FieldOwner = "openstack.k-orc.cloud/portcontroller"
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

	// The time to wait between checking if a delete was successful
	deletePollingPeriod = 5 * time.Second
)

// orcPortReconciler reconciles an ORC Port.
type orcPortReconciler struct {
	client       client.Client
	recorder     record.EventRecorder
	scopeFactory scope.Factory
}

func New(client client.Client, recorder record.EventRecorder, scopeFactory scope.Factory) ctrlexport.SetupWithManager {
	return &orcPortReconciler{
		client:       client,
		recorder:     recorder,
		scopeFactory: scopeFactory,
	}
}

// SetupWithManager sets up the controller with the Manager.
func (r *orcPortReconciler) SetupWithManager(ctx context.Context, mgr ctrl.Manager, options controller.Options) error {
	log := mgr.GetLogger().WithValues("controller", "port")

	getNetworkRefsForPort := func(obj client.Object) []string {
		port, ok := obj.(*orcv1alpha1.Port)
		if !ok {
			return nil
		}
		return []string{string(port.Spec.NetworkRef)}
	}

	// Index ports by referenced network
	const networkRefPath = "spec.resource.networkRef"

	if err := mgr.GetFieldIndexer().IndexField(ctx, &orcv1alpha1.Port{}, networkRefPath, func(obj client.Object) []string {
		return getNetworkRefsForPort(obj)
	}); err != nil {
		return fmt.Errorf("adding ports by network index: %w", err)
	}

	getPortsForNetwork := func(obj *orcv1alpha1.Network) ([]orcv1alpha1.Port, error) {
		k8sClient := mgr.GetClient()
		portList := &orcv1alpha1.PortList{}
		if err := k8sClient.List(ctx, portList, client.InNamespace(obj.Namespace), client.MatchingFields{networkRefPath: obj.Name}); err != nil {
			return nil, err
		}

		return portList.Items, nil
	}

	err := ctrlcommon.AddDeletionGuard(mgr, Finalizer, FieldOwner, getNetworkRefsForPort, getPortsForNetwork)
	if err != nil {
		return err
	}

	return ctrl.NewControllerManagedBy(mgr).
		For(&orcv1alpha1.Port{}, builder.WithPredicates(ctrlcommon.NeedsReconcilePredicate(log))).
		Watches(&orcv1alpha1.Network{},
			handler.EnqueueRequestsFromMapFunc(func(ctx context.Context, obj client.Object) []reconcile.Request {
				log := log.WithValues("watch", "Network", "name", obj.GetName(), "namespace", obj.GetNamespace())

				network, ok := obj.(*orcv1alpha1.Network)
				if !ok {
					log.Info("Watch got unexpected object type", "type", fmt.Sprintf("%T", obj))
					return nil
				}

				ports, err := getPortsForNetwork(network)
				if err != nil {
					log.Error(err, "listing Ports")
					return nil
				}
				requests := make([]reconcile.Request, len(ports))
				for i := range ports {
					port := &ports[i]
					request := &requests[i]

					request.Name = port.Name
					request.Namespace = port.Namespace
				}
				return requests
			}),
			builder.WithPredicates(predicates.NewBecameAvailable(log, &orcv1alpha1.Network{})),
		).
		WithOptions(options).
		Complete(r)
}
