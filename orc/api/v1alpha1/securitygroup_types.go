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

// SecurityGroupResourceSpec contains the desired state of a security group
type SecurityGroupResourceSpec struct {
	// Name will be the name of the created resource. If not specified, the
	// name of the ORC object will be used.
	// +optional
	Name *OpenStackName `json:"name,omitempty"`

	// +optional
	Description *OpenStackDescription `json:"description,omitempty"`

	// Tags is a list of tags which will be applied to the security group.
	// +kubebuilder:validation:MaxItems:=32
	// +listType=set
	Tags []NeutronTag `json:"tags,omitempty"`

	// Stateful indicates if the security group is stateful or stateless.
	// +optional
	Stateful *bool `json:"stateful,omitempty"`
}

// SecurityGroupFilter defines an existing resource by its properties
// +kubebuilder:validation:MinProperties:=1
type SecurityGroupFilter struct {
	// Name of the existing resource
	// +optional
	Name *OpenStackName `json:"name,omitempty"`

	// Description of the existing resource
	// +optional
	Description *OpenStackDescription `json:"description,omitempty"`

	// ProjectID specifies the ID of the project which owns the security group.
	// +optional
	ProjectID *UUID `json:"projectID,omitempty"`

	FilterByNeutronTags `json:",inline"`
}

// SecurityGroupResourceStatus represents the observed state of the resource.
type SecurityGroupResourceStatus struct {
	// Human-readable name for the security group. Might not be unique.
	// +optional
	Name string `json:"name,omitempty"`

	// Description is a human-readable description for the resource.
	// +optional
	Description string `json:"description,omitempty"`

	// ProjectID is the project owner of the security group.
	// +optional
	ProjectID string `json:"projectID,omitempty"`

	// Tags is the list of tags on the resource.
	// +listType=atomic
	// +optional
	Tags []string `json:"tags,omitempty"`

	// Stateful indicates if the security group is stateful or stateless.
	// +optional
	Stateful bool `json:"stateful,omitempty"`

	NeutronStatusMetadata `json:",inline"`
}
