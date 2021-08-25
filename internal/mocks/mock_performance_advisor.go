// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mongodb/mongocli/internal/store (interfaces: PerformanceAdvisorNamespacesLister,PerformanceAdvisorSlowQueriesLister,PerformanceAdvisorIndexesLister)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	mongodbatlas "go.mongodb.org/atlas/mongodbatlas"
)

// MockPerformanceAdvisorNamespacesLister is a mock of PerformanceAdvisorNamespacesLister interface.
type MockPerformanceAdvisorNamespacesLister struct {
	ctrl     *gomock.Controller
	recorder *MockPerformanceAdvisorNamespacesListerMockRecorder
}

// MockPerformanceAdvisorNamespacesListerMockRecorder is the mock recorder for MockPerformanceAdvisorNamespacesLister.
type MockPerformanceAdvisorNamespacesListerMockRecorder struct {
	mock *MockPerformanceAdvisorNamespacesLister
}

// NewMockPerformanceAdvisorNamespacesLister creates a new mock instance.
func NewMockPerformanceAdvisorNamespacesLister(ctrl *gomock.Controller) *MockPerformanceAdvisorNamespacesLister {
	mock := &MockPerformanceAdvisorNamespacesLister{ctrl: ctrl}
	mock.recorder = &MockPerformanceAdvisorNamespacesListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPerformanceAdvisorNamespacesLister) EXPECT() *MockPerformanceAdvisorNamespacesListerMockRecorder {
	return m.recorder
}

// PerformanceAdvisorNamespaces mocks base method.
func (m *MockPerformanceAdvisorNamespacesLister) PerformanceAdvisorNamespaces(arg0, arg1 string, arg2 *mongodbatlas.NamespaceOptions) (*mongodbatlas.Namespaces, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PerformanceAdvisorNamespaces", arg0, arg1, arg2)
	ret0, _ := ret[0].(*mongodbatlas.Namespaces)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PerformanceAdvisorNamespaces indicates an expected call of PerformanceAdvisorNamespaces.
func (mr *MockPerformanceAdvisorNamespacesListerMockRecorder) PerformanceAdvisorNamespaces(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PerformanceAdvisorNamespaces", reflect.TypeOf((*MockPerformanceAdvisorNamespacesLister)(nil).PerformanceAdvisorNamespaces), arg0, arg1, arg2)
}

// MockPerformanceAdvisorSlowQueriesLister is a mock of PerformanceAdvisorSlowQueriesLister interface.
type MockPerformanceAdvisorSlowQueriesLister struct {
	ctrl     *gomock.Controller
	recorder *MockPerformanceAdvisorSlowQueriesListerMockRecorder
}

// MockPerformanceAdvisorSlowQueriesListerMockRecorder is the mock recorder for MockPerformanceAdvisorSlowQueriesLister.
type MockPerformanceAdvisorSlowQueriesListerMockRecorder struct {
	mock *MockPerformanceAdvisorSlowQueriesLister
}

// NewMockPerformanceAdvisorSlowQueriesLister creates a new mock instance.
func NewMockPerformanceAdvisorSlowQueriesLister(ctrl *gomock.Controller) *MockPerformanceAdvisorSlowQueriesLister {
	mock := &MockPerformanceAdvisorSlowQueriesLister{ctrl: ctrl}
	mock.recorder = &MockPerformanceAdvisorSlowQueriesListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPerformanceAdvisorSlowQueriesLister) EXPECT() *MockPerformanceAdvisorSlowQueriesListerMockRecorder {
	return m.recorder
}

// PerformanceAdvisorSlowQueries mocks base method.
func (m *MockPerformanceAdvisorSlowQueriesLister) PerformanceAdvisorSlowQueries(arg0, arg1 string, arg2 *mongodbatlas.SlowQueryOptions) (*mongodbatlas.SlowQueries, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PerformanceAdvisorSlowQueries", arg0, arg1, arg2)
	ret0, _ := ret[0].(*mongodbatlas.SlowQueries)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PerformanceAdvisorSlowQueries indicates an expected call of PerformanceAdvisorSlowQueries.
func (mr *MockPerformanceAdvisorSlowQueriesListerMockRecorder) PerformanceAdvisorSlowQueries(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PerformanceAdvisorSlowQueries", reflect.TypeOf((*MockPerformanceAdvisorSlowQueriesLister)(nil).PerformanceAdvisorSlowQueries), arg0, arg1, arg2)
}

// MockPerformanceAdvisorIndexesLister is a mock of PerformanceAdvisorIndexesLister interface.
type MockPerformanceAdvisorIndexesLister struct {
	ctrl     *gomock.Controller
	recorder *MockPerformanceAdvisorIndexesListerMockRecorder
}

// MockPerformanceAdvisorIndexesListerMockRecorder is the mock recorder for MockPerformanceAdvisorIndexesLister.
type MockPerformanceAdvisorIndexesListerMockRecorder struct {
	mock *MockPerformanceAdvisorIndexesLister
}

// NewMockPerformanceAdvisorIndexesLister creates a new mock instance.
func NewMockPerformanceAdvisorIndexesLister(ctrl *gomock.Controller) *MockPerformanceAdvisorIndexesLister {
	mock := &MockPerformanceAdvisorIndexesLister{ctrl: ctrl}
	mock.recorder = &MockPerformanceAdvisorIndexesListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPerformanceAdvisorIndexesLister) EXPECT() *MockPerformanceAdvisorIndexesListerMockRecorder {
	return m.recorder
}

// PerformanceAdvisorIndexes mocks base method.
func (m *MockPerformanceAdvisorIndexesLister) PerformanceAdvisorIndexes(arg0, arg1 string, arg2 *mongodbatlas.SuggestedIndexOptions) (*mongodbatlas.SuggestedIndexes, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PerformanceAdvisorIndexes", arg0, arg1, arg2)
	ret0, _ := ret[0].(*mongodbatlas.SuggestedIndexes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PerformanceAdvisorIndexes indicates an expected call of PerformanceAdvisorIndexes.
func (mr *MockPerformanceAdvisorIndexesListerMockRecorder) PerformanceAdvisorIndexes(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PerformanceAdvisorIndexes", reflect.TypeOf((*MockPerformanceAdvisorIndexesLister)(nil).PerformanceAdvisorIndexes), arg0, arg1, arg2)
}
