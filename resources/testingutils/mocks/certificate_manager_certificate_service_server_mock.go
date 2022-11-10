// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/yandex-cloud/go-genproto/yandex/cloud/certificatemanager/v1 (interfaces: CertificateServiceServer)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	access "github.com/yandex-cloud/go-genproto/yandex/cloud/access"
	certificatemanager "github.com/yandex-cloud/go-genproto/yandex/cloud/certificatemanager/v1"
	operation "github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
	reflect "reflect"
)

// MockCertificateServiceServer is a mock of CertificateServiceServer interface
type MockCertificateServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockCertificateServiceServerMockRecorder
}

// MockCertificateServiceServerMockRecorder is the mock recorder for MockCertificateServiceServer
type MockCertificateServiceServerMockRecorder struct {
	mock *MockCertificateServiceServer
}

// NewMockCertificateServiceServer creates a new mock instance
func NewMockCertificateServiceServer(ctrl *gomock.Controller) *MockCertificateServiceServer {
	mock := &MockCertificateServiceServer{ctrl: ctrl}
	mock.recorder = &MockCertificateServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCertificateServiceServer) EXPECT() *MockCertificateServiceServerMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockCertificateServiceServer) Create(arg0 context.Context, arg1 *certificatemanager.CreateCertificateRequest) (*operation.Operation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(*operation.Operation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockCertificateServiceServerMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockCertificateServiceServer)(nil).Create), arg0, arg1)
}

// Delete mocks base method
func (m *MockCertificateServiceServer) Delete(arg0 context.Context, arg1 *certificatemanager.DeleteCertificateRequest) (*operation.Operation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(*operation.Operation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete
func (mr *MockCertificateServiceServerMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockCertificateServiceServer)(nil).Delete), arg0, arg1)
}

// Get mocks base method
func (m *MockCertificateServiceServer) Get(arg0 context.Context, arg1 *certificatemanager.GetCertificateRequest) (*certificatemanager.Certificate, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(*certificatemanager.Certificate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockCertificateServiceServerMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockCertificateServiceServer)(nil).Get), arg0, arg1)
}

// List mocks base method
func (m *MockCertificateServiceServer) List(arg0 context.Context, arg1 *certificatemanager.ListCertificatesRequest) (*certificatemanager.ListCertificatesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].(*certificatemanager.ListCertificatesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockCertificateServiceServerMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockCertificateServiceServer)(nil).List), arg0, arg1)
}

// ListAccessBindings mocks base method
func (m *MockCertificateServiceServer) ListAccessBindings(arg0 context.Context, arg1 *access.ListAccessBindingsRequest) (*access.ListAccessBindingsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAccessBindings", arg0, arg1)
	ret0, _ := ret[0].(*access.ListAccessBindingsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAccessBindings indicates an expected call of ListAccessBindings
func (mr *MockCertificateServiceServerMockRecorder) ListAccessBindings(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAccessBindings", reflect.TypeOf((*MockCertificateServiceServer)(nil).ListAccessBindings), arg0, arg1)
}

// ListOperations mocks base method
func (m *MockCertificateServiceServer) ListOperations(arg0 context.Context, arg1 *certificatemanager.ListCertificateOperationsRequest) (*certificatemanager.ListCertificateOperationsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListOperations", arg0, arg1)
	ret0, _ := ret[0].(*certificatemanager.ListCertificateOperationsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListOperations indicates an expected call of ListOperations
func (mr *MockCertificateServiceServerMockRecorder) ListOperations(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListOperations", reflect.TypeOf((*MockCertificateServiceServer)(nil).ListOperations), arg0, arg1)
}

// ListVersions mocks base method
func (m *MockCertificateServiceServer) ListVersions(arg0 context.Context, arg1 *certificatemanager.ListVersionsRequest) (*certificatemanager.ListVersionsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListVersions", arg0, arg1)
	ret0, _ := ret[0].(*certificatemanager.ListVersionsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListVersions indicates an expected call of ListVersions
func (mr *MockCertificateServiceServerMockRecorder) ListVersions(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListVersions", reflect.TypeOf((*MockCertificateServiceServer)(nil).ListVersions), arg0, arg1)
}

// RequestNew mocks base method
func (m *MockCertificateServiceServer) RequestNew(arg0 context.Context, arg1 *certificatemanager.RequestNewCertificateRequest) (*operation.Operation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RequestNew", arg0, arg1)
	ret0, _ := ret[0].(*operation.Operation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RequestNew indicates an expected call of RequestNew
func (mr *MockCertificateServiceServerMockRecorder) RequestNew(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RequestNew", reflect.TypeOf((*MockCertificateServiceServer)(nil).RequestNew), arg0, arg1)
}

// SetAccessBindings mocks base method
func (m *MockCertificateServiceServer) SetAccessBindings(arg0 context.Context, arg1 *access.SetAccessBindingsRequest) (*operation.Operation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetAccessBindings", arg0, arg1)
	ret0, _ := ret[0].(*operation.Operation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetAccessBindings indicates an expected call of SetAccessBindings
func (mr *MockCertificateServiceServerMockRecorder) SetAccessBindings(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetAccessBindings", reflect.TypeOf((*MockCertificateServiceServer)(nil).SetAccessBindings), arg0, arg1)
}

// Update mocks base method
func (m *MockCertificateServiceServer) Update(arg0 context.Context, arg1 *certificatemanager.UpdateCertificateRequest) (*operation.Operation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(*operation.Operation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update
func (mr *MockCertificateServiceServerMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockCertificateServiceServer)(nil).Update), arg0, arg1)
}

// UpdateAccessBindings mocks base method
func (m *MockCertificateServiceServer) UpdateAccessBindings(arg0 context.Context, arg1 *access.UpdateAccessBindingsRequest) (*operation.Operation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAccessBindings", arg0, arg1)
	ret0, _ := ret[0].(*operation.Operation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAccessBindings indicates an expected call of UpdateAccessBindings
func (mr *MockCertificateServiceServerMockRecorder) UpdateAccessBindings(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAccessBindings", reflect.TypeOf((*MockCertificateServiceServer)(nil).UpdateAccessBindings), arg0, arg1)
}
