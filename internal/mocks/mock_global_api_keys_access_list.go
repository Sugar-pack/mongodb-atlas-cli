// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mongodb/mongocli/internal/store (interfaces: GlobalAPIKeyWhitelistLister,GlobalAPIKeyWhitelistDescriber,GlobalAPIKeyWhitelistCreator,GlobalAPIKeyWhitelistDeleter)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	mongodbatlas "go.mongodb.org/atlas/mongodbatlas"
	opsmngr "go.mongodb.org/ops-manager/opsmngr"
)

// MockGlobalAPIKeyWhitelistLister is a mock of GlobalAPIKeyWhitelistLister interface.
type MockGlobalAPIKeyWhitelistLister struct {
	ctrl     *gomock.Controller
	recorder *MockGlobalAPIKeyWhitelistListerMockRecorder
}

// MockGlobalAPIKeyWhitelistListerMockRecorder is the mock recorder for MockGlobalAPIKeyWhitelistLister.
type MockGlobalAPIKeyWhitelistListerMockRecorder struct {
	mock *MockGlobalAPIKeyWhitelistLister
}

// NewMockGlobalAPIKeyWhitelistLister creates a new mock instance.
func NewMockGlobalAPIKeyWhitelistLister(ctrl *gomock.Controller) *MockGlobalAPIKeyWhitelistLister {
	mock := &MockGlobalAPIKeyWhitelistLister{ctrl: ctrl}
	mock.recorder = &MockGlobalAPIKeyWhitelistListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGlobalAPIKeyWhitelistLister) EXPECT() *MockGlobalAPIKeyWhitelistListerMockRecorder {
	return m.recorder
}

// GlobalAPIKeyWhitelists mocks base method.
func (m *MockGlobalAPIKeyWhitelistLister) GlobalAPIKeyWhitelists(arg0 *mongodbatlas.ListOptions) (*opsmngr.GlobalWhitelistAPIKeys, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GlobalAPIKeyWhitelists", arg0)
	ret0, _ := ret[0].(*opsmngr.GlobalWhitelistAPIKeys)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GlobalAPIKeyWhitelists indicates an expected call of GlobalAPIKeyWhitelists.
func (mr *MockGlobalAPIKeyWhitelistListerMockRecorder) GlobalAPIKeyWhitelists(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GlobalAPIKeyWhitelists", reflect.TypeOf((*MockGlobalAPIKeyWhitelistLister)(nil).GlobalAPIKeyWhitelists), arg0)
}

// MockGlobalAPIKeyWhitelistDescriber is a mock of GlobalAPIKeyWhitelistDescriber interface.
type MockGlobalAPIKeyWhitelistDescriber struct {
	ctrl     *gomock.Controller
	recorder *MockGlobalAPIKeyWhitelistDescriberMockRecorder
}

// MockGlobalAPIKeyWhitelistDescriberMockRecorder is the mock recorder for MockGlobalAPIKeyWhitelistDescriber.
type MockGlobalAPIKeyWhitelistDescriberMockRecorder struct {
	mock *MockGlobalAPIKeyWhitelistDescriber
}

// NewMockGlobalAPIKeyWhitelistDescriber creates a new mock instance.
func NewMockGlobalAPIKeyWhitelistDescriber(ctrl *gomock.Controller) *MockGlobalAPIKeyWhitelistDescriber {
	mock := &MockGlobalAPIKeyWhitelistDescriber{ctrl: ctrl}
	mock.recorder = &MockGlobalAPIKeyWhitelistDescriberMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGlobalAPIKeyWhitelistDescriber) EXPECT() *MockGlobalAPIKeyWhitelistDescriberMockRecorder {
	return m.recorder
}

// GlobalAPIKeyWhitelist mocks base method.
func (m *MockGlobalAPIKeyWhitelistDescriber) GlobalAPIKeyWhitelist(arg0 string) (*opsmngr.GlobalWhitelistAPIKey, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GlobalAPIKeyWhitelist", arg0)
	ret0, _ := ret[0].(*opsmngr.GlobalWhitelistAPIKey)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GlobalAPIKeyWhitelist indicates an expected call of GlobalAPIKeyWhitelist.
func (mr *MockGlobalAPIKeyWhitelistDescriberMockRecorder) GlobalAPIKeyWhitelist(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GlobalAPIKeyWhitelist", reflect.TypeOf((*MockGlobalAPIKeyWhitelistDescriber)(nil).GlobalAPIKeyWhitelist), arg0)
}

// MockGlobalAPIKeyWhitelistCreator is a mock of GlobalAPIKeyWhitelistCreator interface.
type MockGlobalAPIKeyWhitelistCreator struct {
	ctrl     *gomock.Controller
	recorder *MockGlobalAPIKeyWhitelistCreatorMockRecorder
}

// MockGlobalAPIKeyWhitelistCreatorMockRecorder is the mock recorder for MockGlobalAPIKeyWhitelistCreator.
type MockGlobalAPIKeyWhitelistCreatorMockRecorder struct {
	mock *MockGlobalAPIKeyWhitelistCreator
}

// NewMockGlobalAPIKeyWhitelistCreator creates a new mock instance.
func NewMockGlobalAPIKeyWhitelistCreator(ctrl *gomock.Controller) *MockGlobalAPIKeyWhitelistCreator {
	mock := &MockGlobalAPIKeyWhitelistCreator{ctrl: ctrl}
	mock.recorder = &MockGlobalAPIKeyWhitelistCreatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGlobalAPIKeyWhitelistCreator) EXPECT() *MockGlobalAPIKeyWhitelistCreatorMockRecorder {
	return m.recorder
}

// CreateGlobalAPIKeyWhitelist mocks base method.
func (m *MockGlobalAPIKeyWhitelistCreator) CreateGlobalAPIKeyWhitelist(arg0 *opsmngr.WhitelistAPIKeysReq) (*opsmngr.GlobalWhitelistAPIKey, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateGlobalAPIKeyWhitelist", arg0)
	ret0, _ := ret[0].(*opsmngr.GlobalWhitelistAPIKey)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateGlobalAPIKeyWhitelist indicates an expected call of CreateGlobalAPIKeyWhitelist.
func (mr *MockGlobalAPIKeyWhitelistCreatorMockRecorder) CreateGlobalAPIKeyWhitelist(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateGlobalAPIKeyWhitelist", reflect.TypeOf((*MockGlobalAPIKeyWhitelistCreator)(nil).CreateGlobalAPIKeyWhitelist), arg0)
}

// MockGlobalAPIKeyWhitelistDeleter is a mock of GlobalAPIKeyWhitelistDeleter interface.
type MockGlobalAPIKeyWhitelistDeleter struct {
	ctrl     *gomock.Controller
	recorder *MockGlobalAPIKeyWhitelistDeleterMockRecorder
}

// MockGlobalAPIKeyWhitelistDeleterMockRecorder is the mock recorder for MockGlobalAPIKeyWhitelistDeleter.
type MockGlobalAPIKeyWhitelistDeleterMockRecorder struct {
	mock *MockGlobalAPIKeyWhitelistDeleter
}

// NewMockGlobalAPIKeyWhitelistDeleter creates a new mock instance.
func NewMockGlobalAPIKeyWhitelistDeleter(ctrl *gomock.Controller) *MockGlobalAPIKeyWhitelistDeleter {
	mock := &MockGlobalAPIKeyWhitelistDeleter{ctrl: ctrl}
	mock.recorder = &MockGlobalAPIKeyWhitelistDeleterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGlobalAPIKeyWhitelistDeleter) EXPECT() *MockGlobalAPIKeyWhitelistDeleterMockRecorder {
	return m.recorder
}

// DeleteGlobalAPIKeyWhitelist mocks base method.
func (m *MockGlobalAPIKeyWhitelistDeleter) DeleteGlobalAPIKeyWhitelist(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteGlobalAPIKeyWhitelist", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteGlobalAPIKeyWhitelist indicates an expected call of DeleteGlobalAPIKeyWhitelist.
func (mr *MockGlobalAPIKeyWhitelistDeleterMockRecorder) DeleteGlobalAPIKeyWhitelist(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteGlobalAPIKeyWhitelist", reflect.TypeOf((*MockGlobalAPIKeyWhitelistDeleter)(nil).DeleteGlobalAPIKeyWhitelist), arg0)
}
