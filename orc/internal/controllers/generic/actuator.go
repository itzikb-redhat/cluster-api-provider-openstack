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

package generic

import (
	"context"

	"github.com/go-logr/logr"

	orcv1alpha1 "github.com/k-orc/openstack-resource-controller/api/v1alpha1"
	orcerrors "github.com/k-orc/openstack-resource-controller/internal/util/errors"
)

type ResourceActuator[osResourceT any] interface {
	GetManagementPolicy() orcv1alpha1.ManagementPolicy
	GetResourceID(osResource osResourceT) string

	GetOSResourceByStatusID(ctx context.Context) (osResourceT, error)
	GetOSResourceByImportID(ctx context.Context) (osResourceT, error)
	GetOSResourceByImportFilter(ctx context.Context) (bool, osResourceT, error)
	GetOSResourceBySpec(ctx context.Context) (osResourceT, error)
	CreateResource(ctx context.Context) (osResourceT, error)
}

func GetOrCreateOSResource[osResourcePT *osResourceT, osResourceT any](ctx context.Context, log logr.Logger, actuator ResourceActuator[osResourcePT]) (osResourcePT, error) {
	var osResource osResourcePT
	var err error

	// Get by status ID
	osResource, err = actuator.GetOSResourceByStatusID(ctx)
	if err != nil {
		if orcerrors.IsNotFound(err) {
			// An OpenStack resource we previously referenced has been deleted unexpectedly. We can't recover from this.
			err = orcerrors.Terminal(orcv1alpha1.OpenStackConditionReasonUnrecoverableError, "resource has been deleted from OpenStack")
		}
		return nil, err
	}
	if osResource != nil {
		log.V(4).Info("Got existing OpenStack resource", "ID", actuator.GetResourceID(osResource))
		return osResource, nil
	}

	// Import by ID
	osResource, err = actuator.GetOSResourceByImportID(ctx)
	if err != nil {
		if orcerrors.IsNotFound(err) {
			// We assume that a resource imported by ID must already exist. It's a terminal error if it doesn't.
			err = orcerrors.Terminal(orcv1alpha1.OpenStackConditionReasonUnrecoverableError, "referenced resource does not exist in OpenStack")
		}
		return nil, err
	}
	if osResource != nil {
		log.V(4).Info("Imported existing OpenStack resource by ID", "ID", actuator.GetResourceID(osResource))
		return osResource, nil
	}

	// Import by filter
	var hasImportFilter bool
	hasImportFilter, osResource, err = actuator.GetOSResourceByImportFilter(ctx)
	if err != nil {
		return nil, err
	}
	if hasImportFilter || osResource != nil {
		return osResource, nil
	}

	// Create
	if actuator.GetManagementPolicy() == orcv1alpha1.ManagementPolicyUnmanaged {
		// We never create an unmanaged resource
		// API validation should have ensured that one of the above functions returned
		return nil, orcerrors.Terminal(orcv1alpha1.OpenStackConditionReasonInvalidConfiguration, "Not creating unmanaged resource")
	}

	osResource, err = actuator.GetOSResourceBySpec(ctx)
	if err != nil {
		return nil, err
	}
	if osResource != nil {
		log.V(4).Info("Adopted previously created resource")
		return osResource, nil
	}

	log.V(4).Info("Creating resource")
	return actuator.CreateResource(ctx)
}
