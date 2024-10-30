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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha7

import (
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/listers"
	"k8s.io/client-go/tools/cache"
	v1alpha7 "sigs.k8s.io/cluster-api-provider-openstack/api/v1alpha7"
)

// OpenStackMachineTemplateLister helps list OpenStackMachineTemplates.
// All objects returned here must be treated as read-only.
type OpenStackMachineTemplateLister interface {
	// List lists all OpenStackMachineTemplates in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha7.OpenStackMachineTemplate, err error)
	// OpenStackMachineTemplates returns an object that can list and get OpenStackMachineTemplates.
	OpenStackMachineTemplates(namespace string) OpenStackMachineTemplateNamespaceLister
	OpenStackMachineTemplateListerExpansion
}

// openStackMachineTemplateLister implements the OpenStackMachineTemplateLister interface.
type openStackMachineTemplateLister struct {
	listers.ResourceIndexer[*v1alpha7.OpenStackMachineTemplate]
}

// NewOpenStackMachineTemplateLister returns a new OpenStackMachineTemplateLister.
func NewOpenStackMachineTemplateLister(indexer cache.Indexer) OpenStackMachineTemplateLister {
	return &openStackMachineTemplateLister{listers.New[*v1alpha7.OpenStackMachineTemplate](indexer, v1alpha7.Resource("openstackmachinetemplate"))}
}

// OpenStackMachineTemplates returns an object that can list and get OpenStackMachineTemplates.
func (s *openStackMachineTemplateLister) OpenStackMachineTemplates(namespace string) OpenStackMachineTemplateNamespaceLister {
	return openStackMachineTemplateNamespaceLister{listers.NewNamespaced[*v1alpha7.OpenStackMachineTemplate](s.ResourceIndexer, namespace)}
}

// OpenStackMachineTemplateNamespaceLister helps list and get OpenStackMachineTemplates.
// All objects returned here must be treated as read-only.
type OpenStackMachineTemplateNamespaceLister interface {
	// List lists all OpenStackMachineTemplates in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha7.OpenStackMachineTemplate, err error)
	// Get retrieves the OpenStackMachineTemplate from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha7.OpenStackMachineTemplate, error)
	OpenStackMachineTemplateNamespaceListerExpansion
}

// openStackMachineTemplateNamespaceLister implements the OpenStackMachineTemplateNamespaceLister
// interface.
type openStackMachineTemplateNamespaceLister struct {
	listers.ResourceIndexer[*v1alpha7.OpenStackMachineTemplate]
}
