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
	"time"

	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"

	orcv1alpha1 "github.com/k-orc/openstack-resource-controller/api/v1alpha1"

	ctrlcommon "github.com/k-orc/openstack-resource-controller/internal/controllers/common"
	ctrlexport "github.com/k-orc/openstack-resource-controller/internal/controllers/export"
	"github.com/k-orc/openstack-resource-controller/internal/scope"
)

const (
	Finalizer = "openstack.k-orc.cloud/network"

	FieldOwner = "openstack.k-orc.cloud/networkcontroller"
	// Field owner of the object finalizer.
	SSAFinalizerTxn = "finalizer"
	// Field owner of transient status.
	SSAStatusTxn = "status"
	// Field owner of persistent id field.
	SSAIDTxn = "id"
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

// orcNetworkReconciler reconciles an ORC Subnet.
type orcNetworkReconciler struct {
	client           client.Client
	recorder         record.EventRecorder
	watchFilterValue string
	scopeFactory     scope.Factory
	caCertificates   []byte // PEM encoded ca certificates.
}

func New(client client.Client, recorder record.EventRecorder, watchFilterValue string, scopeFactory scope.Factory, caCertificates []byte) ctrlexport.SetupWithManager {
	return &orcNetworkReconciler{
		client:           client,
		recorder:         recorder,
		watchFilterValue: watchFilterValue,
		scopeFactory:     scopeFactory,
		caCertificates:   caCertificates,
	}
}

// SetupWithManager sets up the controller with the Manager.
func (r *orcNetworkReconciler) SetupWithManager(ctx context.Context, mgr ctrl.Manager, options controller.Options) error {
	log := mgr.GetLogger()

	return ctrl.NewControllerManagedBy(mgr).
		For(&orcv1alpha1.Network{}).
		WithOptions(options).
		WithEventFilter(ctrlcommon.NeedsReconcilePredicate(log)).
		Complete(r)
}
