// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mongodb/mongocli/internal/store (interfaces: OnlineArchiveLister,OnlineArchiveDescriber,OnlineArchiveDeleter)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	mongodbatlas "go.mongodb.org/atlas/mongodbatlas"
	reflect "reflect"
)

// MockOnlineArchiveLister is a mock of OnlineArchiveLister interface
type MockOnlineArchiveLister struct {
	ctrl     *gomock.Controller
	recorder *MockOnlineArchiveListerMockRecorder
}

// MockOnlineArchiveListerMockRecorder is the mock recorder for MockOnlineArchiveLister
type MockOnlineArchiveListerMockRecorder struct {
	mock *MockOnlineArchiveLister
}

// NewMockOnlineArchiveLister creates a new mock instance
func NewMockOnlineArchiveLister(ctrl *gomock.Controller) *MockOnlineArchiveLister {
	mock := &MockOnlineArchiveLister{ctrl: ctrl}
	mock.recorder = &MockOnlineArchiveListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOnlineArchiveLister) EXPECT() *MockOnlineArchiveListerMockRecorder {
	return m.recorder
}

// OnlineArchives mocks base method
func (m *MockOnlineArchiveLister) OnlineArchives(arg0, arg1 string) ([]*mongodbatlas.OnlineArchive, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OnlineArchives", arg0, arg1)
	ret0, _ := ret[0].([]*mongodbatlas.OnlineArchive)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OnlineArchives indicates an expected call of OnlineArchives
func (mr *MockOnlineArchiveListerMockRecorder) OnlineArchives(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnlineArchives", reflect.TypeOf((*MockOnlineArchiveLister)(nil).OnlineArchives), arg0, arg1)
}

// MockOnlineArchiveDescriber is a mock of OnlineArchiveDescriber interface
type MockOnlineArchiveDescriber struct {
	ctrl     *gomock.Controller
	recorder *MockOnlineArchiveDescriberMockRecorder
}

// MockOnlineArchiveDescriberMockRecorder is the mock recorder for MockOnlineArchiveDescriber
type MockOnlineArchiveDescriberMockRecorder struct {
	mock *MockOnlineArchiveDescriber
}

// NewMockOnlineArchiveDescriber creates a new mock instance
func NewMockOnlineArchiveDescriber(ctrl *gomock.Controller) *MockOnlineArchiveDescriber {
	mock := &MockOnlineArchiveDescriber{ctrl: ctrl}
	mock.recorder = &MockOnlineArchiveDescriberMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOnlineArchiveDescriber) EXPECT() *MockOnlineArchiveDescriberMockRecorder {
	return m.recorder
}

// OnlineArchive mocks base method
func (m *MockOnlineArchiveDescriber) OnlineArchive(arg0, arg1, arg2 string) (*mongodbatlas.OnlineArchive, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OnlineArchive", arg0, arg1, arg2)
	ret0, _ := ret[0].(*mongodbatlas.OnlineArchive)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OnlineArchive indicates an expected call of OnlineArchive
func (mr *MockOnlineArchiveDescriberMockRecorder) OnlineArchive(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnlineArchive", reflect.TypeOf((*MockOnlineArchiveDescriber)(nil).OnlineArchive), arg0, arg1, arg2)
}

// MockOnlineArchiveDeleter is a mock of OnlineArchiveDeleter interface
type MockOnlineArchiveDeleter struct {
	ctrl     *gomock.Controller
	recorder *MockOnlineArchiveDeleterMockRecorder
}

// MockOnlineArchiveDeleterMockRecorder is the mock recorder for MockOnlineArchiveDeleter
type MockOnlineArchiveDeleterMockRecorder struct {
	mock *MockOnlineArchiveDeleter
}

// NewMockOnlineArchiveDeleter creates a new mock instance
func NewMockOnlineArchiveDeleter(ctrl *gomock.Controller) *MockOnlineArchiveDeleter {
	mock := &MockOnlineArchiveDeleter{ctrl: ctrl}
	mock.recorder = &MockOnlineArchiveDeleterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOnlineArchiveDeleter) EXPECT() *MockOnlineArchiveDeleterMockRecorder {
	return m.recorder
}

// DeleteOnlineArchive mocks base method
func (m *MockOnlineArchiveDeleter) DeleteOnlineArchive(arg0, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteOnlineArchive", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteOnlineArchive indicates an expected call of DeleteOnlineArchive
func (mr *MockOnlineArchiveDeleterMockRecorder) DeleteOnlineArchive(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteOnlineArchive", reflect.TypeOf((*MockOnlineArchiveDeleter)(nil).DeleteOnlineArchive), arg0, arg1, arg2)
}