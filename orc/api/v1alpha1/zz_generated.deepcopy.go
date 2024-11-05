//go:build !ignore_autogenerated

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AllocationPool) DeepCopyInto(out *AllocationPool) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AllocationPool.
func (in *AllocationPool) DeepCopy() *AllocationPool {
	if in == nil {
		return nil
	}
	out := new(AllocationPool)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudCredentialsReference) DeepCopyInto(out *CloudCredentialsReference) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudCredentialsReference.
func (in *CloudCredentialsReference) DeepCopy() *CloudCredentialsReference {
	if in == nil {
		return nil
	}
	out := new(CloudCredentialsReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FilterByNeutronTags) DeepCopyInto(out *FilterByNeutronTags) {
	*out = *in
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]NeutronTag, len(*in))
		copy(*out, *in)
	}
	if in.TagsAny != nil {
		in, out := &in.TagsAny, &out.TagsAny
		*out = make([]NeutronTag, len(*in))
		copy(*out, *in)
	}
	if in.NotTags != nil {
		in, out := &in.NotTags, &out.NotTags
		*out = make([]NeutronTag, len(*in))
		copy(*out, *in)
	}
	if in.NotTagsAny != nil {
		in, out := &in.NotTagsAny, &out.NotTagsAny
		*out = make([]NeutronTag, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FilterByNeutronTags.
func (in *FilterByNeutronTags) DeepCopy() *FilterByNeutronTags {
	if in == nil {
		return nil
	}
	out := new(FilterByNeutronTags)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HostRoute) DeepCopyInto(out *HostRoute) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HostRoute.
func (in *HostRoute) DeepCopy() *HostRoute {
	if in == nil {
		return nil
	}
	out := new(HostRoute)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPv6Options) DeepCopyInto(out *IPv6Options) {
	*out = *in
	if in.AddressMode != nil {
		in, out := &in.AddressMode, &out.AddressMode
		*out = new(IPv6AddressMode)
		**out = **in
	}
	if in.RAMode != nil {
		in, out := &in.RAMode, &out.RAMode
		*out = new(IPv6RAMode)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPv6Options.
func (in *IPv6Options) DeepCopy() *IPv6Options {
	if in == nil {
		return nil
	}
	out := new(IPv6Options)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Image) DeepCopyInto(out *Image) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Image.
func (in *Image) DeepCopy() *Image {
	if in == nil {
		return nil
	}
	out := new(Image)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Image) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageContent) DeepCopyInto(out *ImageContent) {
	*out = *in
	if in.Download != nil {
		in, out := &in.Download, &out.Download
		*out = new(ImageContentSourceDownload)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageContent.
func (in *ImageContent) DeepCopy() *ImageContent {
	if in == nil {
		return nil
	}
	out := new(ImageContent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageContentSourceDownload) DeepCopyInto(out *ImageContentSourceDownload) {
	*out = *in
	if in.Decompress != nil {
		in, out := &in.Decompress, &out.Decompress
		*out = new(ImageCompression)
		**out = **in
	}
	if in.Hash != nil {
		in, out := &in.Hash, &out.Hash
		*out = new(ImageHash)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageContentSourceDownload.
func (in *ImageContentSourceDownload) DeepCopy() *ImageContentSourceDownload {
	if in == nil {
		return nil
	}
	out := new(ImageContentSourceDownload)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageFilter) DeepCopyInto(out *ImageFilter) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageFilter.
func (in *ImageFilter) DeepCopy() *ImageFilter {
	if in == nil {
		return nil
	}
	out := new(ImageFilter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageHash) DeepCopyInto(out *ImageHash) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageHash.
func (in *ImageHash) DeepCopy() *ImageHash {
	if in == nil {
		return nil
	}
	out := new(ImageHash)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageImport) DeepCopyInto(out *ImageImport) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Filter != nil {
		in, out := &in.Filter, &out.Filter
		*out = new(ImageFilter)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageImport.
func (in *ImageImport) DeepCopy() *ImageImport {
	if in == nil {
		return nil
	}
	out := new(ImageImport)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageList) DeepCopyInto(out *ImageList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Image, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageList.
func (in *ImageList) DeepCopy() *ImageList {
	if in == nil {
		return nil
	}
	out := new(ImageList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ImageList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageProperties) DeepCopyInto(out *ImageProperties) {
	*out = *in
	if in.MinDiskGB != nil {
		in, out := &in.MinDiskGB, &out.MinDiskGB
		*out = new(int)
		**out = **in
	}
	if in.MinMemoryMB != nil {
		in, out := &in.MinMemoryMB, &out.MinMemoryMB
		*out = new(int)
		**out = **in
	}
	if in.Hardware != nil {
		in, out := &in.Hardware, &out.Hardware
		*out = new(ImagePropertiesHardware)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageProperties.
func (in *ImageProperties) DeepCopy() *ImageProperties {
	if in == nil {
		return nil
	}
	out := new(ImageProperties)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImagePropertiesHardware) DeepCopyInto(out *ImagePropertiesHardware) {
	*out = *in
	if in.CPUSockets != nil {
		in, out := &in.CPUSockets, &out.CPUSockets
		*out = new(int)
		**out = **in
	}
	if in.CPUCores != nil {
		in, out := &in.CPUCores, &out.CPUCores
		*out = new(int)
		**out = **in
	}
	if in.CPUThreads != nil {
		in, out := &in.CPUThreads, &out.CPUThreads
		*out = new(int)
		**out = **in
	}
	if in.CPUPolicy != nil {
		in, out := &in.CPUPolicy, &out.CPUPolicy
		*out = new(string)
		**out = **in
	}
	if in.CPUThreadPolicy != nil {
		in, out := &in.CPUThreadPolicy, &out.CPUThreadPolicy
		*out = new(string)
		**out = **in
	}
	if in.CDROMBus != nil {
		in, out := &in.CDROMBus, &out.CDROMBus
		*out = new(ImageHWBus)
		**out = **in
	}
	if in.DiskBus != nil {
		in, out := &in.DiskBus, &out.DiskBus
		*out = new(ImageHWBus)
		**out = **in
	}
	if in.SCSIModel != nil {
		in, out := &in.SCSIModel, &out.SCSIModel
		*out = new(string)
		**out = **in
	}
	if in.VIFModel != nil {
		in, out := &in.VIFModel, &out.VIFModel
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImagePropertiesHardware.
func (in *ImagePropertiesHardware) DeepCopy() *ImagePropertiesHardware {
	if in == nil {
		return nil
	}
	out := new(ImagePropertiesHardware)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageResourceSpec) DeepCopyInto(out *ImageResourceSpec) {
	*out = *in
	if in.Protected != nil {
		in, out := &in.Protected, &out.Protected
		*out = new(bool)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]ImageTag, len(*in))
		copy(*out, *in)
	}
	if in.Visibility != nil {
		in, out := &in.Visibility, &out.Visibility
		*out = new(ImageVisibility)
		**out = **in
	}
	if in.Properties != nil {
		in, out := &in.Properties, &out.Properties
		*out = new(ImageProperties)
		(*in).DeepCopyInto(*out)
	}
	if in.Content != nil {
		in, out := &in.Content, &out.Content
		*out = new(ImageContent)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageResourceSpec.
func (in *ImageResourceSpec) DeepCopy() *ImageResourceSpec {
	if in == nil {
		return nil
	}
	out := new(ImageResourceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageResourceStatus) DeepCopyInto(out *ImageResourceStatus) {
	*out = *in
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
	if in.Hash != nil {
		in, out := &in.Hash, &out.Hash
		*out = new(ImageHash)
		**out = **in
	}
	if in.SizeB != nil {
		in, out := &in.SizeB, &out.SizeB
		*out = new(int64)
		**out = **in
	}
	if in.VirtualSizeB != nil {
		in, out := &in.VirtualSizeB, &out.VirtualSizeB
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageResourceStatus.
func (in *ImageResourceStatus) DeepCopy() *ImageResourceStatus {
	if in == nil {
		return nil
	}
	out := new(ImageResourceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageSpec) DeepCopyInto(out *ImageSpec) {
	*out = *in
	if in.Import != nil {
		in, out := &in.Import, &out.Import
		*out = new(ImageImport)
		(*in).DeepCopyInto(*out)
	}
	if in.Resource != nil {
		in, out := &in.Resource, &out.Resource
		*out = new(ImageResourceSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.ManagedOptions != nil {
		in, out := &in.ManagedOptions, &out.ManagedOptions
		*out = new(ManagedOptions)
		**out = **in
	}
	out.CloudCredentialsRef = in.CloudCredentialsRef
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageSpec.
func (in *ImageSpec) DeepCopy() *ImageSpec {
	if in == nil {
		return nil
	}
	out := new(ImageSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageStatus) DeepCopyInto(out *ImageStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Resource != nil {
		in, out := &in.Resource, &out.Resource
		*out = new(ImageResourceStatus)
		(*in).DeepCopyInto(*out)
	}
	in.ImageStatusExtra.DeepCopyInto(&out.ImageStatusExtra)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageStatus.
func (in *ImageStatus) DeepCopy() *ImageStatus {
	if in == nil {
		return nil
	}
	out := new(ImageStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageStatusExtra) DeepCopyInto(out *ImageStatusExtra) {
	*out = *in
	if in.DownloadAttempts != nil {
		in, out := &in.DownloadAttempts, &out.DownloadAttempts
		*out = new(int)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageStatusExtra.
func (in *ImageStatusExtra) DeepCopy() *ImageStatusExtra {
	if in == nil {
		return nil
	}
	out := new(ImageStatusExtra)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedOptions) DeepCopyInto(out *ManagedOptions) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedOptions.
func (in *ManagedOptions) DeepCopy() *ManagedOptions {
	if in == nil {
		return nil
	}
	out := new(ManagedOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Network) DeepCopyInto(out *Network) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Network.
func (in *Network) DeepCopy() *Network {
	if in == nil {
		return nil
	}
	out := new(Network)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Network) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkFilter) DeepCopyInto(out *NetworkFilter) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(OpenStackName)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(OpenStackDescription)
		**out = **in
	}
	if in.External != nil {
		in, out := &in.External, &out.External
		*out = new(bool)
		**out = **in
	}
	if in.ProjectID != nil {
		in, out := &in.ProjectID, &out.ProjectID
		*out = new(UUID)
		**out = **in
	}
	in.FilterByNeutronTags.DeepCopyInto(&out.FilterByNeutronTags)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkFilter.
func (in *NetworkFilter) DeepCopy() *NetworkFilter {
	if in == nil {
		return nil
	}
	out := new(NetworkFilter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkImport) DeepCopyInto(out *NetworkImport) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Filter != nil {
		in, out := &in.Filter, &out.Filter
		*out = new(NetworkFilter)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkImport.
func (in *NetworkImport) DeepCopy() *NetworkImport {
	if in == nil {
		return nil
	}
	out := new(NetworkImport)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkList) DeepCopyInto(out *NetworkList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Network, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkList.
func (in *NetworkList) DeepCopy() *NetworkList {
	if in == nil {
		return nil
	}
	out := new(NetworkList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NetworkList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkResourceSpec) DeepCopyInto(out *NetworkResourceSpec) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(OpenStackName)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(OpenStackDescription)
		**out = **in
	}
	if in.AdminStateUp != nil {
		in, out := &in.AdminStateUp, &out.AdminStateUp
		*out = new(bool)
		**out = **in
	}
	if in.DNSDomain != nil {
		in, out := &in.DNSDomain, &out.DNSDomain
		*out = new(DNSDomain)
		**out = **in
	}
	if in.MTU != nil {
		in, out := &in.MTU, &out.MTU
		*out = new(MTU)
		**out = **in
	}
	if in.PortSecurityEnabled != nil {
		in, out := &in.PortSecurityEnabled, &out.PortSecurityEnabled
		*out = new(bool)
		**out = **in
	}
	if in.External != nil {
		in, out := &in.External, &out.External
		*out = new(bool)
		**out = **in
	}
	if in.Shared != nil {
		in, out := &in.Shared, &out.Shared
		*out = new(bool)
		**out = **in
	}
	if in.VLANTransparent != nil {
		in, out := &in.VLANTransparent, &out.VLANTransparent
		*out = new(bool)
		**out = **in
	}
	if in.AvailabilityZoneHints != nil {
		in, out := &in.AvailabilityZoneHints, &out.AvailabilityZoneHints
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.IsDefault != nil {
		in, out := &in.IsDefault, &out.IsDefault
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkResourceSpec.
func (in *NetworkResourceSpec) DeepCopy() *NetworkResourceSpec {
	if in == nil {
		return nil
	}
	out := new(NetworkResourceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkResourceStatus) DeepCopyInto(out *NetworkResourceStatus) {
	*out = *in
	if in.AvailabilityZoneHints != nil {
		in, out := &in.AvailabilityZoneHints, &out.AvailabilityZoneHints
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.AvailabilityZones != nil {
		in, out := &in.AvailabilityZones, &out.AvailabilityZones
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.L2Adjacency != nil {
		in, out := &in.L2Adjacency, &out.L2Adjacency
		*out = new(bool)
		**out = **in
	}
	if in.PortSecurityEnabled != nil {
		in, out := &in.PortSecurityEnabled, &out.PortSecurityEnabled
		*out = new(bool)
		**out = **in
	}
	if in.Provider != nil {
		in, out := &in.Provider, &out.Provider
		*out = new(ProviderProperties)
		(*in).DeepCopyInto(*out)
	}
	if in.Subnets != nil {
		in, out := &in.Subnets, &out.Subnets
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.VLANTransparent != nil {
		in, out := &in.VLANTransparent, &out.VLANTransparent
		*out = new(bool)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.NeutronStatusMetadata.DeepCopyInto(&out.NeutronStatusMetadata)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkResourceStatus.
func (in *NetworkResourceStatus) DeepCopy() *NetworkResourceStatus {
	if in == nil {
		return nil
	}
	out := new(NetworkResourceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkSpec) DeepCopyInto(out *NetworkSpec) {
	*out = *in
	if in.Import != nil {
		in, out := &in.Import, &out.Import
		*out = new(NetworkImport)
		(*in).DeepCopyInto(*out)
	}
	if in.Resource != nil {
		in, out := &in.Resource, &out.Resource
		*out = new(NetworkResourceSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.ManagedOptions != nil {
		in, out := &in.ManagedOptions, &out.ManagedOptions
		*out = new(ManagedOptions)
		**out = **in
	}
	out.CloudCredentialsRef = in.CloudCredentialsRef
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkSpec.
func (in *NetworkSpec) DeepCopy() *NetworkSpec {
	if in == nil {
		return nil
	}
	out := new(NetworkSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkStatus) DeepCopyInto(out *NetworkStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Resource != nil {
		in, out := &in.Resource, &out.Resource
		*out = new(NetworkResourceStatus)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkStatus.
func (in *NetworkStatus) DeepCopy() *NetworkStatus {
	if in == nil {
		return nil
	}
	out := new(NetworkStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NeutronStatusMetadata) DeepCopyInto(out *NeutronStatusMetadata) {
	*out = *in
	if in.CreatedAt != nil {
		in, out := &in.CreatedAt, &out.CreatedAt
		*out = (*in).DeepCopy()
	}
	if in.UpdatedAt != nil {
		in, out := &in.UpdatedAt, &out.UpdatedAt
		*out = (*in).DeepCopy()
	}
	if in.RevisionNumber != nil {
		in, out := &in.RevisionNumber, &out.RevisionNumber
		*out = new(NeutronRevisionNumber)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NeutronStatusMetadata.
func (in *NeutronStatusMetadata) DeepCopy() *NeutronStatusMetadata {
	if in == nil {
		return nil
	}
	out := new(NeutronStatusMetadata)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProviderProperties) DeepCopyInto(out *ProviderProperties) {
	*out = *in
	if in.NetworkType != nil {
		in, out := &in.NetworkType, &out.NetworkType
		*out = new(ProviderNetworkType)
		**out = **in
	}
	if in.PhysicalNetwork != nil {
		in, out := &in.PhysicalNetwork, &out.PhysicalNetwork
		*out = new(PhysicalNetwork)
		**out = **in
	}
	if in.SegmentationID != nil {
		in, out := &in.SegmentationID, &out.SegmentationID
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProviderProperties.
func (in *ProviderProperties) DeepCopy() *ProviderProperties {
	if in == nil {
		return nil
	}
	out := new(ProviderProperties)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Subnet) DeepCopyInto(out *Subnet) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Subnet.
func (in *Subnet) DeepCopy() *Subnet {
	if in == nil {
		return nil
	}
	out := new(Subnet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Subnet) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubnetFilter) DeepCopyInto(out *SubnetFilter) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(OpenStackName)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(OpenStackDescription)
		**out = **in
	}
	if in.ProjectID != nil {
		in, out := &in.ProjectID, &out.ProjectID
		*out = new(UUID)
		**out = **in
	}
	if in.IPVersion != nil {
		in, out := &in.IPVersion, &out.IPVersion
		*out = new(IPVersion)
		**out = **in
	}
	if in.GatewayIP != nil {
		in, out := &in.GatewayIP, &out.GatewayIP
		*out = new(IPvAny)
		**out = **in
	}
	if in.CIDR != nil {
		in, out := &in.CIDR, &out.CIDR
		*out = new(CIDR)
		**out = **in
	}
	if in.IPv6 != nil {
		in, out := &in.IPv6, &out.IPv6
		*out = new(IPv6Options)
		(*in).DeepCopyInto(*out)
	}
	in.FilterByNeutronTags.DeepCopyInto(&out.FilterByNeutronTags)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubnetFilter.
func (in *SubnetFilter) DeepCopy() *SubnetFilter {
	if in == nil {
		return nil
	}
	out := new(SubnetFilter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubnetGateway) DeepCopyInto(out *SubnetGateway) {
	*out = *in
	if in.IP != nil {
		in, out := &in.IP, &out.IP
		*out = new(IPvAny)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubnetGateway.
func (in *SubnetGateway) DeepCopy() *SubnetGateway {
	if in == nil {
		return nil
	}
	out := new(SubnetGateway)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubnetImport) DeepCopyInto(out *SubnetImport) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Filter != nil {
		in, out := &in.Filter, &out.Filter
		*out = new(SubnetFilter)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubnetImport.
func (in *SubnetImport) DeepCopy() *SubnetImport {
	if in == nil {
		return nil
	}
	out := new(SubnetImport)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubnetList) DeepCopyInto(out *SubnetList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Subnet, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubnetList.
func (in *SubnetList) DeepCopy() *SubnetList {
	if in == nil {
		return nil
	}
	out := new(SubnetList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SubnetList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubnetRefs) DeepCopyInto(out *SubnetRefs) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubnetRefs.
func (in *SubnetRefs) DeepCopy() *SubnetRefs {
	if in == nil {
		return nil
	}
	out := new(SubnetRefs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubnetResourceSpec) DeepCopyInto(out *SubnetResourceSpec) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(OpenStackName)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(OpenStackDescription)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]NeutronTag, len(*in))
		copy(*out, *in)
	}
	if in.ProjectID != nil {
		in, out := &in.ProjectID, &out.ProjectID
		*out = new(UUID)
		**out = **in
	}
	if in.AllocationPools != nil {
		in, out := &in.AllocationPools, &out.AllocationPools
		*out = make([]AllocationPool, len(*in))
		copy(*out, *in)
	}
	if in.Gateway != nil {
		in, out := &in.Gateway, &out.Gateway
		*out = new(SubnetGateway)
		(*in).DeepCopyInto(*out)
	}
	if in.EnableDHCP != nil {
		in, out := &in.EnableDHCP, &out.EnableDHCP
		*out = new(bool)
		**out = **in
	}
	if in.DNSNameservers != nil {
		in, out := &in.DNSNameservers, &out.DNSNameservers
		*out = make([]IPvAny, len(*in))
		copy(*out, *in)
	}
	if in.DNSPublishFixedIP != nil {
		in, out := &in.DNSPublishFixedIP, &out.DNSPublishFixedIP
		*out = new(bool)
		**out = **in
	}
	if in.HostRoutes != nil {
		in, out := &in.HostRoutes, &out.HostRoutes
		*out = make([]HostRoute, len(*in))
		copy(*out, *in)
	}
	if in.IPv6 != nil {
		in, out := &in.IPv6, &out.IPv6
		*out = new(IPv6Options)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubnetResourceSpec.
func (in *SubnetResourceSpec) DeepCopy() *SubnetResourceSpec {
	if in == nil {
		return nil
	}
	out := new(SubnetResourceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubnetResourceStatus) DeepCopyInto(out *SubnetResourceStatus) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(OpenStackDescription)
		**out = **in
	}
	if in.GatewayIP != nil {
		in, out := &in.GatewayIP, &out.GatewayIP
		*out = new(IPvAny)
		**out = **in
	}
	if in.DNSNameservers != nil {
		in, out := &in.DNSNameservers, &out.DNSNameservers
		*out = make([]IPvAny, len(*in))
		copy(*out, *in)
	}
	if in.AllocationPools != nil {
		in, out := &in.AllocationPools, &out.AllocationPools
		*out = make([]AllocationPool, len(*in))
		copy(*out, *in)
	}
	if in.HostRoutes != nil {
		in, out := &in.HostRoutes, &out.HostRoutes
		*out = make([]HostRoute, len(*in))
		copy(*out, *in)
	}
	if in.IPv6AddressMode != nil {
		in, out := &in.IPv6AddressMode, &out.IPv6AddressMode
		*out = new(IPv6AddressMode)
		**out = **in
	}
	if in.IPv6RAMode != nil {
		in, out := &in.IPv6RAMode, &out.IPv6RAMode
		*out = new(IPv6RAMode)
		**out = **in
	}
	if in.SubnetPoolID != nil {
		in, out := &in.SubnetPoolID, &out.SubnetPoolID
		*out = new(UUID)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]NeutronTag, len(*in))
		copy(*out, *in)
	}
	in.NeutronStatusMetadata.DeepCopyInto(&out.NeutronStatusMetadata)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubnetResourceStatus.
func (in *SubnetResourceStatus) DeepCopy() *SubnetResourceStatus {
	if in == nil {
		return nil
	}
	out := new(SubnetResourceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubnetSpec) DeepCopyInto(out *SubnetSpec) {
	*out = *in
	out.SubnetRefs = in.SubnetRefs
	if in.Import != nil {
		in, out := &in.Import, &out.Import
		*out = new(SubnetImport)
		(*in).DeepCopyInto(*out)
	}
	if in.Resource != nil {
		in, out := &in.Resource, &out.Resource
		*out = new(SubnetResourceSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.ManagedOptions != nil {
		in, out := &in.ManagedOptions, &out.ManagedOptions
		*out = new(ManagedOptions)
		**out = **in
	}
	out.CloudCredentialsRef = in.CloudCredentialsRef
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubnetSpec.
func (in *SubnetSpec) DeepCopy() *SubnetSpec {
	if in == nil {
		return nil
	}
	out := new(SubnetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubnetStatus) DeepCopyInto(out *SubnetStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Resource != nil {
		in, out := &in.Resource, &out.Resource
		*out = new(SubnetResourceStatus)
		(*in).DeepCopyInto(*out)
	}
	in.SubnetStatusExtra.DeepCopyInto(&out.SubnetStatusExtra)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubnetStatus.
func (in *SubnetStatus) DeepCopy() *SubnetStatus {
	if in == nil {
		return nil
	}
	out := new(SubnetStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubnetStatusExtra) DeepCopyInto(out *SubnetStatusExtra) {
	*out = *in
	if in.NetworkID != nil {
		in, out := &in.NetworkID, &out.NetworkID
		*out = new(UUID)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubnetStatusExtra.
func (in *SubnetStatusExtra) DeepCopy() *SubnetStatusExtra {
	if in == nil {
		return nil
	}
	out := new(SubnetStatusExtra)
	in.DeepCopyInto(out)
	return out
}
