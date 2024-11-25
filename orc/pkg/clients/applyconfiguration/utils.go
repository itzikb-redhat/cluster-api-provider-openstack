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

package applyconfiguration

import (
	v1alpha1 "github.com/k-orc/openstack-resource-controller/api/v1alpha1"
	apiv1alpha1 "github.com/k-orc/openstack-resource-controller/pkg/clients/applyconfiguration/api/v1alpha1"
	internal "github.com/k-orc/openstack-resource-controller/pkg/clients/applyconfiguration/internal"
	runtime "k8s.io/apimachinery/pkg/runtime"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	testing "k8s.io/client-go/testing"
)

// ForKind returns an apply configuration type for the given GroupVersionKind, or nil if no
// apply configuration type exists for the given GroupVersionKind.
func ForKind(kind schema.GroupVersionKind) interface{} {
	switch kind {
	// Group=openstack.k-orc.cloud, Version=v1alpha1
	case v1alpha1.SchemeGroupVersion.WithKind("AllocationPool"):
		return &apiv1alpha1.AllocationPoolApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("AllocationPoolStatus"):
		return &apiv1alpha1.AllocationPoolStatusApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("AllowedAddressPairStatus"):
		return &apiv1alpha1.AllowedAddressPairStatusApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("CloudCredentialsReference"):
		return &apiv1alpha1.CloudCredentialsReferenceApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("ExternalGateway"):
		return &apiv1alpha1.ExternalGatewayApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("ExternalGatewayStatus"):
		return &apiv1alpha1.ExternalGatewayStatusApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("FilterByNeutronTags"):
		return &apiv1alpha1.FilterByNeutronTagsApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("FixedIPStatus"):
		return &apiv1alpha1.FixedIPStatusApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("Flavor"):
		return &apiv1alpha1.FlavorApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("FlavorFilter"):
		return &apiv1alpha1.FlavorFilterApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("FlavorImport"):
		return &apiv1alpha1.FlavorImportApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("FlavorResourceSpec"):
		return &apiv1alpha1.FlavorResourceSpecApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("FlavorResourceStatus"):
		return &apiv1alpha1.FlavorResourceStatusApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("FlavorSpec"):
		return &apiv1alpha1.FlavorSpecApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("FlavorStatus"):
		return &apiv1alpha1.FlavorStatusApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("HostRoute"):
		return &apiv1alpha1.HostRouteApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("HostRouteStatus"):
		return &apiv1alpha1.HostRouteStatusApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("Image"):
		return &apiv1alpha1.ImageApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("ImageContent"):
		return &apiv1alpha1.ImageContentApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("ImageContentSourceDownload"):
		return &apiv1alpha1.ImageContentSourceDownloadApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("ImageFilter"):
		return &apiv1alpha1.ImageFilterApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("ImageHash"):
		return &apiv1alpha1.ImageHashApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("ImageImport"):
		return &apiv1alpha1.ImageImportApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("ImageProperties"):
		return &apiv1alpha1.ImagePropertiesApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("ImagePropertiesHardware"):
		return &apiv1alpha1.ImagePropertiesHardwareApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("ImageResourceSpec"):
		return &apiv1alpha1.ImageResourceSpecApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("ImageResourceStatus"):
		return &apiv1alpha1.ImageResourceStatusApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("ImageSpec"):
		return &apiv1alpha1.ImageSpecApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("ImageStatus"):
		return &apiv1alpha1.ImageStatusApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("ImageStatusExtra"):
		return &apiv1alpha1.ImageStatusExtraApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("IPv6Options"):
		return &apiv1alpha1.IPv6OptionsApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("ManagedOptions"):
		return &apiv1alpha1.ManagedOptionsApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("Network"):
		return &apiv1alpha1.NetworkApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("NetworkFilter"):
		return &apiv1alpha1.NetworkFilterApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("NetworkImport"):
		return &apiv1alpha1.NetworkImportApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("NetworkResourceSpec"):
		return &apiv1alpha1.NetworkResourceSpecApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("NetworkResourceStatus"):
		return &apiv1alpha1.NetworkResourceStatusApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("NetworkSpec"):
		return &apiv1alpha1.NetworkSpecApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("NetworkStatus"):
		return &apiv1alpha1.NetworkStatusApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("NeutronStatusMetadata"):
		return &apiv1alpha1.NeutronStatusMetadataApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("Port"):
		return &apiv1alpha1.PortApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("PortFilter"):
		return &apiv1alpha1.PortFilterApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("PortImport"):
		return &apiv1alpha1.PortImportApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("PortRefs"):
		return &apiv1alpha1.PortRefsApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("PortResourceSpec"):
		return &apiv1alpha1.PortResourceSpecApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("PortResourceStatus"):
		return &apiv1alpha1.PortResourceStatusApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("PortSpec"):
		return &apiv1alpha1.PortSpecApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("PortStatus"):
		return &apiv1alpha1.PortStatusApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("PortStatusExtra"):
		return &apiv1alpha1.PortStatusExtraApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("ProviderProperties"):
		return &apiv1alpha1.ProviderPropertiesApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("Router"):
		return &apiv1alpha1.RouterApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("RouterFilter"):
		return &apiv1alpha1.RouterFilterApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("RouterImport"):
		return &apiv1alpha1.RouterImportApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("RouterInterface"):
		return &apiv1alpha1.RouterInterfaceApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("RouterInterfaceSpec"):
		return &apiv1alpha1.RouterInterfaceSpecApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("RouterInterfaceStatus"):
		return &apiv1alpha1.RouterInterfaceStatusApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("RouterResourceSpec"):
		return &apiv1alpha1.RouterResourceSpecApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("RouterResourceStatus"):
		return &apiv1alpha1.RouterResourceStatusApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("RouterSpec"):
		return &apiv1alpha1.RouterSpecApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("RouterStatus"):
		return &apiv1alpha1.RouterStatusApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("Subnet"):
		return &apiv1alpha1.SubnetApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("SubnetFilter"):
		return &apiv1alpha1.SubnetFilterApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("SubnetGateway"):
		return &apiv1alpha1.SubnetGatewayApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("SubnetImport"):
		return &apiv1alpha1.SubnetImportApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("SubnetRefs"):
		return &apiv1alpha1.SubnetRefsApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("SubnetResourceSpec"):
		return &apiv1alpha1.SubnetResourceSpecApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("SubnetResourceStatus"):
		return &apiv1alpha1.SubnetResourceStatusApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("SubnetSpec"):
		return &apiv1alpha1.SubnetSpecApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("SubnetStatus"):
		return &apiv1alpha1.SubnetStatusApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("SubnetStatusExtra"):
		return &apiv1alpha1.SubnetStatusExtraApplyConfiguration{}

	}
	return nil
}

func NewTypeConverter(scheme *runtime.Scheme) *testing.TypeConverter {
	return &testing.TypeConverter{Scheme: scheme, TypeResolver: internal.Parser()}
}
