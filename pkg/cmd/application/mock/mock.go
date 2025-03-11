// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/cmd/application/app_interface.go

// Package mock_application is a generated GoMock package.
package mock_application

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

// AskDeleteConfirmation mocks base method.
func (m *MockInterface) AskDeleteConfirmation() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AskDeleteConfirmation")
}

// AskDeleteConfirmation indicates an expected call of AskDeleteConfirmation.
func (mr *MockInterfaceMockRecorder) AskDeleteConfirmation() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AskDeleteConfirmation", reflect.TypeOf((*MockInterface)(nil).AskDeleteConfirmation))
}

// DeleteApplication mocks base method.
func (m *MockInterface) DeleteApplication() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DeleteApplication")
}

// DeleteApplication indicates an expected call of DeleteApplication.
func (mr *MockInterfaceMockRecorder) DeleteApplication() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteApplication", reflect.TypeOf((*MockInterface)(nil).DeleteApplication))
}

// DeleteRunner mocks base method.
func (m *MockInterface) DeleteRunner(arg0 *cobra.Command, arg1 []string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DeleteRunner", arg0, arg1)
}

// DeleteRunner indicates an expected call of DeleteRunner.
func (mr *MockInterfaceMockRecorder) DeleteRunner(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRunner", reflect.TypeOf((*MockInterface)(nil).DeleteRunner), arg0, arg1)
}

// GetApps mocks base method.
func (m *MockInterface) GetApps() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "GetApps")
}

// GetApps indicates an expected call of GetApps.
func (mr *MockInterfaceMockRecorder) GetApps() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetApps", reflect.TypeOf((*MockInterface)(nil).GetApps))
}

// GetDefaultApplicationRunner mocks base method.
func (m *MockInterface) GetDefaultApplicationRunner(arg0 *cobra.Command, arg1 []string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "GetDefaultApplicationRunner", arg0, arg1)
}

// GetDefaultApplicationRunner indicates an expected call of GetDefaultApplicationRunner.
func (mr *MockInterfaceMockRecorder) GetDefaultApplicationRunner(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDefaultApplicationRunner", reflect.TypeOf((*MockInterface)(nil).GetDefaultApplicationRunner), arg0, arg1)
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

// GetRunner mocks base method.
func (m *MockInterface) GetRunner(arg0 *cobra.Command, arg1 []string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "GetRunner", arg0, arg1)
}

// GetRunner indicates an expected call of GetRunner.
func (mr *MockInterfaceMockRecorder) GetRunner(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRunner", reflect.TypeOf((*MockInterface)(nil).GetRunner), arg0, arg1)
}

// GetSingleApplication mocks base method.
func (m *MockInterface) GetSingleApplication(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "GetSingleApplication", arg0)
}

// GetSingleApplication indicates an expected call of GetSingleApplication.
func (mr *MockInterfaceMockRecorder) GetSingleApplication(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSingleApplication", reflect.TypeOf((*MockInterface)(nil).GetSingleApplication), arg0)
}

// PrintApps mocks base method.
func (m *MockInterface) PrintApps() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "PrintApps")
}

// PrintApps indicates an expected call of PrintApps.
func (mr *MockInterfaceMockRecorder) PrintApps() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrintApps", reflect.TypeOf((*MockInterface)(nil).PrintApps))
}

// RemoveDefaultApplicationRunner mocks base method.
func (m *MockInterface) RemoveDefaultApplicationRunner(arg0 *cobra.Command, arg1 []string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RemoveDefaultApplicationRunner", arg0, arg1)
}

// RemoveDefaultApplicationRunner indicates an expected call of RemoveDefaultApplicationRunner.
func (mr *MockInterfaceMockRecorder) RemoveDefaultApplicationRunner(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveDefaultApplicationRunner", reflect.TypeOf((*MockInterface)(nil).RemoveDefaultApplicationRunner), arg0, arg1)
}

// RenameApplication mocks base method.
func (m *MockInterface) RenameApplication() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RenameApplication")
}

// RenameApplication indicates an expected call of RenameApplication.
func (mr *MockInterfaceMockRecorder) RenameApplication() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RenameApplication", reflect.TypeOf((*MockInterface)(nil).RenameApplication))
}

// RenameRunner mocks base method.
func (m *MockInterface) RenameRunner(arg0 *cobra.Command, arg1 []string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RenameRunner", arg0, arg1)
}

// RenameRunner indicates an expected call of RenameRunner.
func (mr *MockInterfaceMockRecorder) RenameRunner(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RenameRunner", reflect.TypeOf((*MockInterface)(nil).RenameRunner), arg0, arg1)
}

// SetDefaultRunner mocks base method.
func (m *MockInterface) SetDefaultRunner(arg0 *cobra.Command, arg1 []string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetDefaultRunner", arg0, arg1)
}

// SetDefaultRunner indicates an expected call of SetDefaultRunner.
func (mr *MockInterfaceMockRecorder) SetDefaultRunner(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetDefaultRunner", reflect.TypeOf((*MockInterface)(nil).SetDefaultRunner), arg0, arg1)
}
