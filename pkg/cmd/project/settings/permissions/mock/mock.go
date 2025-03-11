// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/cmd/project/settings/permissions/permissionInterface.go

// Package mock_permissions is a generated GoMock package.
package mock_permissions

import (
	reflect "reflect"

	api "github.com/berrybytes/zocli/api"
	gomock "github.com/golang/mock/gomock"
	cobra "github.com/spf13/cobra"
)

// MockInterface is a mock of Interface interface.
type MockInterface struct {
	ctrl     *gomock.Controller
	recorder *MockInterfaceMockRecorder
}

// MockInterfaceMockRecorder is the mock recorder for MockInterface.
type MockInterfaceMockRecorder struct {
	mock *MockInterface
}

// NewMockInterface creates a new mock instance.
func NewMockInterface(ctrl *gomock.Controller) *MockInterface {
	mock := &MockInterface{ctrl: ctrl}
	mock.recorder = &MockInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInterface) EXPECT() *MockInterfaceMockRecorder {
	return m.recorder
}

// DeletePermission mocks base method.
func (m *MockInterface) DeletePermission(arg0 *cobra.Command, arg1 []string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DeletePermission", arg0, arg1)
}

// DeletePermission indicates an expected call of DeletePermission.
func (mr *MockInterfaceMockRecorder) DeletePermission(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePermission", reflect.TypeOf((*MockInterface)(nil).DeletePermission), arg0, arg1)
}

// GetAllRoles mocks base method.
func (m *MockInterface) GetAllRoles() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "GetAllRoles")
}

// GetAllRoles indicates an expected call of GetAllRoles.
func (mr *MockInterfaceMockRecorder) GetAllRoles() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllRoles", reflect.TypeOf((*MockInterface)(nil).GetAllRoles))
}

// GetPermission mocks base method.
func (m *MockInterface) GetPermission() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "GetPermission")
}

// GetPermission indicates an expected call of GetPermission.
func (mr *MockInterfaceMockRecorder) GetPermission() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPermission", reflect.TypeOf((*MockInterface)(nil).GetPermission))
}

// GetProjectDetail mocks base method.
func (m *MockInterface) GetProjectDetail(id string) *api.Project {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProjectDetail", id)
	ret0, _ := ret[0].(*api.Project)
	return ret0
}

// GetProjectDetail indicates an expected call of GetProjectDetail.
func (mr *MockInterfaceMockRecorder) GetProjectDetail(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProjectDetail", reflect.TypeOf((*MockInterface)(nil).GetProjectDetail), id)
}

// GetProjectDetailByName mocks base method.
func (m *MockInterface) GetProjectDetailByName() *api.Project {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProjectDetailByName")
	ret0, _ := ret[0].(*api.Project)
	return ret0
}

// GetProjectDetailByName indicates an expected call of GetProjectDetailByName.
func (mr *MockInterfaceMockRecorder) GetProjectDetailByName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProjectDetailByName", reflect.TypeOf((*MockInterface)(nil).GetProjectDetailByName))
}

// ListPermissions mocks base method.
func (m *MockInterface) ListPermissions(arg0 *cobra.Command, arg1 []string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ListPermissions", arg0, arg1)
}

// ListPermissions indicates an expected call of ListPermissions.
func (mr *MockInterfaceMockRecorder) ListPermissions(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPermissions", reflect.TypeOf((*MockInterface)(nil).ListPermissions), arg0, arg1)
}

// PrintPermissions mocks base method.
func (m *MockInterface) PrintPermissions() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "PrintPermissions")
}

// PrintPermissions indicates an expected call of PrintPermissions.
func (mr *MockInterfaceMockRecorder) PrintPermissions() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrintPermissions", reflect.TypeOf((*MockInterface)(nil).PrintPermissions))
}

// UpdatePermissions mocks base method.
func (m *MockInterface) UpdatePermissions(arg0 *cobra.Command, arg1 []string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdatePermissions", arg0, arg1)
}

// UpdatePermissions indicates an expected call of UpdatePermissions.
func (mr *MockInterfaceMockRecorder) UpdatePermissions(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePermissions", reflect.TypeOf((*MockInterface)(nil).UpdatePermissions), arg0, arg1)
}
