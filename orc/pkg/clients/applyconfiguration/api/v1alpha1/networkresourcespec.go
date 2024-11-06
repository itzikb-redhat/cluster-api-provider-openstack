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

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/k-orc/openstack-resource-controller/api/v1alpha1"
)

// NetworkResourceSpecApplyConfiguration represents a declarative configuration of the NetworkResourceSpec type for use
// with apply.
type NetworkResourceSpecApplyConfiguration struct {
	Name                  *v1alpha1.OpenStackName        `json:"name,omitempty"`
	Description           *v1alpha1.OpenStackDescription `json:"description,omitempty"`
	Tags                  []v1alpha1.NeutronTag          `json:"tags,omitempty"`
	AdminStateUp          *bool                          `json:"adminStateUp,omitempty"`
	DNSDomain             *v1alpha1.DNSDomain            `json:"dnsDomain,omitempty"`
	MTU                   *v1alpha1.MTU                  `json:"mtu,omitempty"`
	PortSecurityEnabled   *bool                          `json:"portSecurityEnabled,omitempty"`
	External              *bool                          `json:"external,omitempty"`
	Shared                *bool                          `json:"shared,omitempty"`
	AvailabilityZoneHints []string                       `json:"availabilityZoneHints,omitempty"`
}

// NetworkResourceSpecApplyConfiguration constructs a declarative configuration of the NetworkResourceSpec type for use with
// apply.
func NetworkResourceSpec() *NetworkResourceSpecApplyConfiguration {
	return &NetworkResourceSpecApplyConfiguration{}
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *NetworkResourceSpecApplyConfiguration) WithName(value v1alpha1.OpenStackName) *NetworkResourceSpecApplyConfiguration {
	b.Name = &value
	return b
}

// WithDescription sets the Description field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Description field is set to the value of the last call.
func (b *NetworkResourceSpecApplyConfiguration) WithDescription(value v1alpha1.OpenStackDescription) *NetworkResourceSpecApplyConfiguration {
	b.Description = &value
	return b
}

// WithTags adds the given value to the Tags field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Tags field.
func (b *NetworkResourceSpecApplyConfiguration) WithTags(values ...v1alpha1.NeutronTag) *NetworkResourceSpecApplyConfiguration {
	for i := range values {
		b.Tags = append(b.Tags, values[i])
	}
	return b
}

// WithAdminStateUp sets the AdminStateUp field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the AdminStateUp field is set to the value of the last call.
func (b *NetworkResourceSpecApplyConfiguration) WithAdminStateUp(value bool) *NetworkResourceSpecApplyConfiguration {
	b.AdminStateUp = &value
	return b
}

// WithDNSDomain sets the DNSDomain field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DNSDomain field is set to the value of the last call.
func (b *NetworkResourceSpecApplyConfiguration) WithDNSDomain(value v1alpha1.DNSDomain) *NetworkResourceSpecApplyConfiguration {
	b.DNSDomain = &value
	return b
}

// WithMTU sets the MTU field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the MTU field is set to the value of the last call.
func (b *NetworkResourceSpecApplyConfiguration) WithMTU(value v1alpha1.MTU) *NetworkResourceSpecApplyConfiguration {
	b.MTU = &value
	return b
}

// WithPortSecurityEnabled sets the PortSecurityEnabled field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the PortSecurityEnabled field is set to the value of the last call.
func (b *NetworkResourceSpecApplyConfiguration) WithPortSecurityEnabled(value bool) *NetworkResourceSpecApplyConfiguration {
	b.PortSecurityEnabled = &value
	return b
}

// WithExternal sets the External field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the External field is set to the value of the last call.
func (b *NetworkResourceSpecApplyConfiguration) WithExternal(value bool) *NetworkResourceSpecApplyConfiguration {
	b.External = &value
	return b
}

// WithShared sets the Shared field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Shared field is set to the value of the last call.
func (b *NetworkResourceSpecApplyConfiguration) WithShared(value bool) *NetworkResourceSpecApplyConfiguration {
	b.Shared = &value
	return b
}

// WithAvailabilityZoneHints adds the given value to the AvailabilityZoneHints field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the AvailabilityZoneHints field.
func (b *NetworkResourceSpecApplyConfiguration) WithAvailabilityZoneHints(values ...string) *NetworkResourceSpecApplyConfiguration {
	for i := range values {
		b.AvailabilityZoneHints = append(b.AvailabilityZoneHints, values[i])
	}
	return b
}
