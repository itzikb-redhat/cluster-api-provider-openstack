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

package v1alpha1

// NetworkSegment contains provider-network properties. Currently only
// available in status.
type NetworkSegment struct {
	// ProviderNetworkType is the type of physical network that this
	// network should be mapped to. For example, flat, vlan, vxlan, or gre.
	// Valid values depend on a networking back-end.
	ProviderNetworkType *string `json:"providerNetworkType,omitempty"`

	// ProviderPhysicalNetwork is the physical network where this network
	// should be implemented. The Networking API v2.0 does not provide a
	// way to list available physical networks. For example, the Open
	// vSwitch plug-in configuration file defines a symbolic name that maps
	// to specific bridges on each compute host.
	ProviderPhysicalNetwork *string `json:"providerPhysicalNetwork,omitempty"`

	// ProviderSegmentationID is the ID of the isolated segment on the
	// physical network. The network_type attribute defines the
	// segmentation model. For example, if the network_type value is vlan,
	// this ID is a vlan identifier. If the network_type value is gre, this
	// ID is a gre key.
	ProviderSegmentationID *int32 `json:"providerSegmentationID,omitempty"`
}

// NetworkResourceSpec contains the desired state of a network
// +kubebuilder:validation:XValidation:rule="has(self.name) ? self.name == oldSelf.name : !has(oldSelf.name)",message="name is immutable"
// +kubebuilder:validation:XValidation:rule="has(self.description) ? self.description == oldSelf.description : !has(oldSelf.description)",message="description is immutable"
// +kubebuilder:validation:XValidation:rule="has(self.adminStateUp) ? self.adminStateUp == oldSelf.adminStateUp : !has(oldSelf.adminStateUp)",message="adminStateUp is immutable"
// +kubebuilder:validation:XValidation:rule="has(self.dnsDomain) ? self.dnsDomain == oldSelf.dnsDomain : !has(oldSelf.dnsDomain)",message="dnsDomain is immutable"
// +kubebuilder:validation:XValidation:rule="has(self.mtu) ? self.mtu == oldSelf.mtu : !has(oldSelf.mtu)",message="mtu is immutable"
// +kubebuilder:validation:XValidation:rule="has(self.portSecurityEnabled) ? self.portSecurityEnabled == oldSelf.portSecurityEnabled : !has(oldSelf.portSecurityEnabled)",message="portSecurityEnabled is immutable"
// +kubebuilder:validation:XValidation:rule="has(self.qosPolicyID) ? self.qosPolicyID == oldSelf.qosPolicyID : !has(oldSelf.qosPolicyID)",message="qosPolicyID is immutable"
// +kubebuilder:validation:XValidation:rule="has(self.external) ? self.external == oldSelf.external : !has(oldSelf.external)",message="external is immutable"
// +kubebuilder:validation:XValidation:rule="has(self.shared) ? self.shared == oldSelf.shared : !has(oldSelf.shared)",message="shared is immutable"
// +kubebuilder:validation:XValidation:rule="has(self.vlanTransparent) ? self.vlanTransparent == oldSelf.vlanTransparent : !has(oldSelf.vlanTransparent)",message="vlanTransparent is immutable"
// +kubebuilder:validation:XValidation:rule="has(self.availabilityZoneHints) ? self.availabilityZoneHints == oldSelf.availabilityZoneHints : !has(oldSelf.availabilityZoneHints)",message="availabilityZoneHints is immutable"
type NetworkResourceSpec struct {
	// Name will be the name of the created resource. If not specified, the
	// name of the ORC object will be used.
	// +kubebuilder:validation:MinLength:=1
	// +kubebuilder:validation:MaxLength:=1000
	// +optional
	Name *string `json:"name,omitempty"`

	// +optional
	Description *string `json:"description,omitempty"`

	// +optional
	AdminStateUp *bool `json:"adminStateUp,omitempty"`

	// +optional
	DNSDomain *string `json:"dnsDomain,omitempty"`

	// MTU is the the maximum transmission unit value to address
	// fragmentation. Minimum value is 68 for IPv4, and 1280 for IPv6.
	// +optional
	// +kubebuilder:validation:Minimum:=68
	MTU *int32 `json:"mtu,omitempty"`

	// PortSecurityEnabled is the port security status of the network.
	// Valid values are enabled (true) and disabled (false). This value is
	// used as the default value of port_security_enabled field of a newly
	// created port.
	// +optional
	PortSecurityEnabled *bool `json:"portSecurityEnabled,omitempty"`

	// QOSPolicyID is the ID of the QoS policy associated with the network.
	// +optional
	QOSPolicyID *string `json:"qosPolicyID,omitempty"`

	// External indicates whether the network has an external routing
	// facility that’s not managed by the networking service.
	// +optional
	External *bool `json:"external,omitempty"`

	// Shared indicates whether this resource is shared across all
	// projects. By default, only administrative users can change this
	// value.
	// +optional
	Shared *bool `json:"shared,omitempty"`

	// VLANTransparent indicates the VLAN transparency mode of the network,
	// which is VLAN transparent (true) or not VLAN transparent (false).
	// +optional
	VLANTransparent *bool `json:"vlanTransparent,omitempty"`

	// AvailabilityZoneHints is the availability zone candidate for the network.
	// +listType=set
	// +optional
	AvailabilityZoneHints []string `json:"availabilityZoneHints,omitempty"`
}

// NetworkFilter defines an existing resource by its properties
// +kubebuilder:validation:MinProperties:=1
type NetworkFilter struct {
	// Name of the existing resource
	// +optional
	Name *string `json:"name,omitempty"`

	// Description of the existing resource
	// +optional
	Description *string `json:"description,omitempty"`

	// External indicates whether the network has an external routing
	// facility that’s not managed by the networking service.
	// +optional
	External *bool `json:"external,omitempty"`
}

// NetworkResourceStatus represents the observed state of the resource.
type NetworkResourceStatus struct {
	// AdminStateUp is the administrative state of the network,
	// which is up (true) or down (false).
	// +optional
	AdminStateUp *bool `json:"adminStateUp,omitempty"`

	// AvailabilityZoneHints is the availability zone candidate for the
	// network.
	// +listType=atomic
	// +optional
	AvailabilityZoneHints []string `json:"availabilityZoneHints,omitempty"`

	// Availability is the availability zone for the network.
	// +listType=atomic
	// +optional
	AvailabilityZones []string `json:"availabilityZones,omitempty"`

	// CreatedAt contains the timestamp of when the resource was created.
	// +optional
	CreatedAt *string `json:"createdAt,omitempty"`

	// +optional
	DNSDomain *string `json:"dnsDomain,omitempty"`

	// UUID for the network
	// +optional
	ID string `json:"id,omitempty"`

	// IPV4AddressScope is the ID of the IPv4 address scope that the
	// network is associated with.
	// +optional
	IPV4AddressScope string `json:"ipv4AddressScope,omitempty"`

	// IPV6AddressScope is the ID of the IPv6 address scope that the
	// network is associated with.
	// +optional
	IPV6AddressScope string `json:"ipv6AddressScope,omitempty"`

	// L2Adjacency indicates whether L2 connectivity is available
	// throughout the network.
	// +optional
	L2Adjacency *bool `json:"l2Adjacency,omitempty"`

	// MTU is the the maximum transmission unit value to address
	// fragmentation. Minimum value is 68 for IPv4, and 1280 for IPv6.
	// +optional
	MTU int32 `json:"mtu,omitempty"`

	// Human-readable name for the network. Might not be unique.
	// +optional
	Name string `json:"name,omitempty"`

	// PortSecurityEnabled is the port security status of the network.
	// Valid values are enabled (true) and disabled (false). This value is
	// used as the default value of port_security_enabled field of a newly
	// created port.
	// +optional
	PortSecurityEnabled *bool `json:"portSecurityEnabled,omitempty"`

	// ProjectID is the project owner of the network.
	// +optional
	ProjectID string `json:"projectID,omitempty"`

	// +optional
	Segment NetworkSegment `json:",inline"`

	// QOSPolicyID is the ID of the QoS policy associated with the network.
	// +optional
	QOSPolicyID string `json:"qosPolicyID,omitempty"`

	// RevisionNumber is the revision number of the resource.
	// +optional
	RevisionNumber int32 `json:"revisionNumber,omitempty"`

	// External defines whether the network may be used for creation of
	// floating IPs. Only networks with this flag may be an external
	// gateway for routers. The network must have an external routing
	// facility that is not managed by the networking service. If the
	// network is updated from external to internal the unused floating IPs
	// of this network are automatically deleted when extension
	// floatingip-autodelete-internal is present.
	// +optional
	External bool `json:"external,omitempty"`

	// Segment is a list of provider segment objects.
	// +listType=atomic
	// +optional
	Segments []NetworkSegment `json:"segments,omitempty"`

	// Specifies whether the network resource can be accessed by any tenant.
	// +optional
	Shared bool `json:"shared,omitempty"`

	// Indicates whether network is currently operational. Possible values
	// include `ACTIVE', `DOWN', `BUILD', or `ERROR'. Plug-ins might define
	// additional values.
	// +optional
	Status *string `json:"status,omitempty"`

	// Subnets associated with this network.
	// +listType=atomic
	// +optional
	Subnets []string `json:"subnets,omitempty"`

	// TenantID is the project owner of the network.
	// +optional
	TenantID *string `json:"tenantID,omitempty"`

	// UpdatedAt contains the timestamp of when the resource was last
	// changed.
	// +optional
	UpdatedAt *string `json:"updatedAt,omitempty"`

	// VLANTransparent indicates the VLAN transparency mode of the network,
	// which is VLAN transparent (true) or not VLAN transparent (false).
	// +optional
	VLANTransparent *bool `json:"vlanTransparent,omitempty"`

	// Description is a human-readable description for the resource.
	// +optional
	Description *string `json:"description,omitempty"`

	// +optional
	IsDefault *bool `json:"isDefault,omitempty"`

	// Tags is the list of tags on the resource.
	// +listType=atomic
	// +optional
	Tags []string `json:"tags,omitempty"`
}
