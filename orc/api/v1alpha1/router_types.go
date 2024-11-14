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

package v1alpha1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// RouterFilter specifies a query to select an OpenStack router. At least one property must be set.
// +kubebuilder:validation:MinProperties:=1
type RouterFilter struct {
	Name        OpenStackName        `json:"name,omitempty"`
	Description OpenStackDescription `json:"description,omitempty"`
	ProjectID   UUID                 `json:"projectID,omitempty"`

	FilterByNeutronTags `json:",inline"`
}

type ExternalGateway struct {
	NetworkRef KubernetesNameRef `json:"networkRef"`
}

type ExternalGatewayStatus struct {
	NetworkID string `json:"networkID"`
}

type RouterResourceSpec struct {
	// Name is the human-readable name of the subnet. Might not be unique.
	// +optional
	Name *OpenStackName `json:"name,omitempty"`

	// Description for the subnet.
	// +optional
	Description OpenStackDescription `json:"description,omitempty"`

	// Tags optionally set via extensions/attributestags
	// +listType=set
	Tags []NeutronTag `json:"tags,omitempty"`

	AdminStateUp *bool `json:"adminStateUp,omitempty"`

	// +listType=atomic
	// +optional
	ExternalGateways []ExternalGateway `json:"externalGateways,omitempty"`

	Distributed *bool `json:"distributed,omitempty"`

	// +listType=set
	// +optional
	AvailabilityZoneHints []AvailabilityZoneHint `json:"availabilityZoneHints,omitempty"`

	NeutronStatusMetadata `json:",inline"`
}

type RouterResourceStatus struct {
	// Name is the human-readable name of the resource. Might not be unique.
	// +optional
	Name string `json:"name,omitempty"`

	// Description is a human-readable description for the resource.
	// +optional
	Description string `json:"description,omitempty"`

	// ProjectID is the project owner of the resource.
	// +optional
	ProjectID string `json:"projectID,omitempty"`

	// Status indicates the current status of the resource.
	// +optional
	Status string `json:"status,omitempty"`

	// Tags is the list of tags on the resource.
	// +listType=atomic
	// +optional
	Tags []string `json:"tags,omitempty"`

	AdminStateUp bool `json:"adminStateUp"`

	// +listType=atomic
	// +optional
	ExternalGateways []ExternalGatewayStatus `json:"externalGateways,omitempty"`

	// +listType=atomic
	// +optional
	AvailabilityZoneHints []string `json:"availabilityZoneHints,omitempty"`
}

// +genclient
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Available",type="string",JSONPath=".status.conditions[?(@.type=='Available')].status",description="Availability status of resource"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Available')].message",description="Message describing current availability status"
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp",description="Time duration since creation"

// Router is the Schema for an ORC resource.
type RouterInterface struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RouterInterfaceSpec   `json:"spec,omitempty"`
	Status RouterInterfaceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RouterList contains a list of Router.
type RouterInterfaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Router `json:"items"`
}

// +kubebuilder:validation:Enum:=Subnet
// +kubebuilder:validation:MinLength:=1
// +kubebuilder:validation:MaxLength:=8
type RouterInterfaceType string

const (
	RouterInterfaceTypeSubnet RouterInterfaceType = "Subnet"
)

// +kubebuilder:validation:XValidation:rule="self.type == 'Subnet' ? has(self.subnetRef) : !has(self.subnetRef)",message="subnetRef is required when type is 'Subnet' and not permitted otherwise"
type RouterInterfaceSpec struct {
	// +required
	// +unionDiscriminator
	Type RouterInterfaceType `json:"type"`

	// +required
	RouterRef ORCNameRef `json:"routerRef"`

	// +unionMember,optional
	SubnetRef *ORCNameRef `json:"subnetRef,omitempty"`
}

type RouterInterfaceStatus struct {
	// Conditions represents the observed status of the object.
	// Known .status.conditions.type are: "Available", "Progressing"
	//
	// Available represents the availability of the OpenStack resource. If it is
	// true then the resource is ready for use.
	//
	// Progressing indicates whether the controller is still attempting to
	// reconcile the current state of the OpenStack resource to the desired
	// state. Progressing will be False either because the desired state has
	// been achieved, or because some terminal error prevents it from ever being
	// achieved and the controller is no longer attempting to reconcile. If
	// Progressing is True, an observer waiting on the resource should continue
	// to wait.
	//
	// +patchMergeKey=type
	// +patchStrategy=merge
	// +listType=map
	// +listMapKey=type
	Conditions []metav1.Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type"`
}

var _ ObjectWithConditions = &Router{}

func (i *RouterInterface) GetConditions() []metav1.Condition {
	return i.Status.Conditions
}

func init() {
	SchemeBuilder.Register(&RouterInterface{}, &RouterInterfaceList{})
}
