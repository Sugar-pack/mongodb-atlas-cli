// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mongodb/mongocli/internal/store (interfaces: MaintenanceWindowUpdater,MaintenanceWindowClearer,MaintenanceWindowDeferrer,MaintenanceWindowDescriber)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	mongodbatlas "go.mongodb.org/atlas/mongodbatlas"
	reflect "reflect"
)

// MockMaintenanceWindowUpdater is a mock of MaintenanceWindowUpdater interface
type MockMaintenanceWindowUpdater struct {
	ctrl     *gomock.Controller
	recorder *MockMaintenanceWindowUpdaterMockRecorder
}

// MockMaintenanceWindowUpdaterMockRecorder is the mock recorder for MockMaintenanceWindowUpdater
type MockMaintenanceWindowUpdaterMockRecorder struct {
	mock *MockMaintenanceWindowUpdater
}

// NewMockMaintenanceWindowUpdater creates a new mock instance
func NewMockMaintenanceWindowUpdater(ctrl *gomock.Controller) *MockMaintenanceWindowUpdater {
	mock := &MockMaintenanceWindowUpdater{ctrl: ctrl}
	mock.recorder = &MockMaintenanceWindowUpdaterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMaintenanceWindowUpdater) EXPECT() *MockMaintenanceWindowUpdaterMockRecorder {
	return m.recorder
}

// UpdateMaintenanceWindow mocks base method
func (m *MockMaintenanceWindowUpdater) UpdateMaintenanceWindow(arg0 string, arg1 *mongodbatlas.MaintenanceWindow) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateMaintenanceWindow", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateMaintenanceWindow indicates an expected call of UpdateMaintenanceWindow
func (mr *MockMaintenanceWindowUpdaterMockRecorder) UpdateMaintenanceWindow(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMaintenanceWindow", reflect.TypeOf((*MockMaintenanceWindowUpdater)(nil).UpdateMaintenanceWindow), arg0, arg1)
}

// MockMaintenanceWindowClearer is a mock of MaintenanceWindowClearer interface
type MockMaintenanceWindowClearer struct {
	ctrl     *gomock.Controller
	recorder *MockMaintenanceWindowClearerMockRecorder
}

// MockMaintenanceWindowClearerMockRecorder is the mock recorder for MockMaintenanceWindowClearer
type MockMaintenanceWindowClearerMockRecorder struct {
	mock *MockMaintenanceWindowClearer
}

// NewMockMaintenanceWindowClearer creates a new mock instance
func NewMockMaintenanceWindowClearer(ctrl *gomock.Controller) *MockMaintenanceWindowClearer {
	mock := &MockMaintenanceWindowClearer{ctrl: ctrl}
	mock.recorder = &MockMaintenanceWindowClearerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMaintenanceWindowClearer) EXPECT() *MockMaintenanceWindowClearerMockRecorder {
	return m.recorder
}

// ClearMaintenanceWindow mocks base method
func (m *MockMaintenanceWindowClearer) ClearMaintenanceWindow(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClearMaintenanceWindow", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// ClearMaintenanceWindow indicates an expected call of ClearMaintenanceWindow
func (mr *MockMaintenanceWindowClearerMockRecorder) ClearMaintenanceWindow(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClearMaintenanceWindow", reflect.TypeOf((*MockMaintenanceWindowClearer)(nil).ClearMaintenanceWindow), arg0)
}

// MockMaintenanceWindowDeferrer is a mock of MaintenanceWindowDeferrer interface
type MockMaintenanceWindowDeferrer struct {
	ctrl     *gomock.Controller
	recorder *MockMaintenanceWindowDeferrerMockRecorder
}

// MockMaintenanceWindowDeferrerMockRecorder is the mock recorder for MockMaintenanceWindowDeferrer
type MockMaintenanceWindowDeferrerMockRecorder struct {
	mock *MockMaintenanceWindowDeferrer
}

// NewMockMaintenanceWindowDeferrer creates a new mock instance
func NewMockMaintenanceWindowDeferrer(ctrl *gomock.Controller) *MockMaintenanceWindowDeferrer {
	mock := &MockMaintenanceWindowDeferrer{ctrl: ctrl}
	mock.recorder = &MockMaintenanceWindowDeferrerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMaintenanceWindowDeferrer) EXPECT() *MockMaintenanceWindowDeferrerMockRecorder {
	return m.recorder
}

// DeferMaintenanceWindow mocks base method
func (m *MockMaintenanceWindowDeferrer) DeferMaintenanceWindow(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeferMaintenanceWindow", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeferMaintenanceWindow indicates an expected call of DeferMaintenanceWindow
func (mr *MockMaintenanceWindowDeferrerMockRecorder) DeferMaintenanceWindow(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeferMaintenanceWindow", reflect.TypeOf((*MockMaintenanceWindowDeferrer)(nil).DeferMaintenanceWindow), arg0)
}

// MockMaintenanceWindowDescriber is a mock of MaintenanceWindowDescriber interface
type MockMaintenanceWindowDescriber struct {
	ctrl     *gomock.Controller
	recorder *MockMaintenanceWindowDescriberMockRecorder
}

// MockMaintenanceWindowDescriberMockRecorder is the mock recorder for MockMaintenanceWindowDescriber
type MockMaintenanceWindowDescriberMockRecorder struct {
	mock *MockMaintenanceWindowDescriber
}

// NewMockMaintenanceWindowDescriber creates a new mock instance
func NewMockMaintenanceWindowDescriber(ctrl *gomock.Controller) *MockMaintenanceWindowDescriber {
	mock := &MockMaintenanceWindowDescriber{ctrl: ctrl}
	mock.recorder = &MockMaintenanceWindowDescriberMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMaintenanceWindowDescriber) EXPECT() *MockMaintenanceWindowDescriberMockRecorder {
	return m.recorder
}

// MaintenanceWindow mocks base method
func (m *MockMaintenanceWindowDescriber) MaintenanceWindow(arg0 string) (*mongodbatlas.MaintenanceWindow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MaintenanceWindow", arg0)
	ret0, _ := ret[0].(*mongodbatlas.MaintenanceWindow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MaintenanceWindow indicates an expected call of MaintenanceWindow
func (mr *MockMaintenanceWindowDescriberMockRecorder) MaintenanceWindow(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MaintenanceWindow", reflect.TypeOf((*MockMaintenanceWindowDescriber)(nil).MaintenanceWindow), arg0)
}