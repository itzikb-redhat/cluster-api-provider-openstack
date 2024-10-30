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

// Code generated by client-gen. DO NOT EDIT.

package v1beta1

import (
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
	v1beta1 "sigs.k8s.io/cluster-api-provider-openstack/api/v1beta1"
	apiv1beta1 "sigs.k8s.io/cluster-api-provider-openstack/pkg/generated/applyconfiguration/api/v1beta1"
	scheme "sigs.k8s.io/cluster-api-provider-openstack/pkg/generated/clientset/clientset/scheme"
)

// OpenStackMachineTemplatesGetter has a method to return a OpenStackMachineTemplateInterface.
// A group's client should implement this interface.
type OpenStackMachineTemplatesGetter interface {
	OpenStackMachineTemplates(namespace string) OpenStackMachineTemplateInterface
}

// OpenStackMachineTemplateInterface has methods to work with OpenStackMachineTemplate resources.
type OpenStackMachineTemplateInterface interface {
	Create(ctx context.Context, openStackMachineTemplate *v1beta1.OpenStackMachineTemplate, opts v1.CreateOptions) (*v1beta1.OpenStackMachineTemplate, error)
	Update(ctx context.Context, openStackMachineTemplate *v1beta1.OpenStackMachineTemplate, opts v1.UpdateOptions) (*v1beta1.OpenStackMachineTemplate, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.OpenStackMachineTemplate, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.OpenStackMachineTemplateList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.OpenStackMachineTemplate, err error)
	Apply(ctx context.Context, openStackMachineTemplate *apiv1beta1.OpenStackMachineTemplateApplyConfiguration, opts v1.ApplyOptions) (result *v1beta1.OpenStackMachineTemplate, err error)
	OpenStackMachineTemplateExpansion
}

// openStackMachineTemplates implements OpenStackMachineTemplateInterface
type openStackMachineTemplates struct {
	*gentype.ClientWithListAndApply[*v1beta1.OpenStackMachineTemplate, *v1beta1.OpenStackMachineTemplateList, *apiv1beta1.OpenStackMachineTemplateApplyConfiguration]
}

// newOpenStackMachineTemplates returns a OpenStackMachineTemplates
func newOpenStackMachineTemplates(c *InfrastructureV1beta1Client, namespace string) *openStackMachineTemplates {
	return &openStackMachineTemplates{
		gentype.NewClientWithListAndApply[*v1beta1.OpenStackMachineTemplate, *v1beta1.OpenStackMachineTemplateList, *apiv1beta1.OpenStackMachineTemplateApplyConfiguration](
			"openstackmachinetemplates",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *v1beta1.OpenStackMachineTemplate { return &v1beta1.OpenStackMachineTemplate{} },
			func() *v1beta1.OpenStackMachineTemplateList { return &v1beta1.OpenStackMachineTemplateList{} }),
	}
}
