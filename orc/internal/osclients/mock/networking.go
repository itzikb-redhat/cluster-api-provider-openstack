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
// Source: ../networking.go
//
// Generated by this command:
//
//	mockgen -package mock -destination=networking.go -source=../networking.go github.com/k-orc/openstack-resource-controller/internal/osclients/mock NetworkClient
//

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	extensions "github.com/gophercloud/gophercloud/v2/openstack/networking/v2/extensions"
	attributestags "github.com/gophercloud/gophercloud/v2/openstack/networking/v2/extensions/attributestags"
	floatingips "github.com/gophercloud/gophercloud/v2/openstack/networking/v2/extensions/layer3/floatingips"
	routers "github.com/gophercloud/gophercloud/v2/openstack/networking/v2/extensions/layer3/routers"
	groups "github.com/gophercloud/gophercloud/v2/openstack/networking/v2/extensions/security/groups"
	rules "github.com/gophercloud/gophercloud/v2/openstack/networking/v2/extensions/security/rules"
	trunks "github.com/gophercloud/gophercloud/v2/openstack/networking/v2/extensions/trunks"
	networks "github.com/gophercloud/gophercloud/v2/openstack/networking/v2/networks"
	ports "github.com/gophercloud/gophercloud/v2/openstack/networking/v2/ports"
	subnets "github.com/gophercloud/gophercloud/v2/openstack/networking/v2/subnets"
	pagination "github.com/gophercloud/gophercloud/v2/pagination"
	gomock "go.uber.org/mock/gomock"
)

// MockNetworkClient is a mock of NetworkClient interface.
type MockNetworkClient struct {
	ctrl     *gomock.Controller
	recorder *MockNetworkClientMockRecorder
}

// MockNetworkClientMockRecorder is the mock recorder for MockNetworkClient.
type MockNetworkClientMockRecorder struct {
	mock *MockNetworkClient
}

// NewMockNetworkClient creates a new mock instance.
func NewMockNetworkClient(ctrl *gomock.Controller) *MockNetworkClient {
	mock := &MockNetworkClient{ctrl: ctrl}
	mock.recorder = &MockNetworkClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNetworkClient) EXPECT() *MockNetworkClientMockRecorder {
	return m.recorder
}

// AddRouterInterface mocks base method.
func (m *MockNetworkClient) AddRouterInterface(ctx context.Context, id string, opts routers.AddInterfaceOptsBuilder) (*routers.InterfaceInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddRouterInterface", ctx, id, opts)
	ret0, _ := ret[0].(*routers.InterfaceInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddRouterInterface indicates an expected call of AddRouterInterface.
func (mr *MockNetworkClientMockRecorder) AddRouterInterface(ctx, id, opts any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddRouterInterface", reflect.TypeOf((*MockNetworkClient)(nil).AddRouterInterface), ctx, id, opts)
}

// CreateFloatingIP mocks base method.
func (m *MockNetworkClient) CreateFloatingIP(opts floatingips.CreateOptsBuilder) (*floatingips.FloatingIP, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateFloatingIP", opts)
	ret0, _ := ret[0].(*floatingips.FloatingIP)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateFloatingIP indicates an expected call of CreateFloatingIP.
func (mr *MockNetworkClientMockRecorder) CreateFloatingIP(opts any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFloatingIP", reflect.TypeOf((*MockNetworkClient)(nil).CreateFloatingIP), opts)
}

// CreateNetwork mocks base method.
func (m *MockNetworkClient) CreateNetwork(ctx context.Context, opts networks.CreateOptsBuilder) networks.CreateResult {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateNetwork", ctx, opts)
	ret0, _ := ret[0].(networks.CreateResult)
	return ret0
}

// CreateNetwork indicates an expected call of CreateNetwork.
func (mr *MockNetworkClientMockRecorder) CreateNetwork(ctx, opts any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateNetwork", reflect.TypeOf((*MockNetworkClient)(nil).CreateNetwork), ctx, opts)
}

// CreatePort mocks base method.
func (m *MockNetworkClient) CreatePort(ctx context.Context, opts ports.CreateOptsBuilder) (*ports.Port, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePort", ctx, opts)
	ret0, _ := ret[0].(*ports.Port)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePort indicates an expected call of CreatePort.
func (mr *MockNetworkClientMockRecorder) CreatePort(ctx, opts any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePort", reflect.TypeOf((*MockNetworkClient)(nil).CreatePort), ctx, opts)
}

// CreateRouter mocks base method.
func (m *MockNetworkClient) CreateRouter(ctx context.Context, opts routers.CreateOptsBuilder) (*routers.Router, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRouter", ctx, opts)
	ret0, _ := ret[0].(*routers.Router)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateRouter indicates an expected call of CreateRouter.
func (mr *MockNetworkClientMockRecorder) CreateRouter(ctx, opts any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRouter", reflect.TypeOf((*MockNetworkClient)(nil).CreateRouter), ctx, opts)
}

// CreateSecGroup mocks base method.
func (m *MockNetworkClient) CreateSecGroup(ctx context.Context, opts groups.CreateOptsBuilder) (*groups.SecGroup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSecGroup", ctx, opts)
	ret0, _ := ret[0].(*groups.SecGroup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSecGroup indicates an expected call of CreateSecGroup.
func (mr *MockNetworkClientMockRecorder) CreateSecGroup(ctx, opts any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSecGroup", reflect.TypeOf((*MockNetworkClient)(nil).CreateSecGroup), ctx, opts)
}

// CreateSecGroupRule mocks base method.
func (m *MockNetworkClient) CreateSecGroupRule(opts rules.CreateOptsBuilder) (*rules.SecGroupRule, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSecGroupRule", opts)
	ret0, _ := ret[0].(*rules.SecGroupRule)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSecGroupRule indicates an expected call of CreateSecGroupRule.
func (mr *MockNetworkClientMockRecorder) CreateSecGroupRule(opts any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSecGroupRule", reflect.TypeOf((*MockNetworkClient)(nil).CreateSecGroupRule), opts)
}

// CreateSubnet mocks base method.
func (m *MockNetworkClient) CreateSubnet(ctx context.Context, opts subnets.CreateOptsBuilder) (*subnets.Subnet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSubnet", ctx, opts)
	ret0, _ := ret[0].(*subnets.Subnet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSubnet indicates an expected call of CreateSubnet.
func (mr *MockNetworkClientMockRecorder) CreateSubnet(ctx, opts any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSubnet", reflect.TypeOf((*MockNetworkClient)(nil).CreateSubnet), ctx, opts)
}

// CreateTrunk mocks base method.
func (m *MockNetworkClient) CreateTrunk(opts trunks.CreateOptsBuilder) (*trunks.Trunk, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTrunk", opts)
	ret0, _ := ret[0].(*trunks.Trunk)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTrunk indicates an expected call of CreateTrunk.
func (mr *MockNetworkClientMockRecorder) CreateTrunk(opts any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTrunk", reflect.TypeOf((*MockNetworkClient)(nil).CreateTrunk), opts)
}

// DeleteFloatingIP mocks base method.
func (m *MockNetworkClient) DeleteFloatingIP(id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteFloatingIP", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteFloatingIP indicates an expected call of DeleteFloatingIP.
func (mr *MockNetworkClientMockRecorder) DeleteFloatingIP(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFloatingIP", reflect.TypeOf((*MockNetworkClient)(nil).DeleteFloatingIP), id)
}

// DeleteNetwork mocks base method.
func (m *MockNetworkClient) DeleteNetwork(ctx context.Context, id string) networks.DeleteResult {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteNetwork", ctx, id)
	ret0, _ := ret[0].(networks.DeleteResult)
	return ret0
}

// DeleteNetwork indicates an expected call of DeleteNetwork.
func (mr *MockNetworkClientMockRecorder) DeleteNetwork(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteNetwork", reflect.TypeOf((*MockNetworkClient)(nil).DeleteNetwork), ctx, id)
}

// DeletePort mocks base method.
func (m *MockNetworkClient) DeletePort(ctx context.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePort", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePort indicates an expected call of DeletePort.
func (mr *MockNetworkClientMockRecorder) DeletePort(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePort", reflect.TypeOf((*MockNetworkClient)(nil).DeletePort), ctx, id)
}

// DeleteRouter mocks base method.
func (m *MockNetworkClient) DeleteRouter(ctx context.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteRouter", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteRouter indicates an expected call of DeleteRouter.
func (mr *MockNetworkClientMockRecorder) DeleteRouter(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRouter", reflect.TypeOf((*MockNetworkClient)(nil).DeleteRouter), ctx, id)
}

// DeleteSecGroup mocks base method.
func (m *MockNetworkClient) DeleteSecGroup(ctx context.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSecGroup", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSecGroup indicates an expected call of DeleteSecGroup.
func (mr *MockNetworkClientMockRecorder) DeleteSecGroup(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSecGroup", reflect.TypeOf((*MockNetworkClient)(nil).DeleteSecGroup), ctx, id)
}

// DeleteSecGroupRule mocks base method.
func (m *MockNetworkClient) DeleteSecGroupRule(id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSecGroupRule", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSecGroupRule indicates an expected call of DeleteSecGroupRule.
func (mr *MockNetworkClientMockRecorder) DeleteSecGroupRule(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSecGroupRule", reflect.TypeOf((*MockNetworkClient)(nil).DeleteSecGroupRule), id)
}

// DeleteSubnet mocks base method.
func (m *MockNetworkClient) DeleteSubnet(ctx context.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSubnet", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSubnet indicates an expected call of DeleteSubnet.
func (mr *MockNetworkClientMockRecorder) DeleteSubnet(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSubnet", reflect.TypeOf((*MockNetworkClient)(nil).DeleteSubnet), ctx, id)
}

// DeleteTrunk mocks base method.
func (m *MockNetworkClient) DeleteTrunk(id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTrunk", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTrunk indicates an expected call of DeleteTrunk.
func (mr *MockNetworkClientMockRecorder) DeleteTrunk(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTrunk", reflect.TypeOf((*MockNetworkClient)(nil).DeleteTrunk), id)
}

// GetFloatingIP mocks base method.
func (m *MockNetworkClient) GetFloatingIP(id string) (*floatingips.FloatingIP, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFloatingIP", id)
	ret0, _ := ret[0].(*floatingips.FloatingIP)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFloatingIP indicates an expected call of GetFloatingIP.
func (mr *MockNetworkClientMockRecorder) GetFloatingIP(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFloatingIP", reflect.TypeOf((*MockNetworkClient)(nil).GetFloatingIP), id)
}

// GetNetwork mocks base method.
func (m *MockNetworkClient) GetNetwork(ctx context.Context, id string) networks.GetResult {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNetwork", ctx, id)
	ret0, _ := ret[0].(networks.GetResult)
	return ret0
}

// GetNetwork indicates an expected call of GetNetwork.
func (mr *MockNetworkClientMockRecorder) GetNetwork(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNetwork", reflect.TypeOf((*MockNetworkClient)(nil).GetNetwork), ctx, id)
}

// GetPort mocks base method.
func (m *MockNetworkClient) GetPort(ctx context.Context, id string) (*ports.Port, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPort", ctx, id)
	ret0, _ := ret[0].(*ports.Port)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPort indicates an expected call of GetPort.
func (mr *MockNetworkClientMockRecorder) GetPort(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPort", reflect.TypeOf((*MockNetworkClient)(nil).GetPort), ctx, id)
}

// GetRouter mocks base method.
func (m *MockNetworkClient) GetRouter(ctx context.Context, id string) (*routers.Router, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRouter", ctx, id)
	ret0, _ := ret[0].(*routers.Router)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRouter indicates an expected call of GetRouter.
func (mr *MockNetworkClientMockRecorder) GetRouter(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRouter", reflect.TypeOf((*MockNetworkClient)(nil).GetRouter), ctx, id)
}

// GetSecGroup mocks base method.
func (m *MockNetworkClient) GetSecGroup(ctx context.Context, id string) (*groups.SecGroup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSecGroup", ctx, id)
	ret0, _ := ret[0].(*groups.SecGroup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSecGroup indicates an expected call of GetSecGroup.
func (mr *MockNetworkClientMockRecorder) GetSecGroup(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSecGroup", reflect.TypeOf((*MockNetworkClient)(nil).GetSecGroup), ctx, id)
}

// GetSecGroupRule mocks base method.
func (m *MockNetworkClient) GetSecGroupRule(id string) (*rules.SecGroupRule, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSecGroupRule", id)
	ret0, _ := ret[0].(*rules.SecGroupRule)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSecGroupRule indicates an expected call of GetSecGroupRule.
func (mr *MockNetworkClientMockRecorder) GetSecGroupRule(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSecGroupRule", reflect.TypeOf((*MockNetworkClient)(nil).GetSecGroupRule), id)
}

// GetSubnet mocks base method.
func (m *MockNetworkClient) GetSubnet(ctx context.Context, id string) (*subnets.Subnet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSubnet", ctx, id)
	ret0, _ := ret[0].(*subnets.Subnet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSubnet indicates an expected call of GetSubnet.
func (mr *MockNetworkClientMockRecorder) GetSubnet(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubnet", reflect.TypeOf((*MockNetworkClient)(nil).GetSubnet), ctx, id)
}

// ListExtensions mocks base method.
func (m *MockNetworkClient) ListExtensions() ([]extensions.Extension, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListExtensions")
	ret0, _ := ret[0].([]extensions.Extension)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListExtensions indicates an expected call of ListExtensions.
func (mr *MockNetworkClientMockRecorder) ListExtensions() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListExtensions", reflect.TypeOf((*MockNetworkClient)(nil).ListExtensions))
}

// ListFloatingIP mocks base method.
func (m *MockNetworkClient) ListFloatingIP(opts floatingips.ListOptsBuilder) ([]floatingips.FloatingIP, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListFloatingIP", opts)
	ret0, _ := ret[0].([]floatingips.FloatingIP)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListFloatingIP indicates an expected call of ListFloatingIP.
func (mr *MockNetworkClientMockRecorder) ListFloatingIP(opts any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListFloatingIP", reflect.TypeOf((*MockNetworkClient)(nil).ListFloatingIP), opts)
}

// ListNetwork mocks base method.
func (m *MockNetworkClient) ListNetwork(opts networks.ListOptsBuilder) pagination.Pager {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListNetwork", opts)
	ret0, _ := ret[0].(pagination.Pager)
	return ret0
}

// ListNetwork indicates an expected call of ListNetwork.
func (mr *MockNetworkClientMockRecorder) ListNetwork(opts any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListNetwork", reflect.TypeOf((*MockNetworkClient)(nil).ListNetwork), opts)
}

// ListPort mocks base method.
func (m *MockNetworkClient) ListPort(ctx context.Context, opts ports.ListOptsBuilder) ([]ports.Port, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPort", ctx, opts)
	ret0, _ := ret[0].([]ports.Port)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPort indicates an expected call of ListPort.
func (mr *MockNetworkClientMockRecorder) ListPort(ctx, opts any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPort", reflect.TypeOf((*MockNetworkClient)(nil).ListPort), ctx, opts)
}

// ListRouter mocks base method.
func (m *MockNetworkClient) ListRouter(ctx context.Context, opts routers.ListOpts) ([]routers.Router, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListRouter", ctx, opts)
	ret0, _ := ret[0].([]routers.Router)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListRouter indicates an expected call of ListRouter.
func (mr *MockNetworkClientMockRecorder) ListRouter(ctx, opts any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRouter", reflect.TypeOf((*MockNetworkClient)(nil).ListRouter), ctx, opts)
}

// ListSecGroup mocks base method.
func (m *MockNetworkClient) ListSecGroup(ctx context.Context, opts groups.ListOpts) ([]groups.SecGroup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListSecGroup", ctx, opts)
	ret0, _ := ret[0].([]groups.SecGroup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSecGroup indicates an expected call of ListSecGroup.
func (mr *MockNetworkClientMockRecorder) ListSecGroup(ctx, opts any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSecGroup", reflect.TypeOf((*MockNetworkClient)(nil).ListSecGroup), ctx, opts)
}

// ListSecGroupRule mocks base method.
func (m *MockNetworkClient) ListSecGroupRule(opts rules.ListOpts) ([]rules.SecGroupRule, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListSecGroupRule", opts)
	ret0, _ := ret[0].([]rules.SecGroupRule)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSecGroupRule indicates an expected call of ListSecGroupRule.
func (mr *MockNetworkClientMockRecorder) ListSecGroupRule(opts any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSecGroupRule", reflect.TypeOf((*MockNetworkClient)(nil).ListSecGroupRule), opts)
}

// ListSubnet mocks base method.
func (m *MockNetworkClient) ListSubnet(ctx context.Context, opts subnets.ListOptsBuilder) ([]subnets.Subnet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListSubnet", ctx, opts)
	ret0, _ := ret[0].([]subnets.Subnet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSubnet indicates an expected call of ListSubnet.
func (mr *MockNetworkClientMockRecorder) ListSubnet(ctx, opts any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSubnet", reflect.TypeOf((*MockNetworkClient)(nil).ListSubnet), ctx, opts)
}

// ListTrunk mocks base method.
func (m *MockNetworkClient) ListTrunk(opts trunks.ListOptsBuilder) ([]trunks.Trunk, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTrunk", opts)
	ret0, _ := ret[0].([]trunks.Trunk)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTrunk indicates an expected call of ListTrunk.
func (mr *MockNetworkClientMockRecorder) ListTrunk(opts any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTrunk", reflect.TypeOf((*MockNetworkClient)(nil).ListTrunk), opts)
}

// ListTrunkSubports mocks base method.
func (m *MockNetworkClient) ListTrunkSubports(trunkID string) ([]trunks.Subport, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTrunkSubports", trunkID)
	ret0, _ := ret[0].([]trunks.Subport)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTrunkSubports indicates an expected call of ListTrunkSubports.
func (mr *MockNetworkClientMockRecorder) ListTrunkSubports(trunkID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTrunkSubports", reflect.TypeOf((*MockNetworkClient)(nil).ListTrunkSubports), trunkID)
}

// RemoveRouterInterface mocks base method.
func (m *MockNetworkClient) RemoveRouterInterface(ctx context.Context, id string, opts routers.RemoveInterfaceOptsBuilder) (*routers.InterfaceInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveRouterInterface", ctx, id, opts)
	ret0, _ := ret[0].(*routers.InterfaceInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoveRouterInterface indicates an expected call of RemoveRouterInterface.
func (mr *MockNetworkClientMockRecorder) RemoveRouterInterface(ctx, id, opts any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveRouterInterface", reflect.TypeOf((*MockNetworkClient)(nil).RemoveRouterInterface), ctx, id, opts)
}

// RemoveSubports mocks base method.
func (m *MockNetworkClient) RemoveSubports(id string, opts trunks.RemoveSubportsOpts) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveSubports", id, opts)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveSubports indicates an expected call of RemoveSubports.
func (mr *MockNetworkClientMockRecorder) RemoveSubports(id, opts any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveSubports", reflect.TypeOf((*MockNetworkClient)(nil).RemoveSubports), id, opts)
}

// ReplaceAllAttributesTags mocks base method.
func (m *MockNetworkClient) ReplaceAllAttributesTags(ctx context.Context, resourceType, resourceID string, opts attributestags.ReplaceAllOptsBuilder) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReplaceAllAttributesTags", ctx, resourceType, resourceID, opts)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReplaceAllAttributesTags indicates an expected call of ReplaceAllAttributesTags.
func (mr *MockNetworkClientMockRecorder) ReplaceAllAttributesTags(ctx, resourceType, resourceID, opts any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReplaceAllAttributesTags", reflect.TypeOf((*MockNetworkClient)(nil).ReplaceAllAttributesTags), ctx, resourceType, resourceID, opts)
}

// UpdateFloatingIP mocks base method.
func (m *MockNetworkClient) UpdateFloatingIP(id string, opts floatingips.UpdateOptsBuilder) (*floatingips.FloatingIP, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateFloatingIP", id, opts)
	ret0, _ := ret[0].(*floatingips.FloatingIP)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateFloatingIP indicates an expected call of UpdateFloatingIP.
func (mr *MockNetworkClientMockRecorder) UpdateFloatingIP(id, opts any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateFloatingIP", reflect.TypeOf((*MockNetworkClient)(nil).UpdateFloatingIP), id, opts)
}

// UpdateNetwork mocks base method.
func (m *MockNetworkClient) UpdateNetwork(ctx context.Context, id string, opts networks.UpdateOptsBuilder) networks.UpdateResult {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateNetwork", ctx, id, opts)
	ret0, _ := ret[0].(networks.UpdateResult)
	return ret0
}

// UpdateNetwork indicates an expected call of UpdateNetwork.
func (mr *MockNetworkClientMockRecorder) UpdateNetwork(ctx, id, opts any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateNetwork", reflect.TypeOf((*MockNetworkClient)(nil).UpdateNetwork), ctx, id, opts)
}

// UpdatePort mocks base method.
func (m *MockNetworkClient) UpdatePort(id string, opts ports.UpdateOptsBuilder) (*ports.Port, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePort", id, opts)
	ret0, _ := ret[0].(*ports.Port)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdatePort indicates an expected call of UpdatePort.
func (mr *MockNetworkClientMockRecorder) UpdatePort(id, opts any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePort", reflect.TypeOf((*MockNetworkClient)(nil).UpdatePort), id, opts)
}

// UpdateRouter mocks base method.
func (m *MockNetworkClient) UpdateRouter(ctx context.Context, id string, opts routers.UpdateOptsBuilder) (*routers.Router, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateRouter", ctx, id, opts)
	ret0, _ := ret[0].(*routers.Router)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateRouter indicates an expected call of UpdateRouter.
func (mr *MockNetworkClientMockRecorder) UpdateRouter(ctx, id, opts any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRouter", reflect.TypeOf((*MockNetworkClient)(nil).UpdateRouter), ctx, id, opts)
}

// UpdateSecGroup mocks base method.
func (m *MockNetworkClient) UpdateSecGroup(ctx context.Context, id string, opts groups.UpdateOptsBuilder) (*groups.SecGroup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateSecGroup", ctx, id, opts)
	ret0, _ := ret[0].(*groups.SecGroup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateSecGroup indicates an expected call of UpdateSecGroup.
func (mr *MockNetworkClientMockRecorder) UpdateSecGroup(ctx, id, opts any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSecGroup", reflect.TypeOf((*MockNetworkClient)(nil).UpdateSecGroup), ctx, id, opts)
}

// UpdateSubnet mocks base method.
func (m *MockNetworkClient) UpdateSubnet(ctx context.Context, id string, opts subnets.UpdateOptsBuilder) (*subnets.Subnet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateSubnet", ctx, id, opts)
	ret0, _ := ret[0].(*subnets.Subnet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateSubnet indicates an expected call of UpdateSubnet.
func (mr *MockNetworkClientMockRecorder) UpdateSubnet(ctx, id, opts any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSubnet", reflect.TypeOf((*MockNetworkClient)(nil).UpdateSubnet), ctx, id, opts)
}
