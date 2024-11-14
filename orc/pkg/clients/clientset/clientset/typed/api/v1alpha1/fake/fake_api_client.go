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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/k-orc/openstack-resource-controller/pkg/clients/clientset/clientset/typed/api/v1alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeOpenstackV1alpha1 struct {
	*testing.Fake
}

func (c *FakeOpenstackV1alpha1) Images(namespace string) v1alpha1.ImageInterface {
	return &FakeImages{c, namespace}
}

func (c *FakeOpenstackV1alpha1) Networks(namespace string) v1alpha1.NetworkInterface {
	return &FakeNetworks{c, namespace}
}

func (c *FakeOpenstackV1alpha1) Ports(namespace string) v1alpha1.PortInterface {
	return &FakePorts{c, namespace}
}

func (c *FakeOpenstackV1alpha1) Routers(namespace string) v1alpha1.RouterInterface {
	return &FakeRouters{c, namespace}
}

func (c *FakeOpenstackV1alpha1) Subnets(namespace string) v1alpha1.SubnetInterface {
	return &FakeSubnets{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeOpenstackV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
