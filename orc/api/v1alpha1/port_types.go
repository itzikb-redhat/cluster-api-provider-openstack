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

type PortRefs struct {
	// NetworkRef is a reference to the ORC Network which this port is associated with.
	// +required
	NetworkRef ORCNameRef `json:"networkRef"`
}

// PortFilter specifies a filter to select a port. At least one parameter must be specified.
// +kubebuilder:validation:MinProperties:=1
type PortFilter struct {
	Name        *OpenStackName        `json:"name,omitempty"`
	Description *OpenStackDescription `json:"description,omitempty"`
	ProjectID   *UUID                 `json:"projectID,omitempty"`

	FilterByNeutronTags `json:",inline"`
}

type PortResourceSpec struct {
	// Name is a human-readable name of the port. If not set, the object's name will be used.
	// +optional
	Name *OpenStackName `json:"name,omitempty"`

	// Description of the port.
	// +optional
	Description *OpenStackDescription `json:"description,omitempty"`

	// Tags is a list of tags which will be applied to the port.
	// +kubebuilder:validation:MaxItems:=32
	// +listType=set
	Tags []NeutronTag `json:"tags,omitempty"`

	// ProjectID is the unique ID of the project which owns the Port. Only
	// administrative users can specify a project UUID other than their own.
	// +optional
	ProjectID *UUID `json:"projectID,omitempty"`
}

type PortResourceStatus struct {
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

	NeutronStatusMetadata `json:",inline"`
}

type PortStatusExtra struct {
	// NetworkID is the UUID of the parent network.
	NetworkID *UUID `json:"networkID,omitempty"`
}
