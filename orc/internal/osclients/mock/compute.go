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
// Code generated by MockGen. DO NOT EDIT.
// Source: ../compute.go
//
// Generated by this command:
//
//	mockgen -package mock -destination=compute.go -source=../compute.go github.com/k-orc/openstack-resource-controller/internal/osclients/mock ComputeClient
//

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	attachinterfaces "github.com/gophercloud/gophercloud/v2/openstack/compute/v2/attachinterfaces"
	availabilityzones "github.com/gophercloud/gophercloud/v2/openstack/compute/v2/availabilityzones"
	flavors "github.com/gophercloud/gophercloud/v2/openstack/compute/v2/flavors"
	servergroups "github.com/gophercloud/gophercloud/v2/openstack/compute/v2/servergroups"
	servers "github.com/gophercloud/gophercloud/v2/openstack/compute/v2/servers"
	osclients "github.com/k-orc/openstack-resource-controller/internal/osclients"
	gomock "go.uber.org/mock/gomock"
)

// MockComputeClient is a mock of ComputeClient interface.
type MockComputeClient struct {
	ctrl     *gomock.Controller
	recorder *MockComputeClientMockRecorder
}

// MockComputeClientMockRecorder is the mock recorder for MockComputeClient.
type MockComputeClientMockRecorder struct {
	mock *MockComputeClient
}

// NewMockComputeClient creates a new mock instance.
func NewMockComputeClient(ctrl *gomock.Controller) *MockComputeClient {
	mock := &MockComputeClient{ctrl: ctrl}
	mock.recorder = &MockComputeClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockComputeClient) EXPECT() *MockComputeClientMockRecorder {
	return m.recorder
}

// CreateFlavor mocks base method.
func (m *MockComputeClient) CreateFlavor(ctx context.Context, opts flavors.CreateOptsBuilder) (*flavors.Flavor, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateFlavor", ctx, opts)
	ret0, _ := ret[0].(*flavors.Flavor)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateFlavor indicates an expected call of CreateFlavor.
func (mr *MockComputeClientMockRecorder) CreateFlavor(ctx, opts any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFlavor", reflect.TypeOf((*MockComputeClient)(nil).CreateFlavor), ctx, opts)
}

// CreateServer mocks base method.
func (m *MockComputeClient) CreateServer(createOpts servers.CreateOptsBuilder, schedulerHints servers.SchedulerHintOptsBuilder) (*servers.Server, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateServer", createOpts, schedulerHints)
	ret0, _ := ret[0].(*servers.Server)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateServer indicates an expected call of CreateServer.
func (mr *MockComputeClientMockRecorder) CreateServer(createOpts, schedulerHints any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateServer", reflect.TypeOf((*MockComputeClient)(nil).CreateServer), createOpts, schedulerHints)
}

// DeleteAttachedInterface mocks base method.
func (m *MockComputeClient) DeleteAttachedInterface(serverID, portID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAttachedInterface", serverID, portID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAttachedInterface indicates an expected call of DeleteAttachedInterface.
func (mr *MockComputeClientMockRecorder) DeleteAttachedInterface(serverID, portID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAttachedInterface", reflect.TypeOf((*MockComputeClient)(nil).DeleteAttachedInterface), serverID, portID)
}

// DeleteFlavor mocks base method.
func (m *MockComputeClient) DeleteFlavor(ctx context.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteFlavor", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteFlavor indicates an expected call of DeleteFlavor.
func (mr *MockComputeClientMockRecorder) DeleteFlavor(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFlavor", reflect.TypeOf((*MockComputeClient)(nil).DeleteFlavor), ctx, id)
}

// DeleteServer mocks base method.
func (m *MockComputeClient) DeleteServer(serverID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteServer", serverID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteServer indicates an expected call of DeleteServer.
func (mr *MockComputeClientMockRecorder) DeleteServer(serverID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteServer", reflect.TypeOf((*MockComputeClient)(nil).DeleteServer), serverID)
}

// GetFlavor mocks base method.
func (m *MockComputeClient) GetFlavor(ctx context.Context, id string) (*flavors.Flavor, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFlavor", ctx, id)
	ret0, _ := ret[0].(*flavors.Flavor)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFlavor indicates an expected call of GetFlavor.
func (mr *MockComputeClientMockRecorder) GetFlavor(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFlavor", reflect.TypeOf((*MockComputeClient)(nil).GetFlavor), ctx, id)
}

// GetFlavorFromName mocks base method.
func (m *MockComputeClient) GetFlavorFromName(flavor string) (*flavors.Flavor, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFlavorFromName", flavor)
	ret0, _ := ret[0].(*flavors.Flavor)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFlavorFromName indicates an expected call of GetFlavorFromName.
func (mr *MockComputeClientMockRecorder) GetFlavorFromName(flavor any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFlavorFromName", reflect.TypeOf((*MockComputeClient)(nil).GetFlavorFromName), flavor)
}

// GetServer mocks base method.
func (m *MockComputeClient) GetServer(serverID string) (*servers.Server, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetServer", serverID)
	ret0, _ := ret[0].(*servers.Server)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetServer indicates an expected call of GetServer.
func (mr *MockComputeClientMockRecorder) GetServer(serverID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetServer", reflect.TypeOf((*MockComputeClient)(nil).GetServer), serverID)
}

// ListAttachedInterfaces mocks base method.
func (m *MockComputeClient) ListAttachedInterfaces(serverID string) ([]attachinterfaces.Interface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAttachedInterfaces", serverID)
	ret0, _ := ret[0].([]attachinterfaces.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAttachedInterfaces indicates an expected call of ListAttachedInterfaces.
func (mr *MockComputeClientMockRecorder) ListAttachedInterfaces(serverID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAttachedInterfaces", reflect.TypeOf((*MockComputeClient)(nil).ListAttachedInterfaces), serverID)
}

// ListAvailabilityZones mocks base method.
func (m *MockComputeClient) ListAvailabilityZones() ([]availabilityzones.AvailabilityZone, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAvailabilityZones")
	ret0, _ := ret[0].([]availabilityzones.AvailabilityZone)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAvailabilityZones indicates an expected call of ListAvailabilityZones.
func (mr *MockComputeClientMockRecorder) ListAvailabilityZones() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAvailabilityZones", reflect.TypeOf((*MockComputeClient)(nil).ListAvailabilityZones))
}

// ListFlavors mocks base method.
func (m *MockComputeClient) ListFlavors(ctx context.Context, listOpts flavors.ListOptsBuilder) <-chan osclients.FlavorResult {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListFlavors", ctx, listOpts)
	ret0, _ := ret[0].(<-chan osclients.FlavorResult)
	return ret0
}

// ListFlavors indicates an expected call of ListFlavors.
func (mr *MockComputeClientMockRecorder) ListFlavors(ctx, listOpts any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListFlavors", reflect.TypeOf((*MockComputeClient)(nil).ListFlavors), ctx, listOpts)
}

// ListServerGroups mocks base method.
func (m *MockComputeClient) ListServerGroups() ([]servergroups.ServerGroup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListServerGroups")
	ret0, _ := ret[0].([]servergroups.ServerGroup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListServerGroups indicates an expected call of ListServerGroups.
func (mr *MockComputeClientMockRecorder) ListServerGroups() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListServerGroups", reflect.TypeOf((*MockComputeClient)(nil).ListServerGroups))
}

// ListServers mocks base method.
func (m *MockComputeClient) ListServers(listOpts servers.ListOptsBuilder) ([]servers.Server, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListServers", listOpts)
	ret0, _ := ret[0].([]servers.Server)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListServers indicates an expected call of ListServers.
func (mr *MockComputeClientMockRecorder) ListServers(listOpts any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListServers", reflect.TypeOf((*MockComputeClient)(nil).ListServers), listOpts)
}
