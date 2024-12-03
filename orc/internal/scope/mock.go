/*
Copyright 2022 The ORC Authors.

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

package scope

import (
	"context"
	"time"

	"github.com/go-logr/logr"
	"github.com/gophercloud/gophercloud/v2/openstack/identity/v3/tokens"
	"go.uber.org/mock/gomock"
	"sigs.k8s.io/controller-runtime/pkg/client"

	orcv1alpha1 "github.com/k-orc/openstack-resource-controller/api/v1alpha1"

	osclients "github.com/k-orc/openstack-resource-controller/internal/osclients"
	"github.com/k-orc/openstack-resource-controller/internal/osclients/mock"
)

// MockScopeFactory implements both the ScopeFactory and ClientScope interfaces. It can be used in place of the default ProviderScopeFactory
// when we want to use mocked service clients which do not attempt to connect to a running OpenStack cloud.
type MockScopeFactory struct {
	ComputeClient *mock.MockComputeClient
	NetworkClient *mock.MockNetworkClient
	VolumeClient  *mock.MockVolumeClient
	ImageClient   *mock.MockImageClient
	LbClient      *mock.MockLbClient

	clientScopeCreateError error
}

func NewMockScopeFactory(mockCtrl *gomock.Controller) *MockScopeFactory {
	computeClient := mock.NewMockComputeClient(mockCtrl)
	volumeClient := mock.NewMockVolumeClient(mockCtrl)
	imageClient := mock.NewMockImageClient(mockCtrl)
	networkClient := mock.NewMockNetworkClient(mockCtrl)
	lbClient := mock.NewMockLbClient(mockCtrl)

	return &MockScopeFactory{
		ComputeClient: computeClient,
		VolumeClient:  volumeClient,
		ImageClient:   imageClient,
		NetworkClient: networkClient,
		LbClient:      lbClient,
	}
}

func (f *MockScopeFactory) SetClientScopeCreateError(err error) {
	f.clientScopeCreateError = err
}

func (f *MockScopeFactory) NewClientScopeFromObject(_ context.Context, _ client.Client, _ logr.Logger, _ ...orcv1alpha1.CloudCredentialsRefProvider) (Scope, error) {
	if f.clientScopeCreateError != nil {
		return nil, f.clientScopeCreateError
	}
	return f, nil
}

func (f *MockScopeFactory) NewComputeClient() (osclients.ComputeClient, error) {
	return f.ComputeClient, nil
}

func (f *MockScopeFactory) NewVolumeClient() (osclients.VolumeClient, error) {
	return f.VolumeClient, nil
}

func (f *MockScopeFactory) NewImageClient() (osclients.ImageClient, error) {
	return f.ImageClient, nil
}

func (f *MockScopeFactory) NewNetworkClient() (osclients.NetworkClient, error) {
	return f.NetworkClient, nil
}

func (f *MockScopeFactory) NewLbClient() (osclients.LbClient, error) {
	return f.LbClient, nil
}

func (f *MockScopeFactory) ExtractToken() (*tokens.Token, error) {
	return &tokens.Token{ExpiresAt: time.Now().Add(24 * time.Hour)}, nil
}
