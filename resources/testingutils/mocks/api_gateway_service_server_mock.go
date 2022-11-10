// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/yandex-cloud/go-genproto/yandex/cloud/serverless/apigateway/v1 (interfaces: ApiGatewayServiceServer)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	access "github.com/yandex-cloud/go-genproto/yandex/cloud/access"
	operation "github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
	apigateway "github.com/yandex-cloud/go-genproto/yandex/cloud/serverless/apigateway/v1"
	reflect "reflect"
)

// MockApiGatewayServiceServer is a mock of ApiGatewayServiceServer interface
type MockApiGatewayServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockApiGatewayServiceServerMockRecorder
}

// MockApiGatewayServiceServerMockRecorder is the mock recorder for MockApiGatewayServiceServer
type MockApiGatewayServiceServerMockRecorder struct {
	mock *MockApiGatewayServiceServer
}

// NewMockApiGatewayServiceServer creates a new mock instance
func NewMockApiGatewayServiceServer(ctrl *gomock.Controller) *MockApiGatewayServiceServer {
	mock := &MockApiGatewayServiceServer{ctrl: ctrl}
	mock.recorder = &MockApiGatewayServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockApiGatewayServiceServer) EXPECT() *MockApiGatewayServiceServerMockRecorder {
	return m.recorder
}

// AddDomain mocks base method
func (m *MockApiGatewayServiceServer) AddDomain(arg0 context.Context, arg1 *apigateway.AddDomainRequest) (*operation.Operation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddDomain", arg0, arg1)
	ret0, _ := ret[0].(*operation.Operation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddDomain indicates an expected call of AddDomain
func (mr *MockApiGatewayServiceServerMockRecorder) AddDomain(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddDomain", reflect.TypeOf((*MockApiGatewayServiceServer)(nil).AddDomain), arg0, arg1)
}

// Create mocks base method
func (m *MockApiGatewayServiceServer) Create(arg0 context.Context, arg1 *apigateway.CreateApiGatewayRequest) (*operation.Operation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(*operation.Operation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockApiGatewayServiceServerMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockApiGatewayServiceServer)(nil).Create), arg0, arg1)
}

// Delete mocks base method
func (m *MockApiGatewayServiceServer) Delete(arg0 context.Context, arg1 *apigateway.DeleteApiGatewayRequest) (*operation.Operation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(*operation.Operation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete
func (mr *MockApiGatewayServiceServerMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockApiGatewayServiceServer)(nil).Delete), arg0, arg1)
}

// Get mocks base method
func (m *MockApiGatewayServiceServer) Get(arg0 context.Context, arg1 *apigateway.GetApiGatewayRequest) (*apigateway.ApiGateway, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(*apigateway.ApiGateway)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockApiGatewayServiceServerMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockApiGatewayServiceServer)(nil).Get), arg0, arg1)
}

// GetOpenapiSpec mocks base method
func (m *MockApiGatewayServiceServer) GetOpenapiSpec(arg0 context.Context, arg1 *apigateway.GetOpenapiSpecRequest) (*apigateway.GetOpenapiSpecResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOpenapiSpec", arg0, arg1)
	ret0, _ := ret[0].(*apigateway.GetOpenapiSpecResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOpenapiSpec indicates an expected call of GetOpenapiSpec
func (mr *MockApiGatewayServiceServerMockRecorder) GetOpenapiSpec(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOpenapiSpec", reflect.TypeOf((*MockApiGatewayServiceServer)(nil).GetOpenapiSpec), arg0, arg1)
}

// List mocks base method
func (m *MockApiGatewayServiceServer) List(arg0 context.Context, arg1 *apigateway.ListApiGatewayRequest) (*apigateway.ListApiGatewayResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].(*apigateway.ListApiGatewayResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockApiGatewayServiceServerMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockApiGatewayServiceServer)(nil).List), arg0, arg1)
}

// ListAccessBindings mocks base method
func (m *MockApiGatewayServiceServer) ListAccessBindings(arg0 context.Context, arg1 *access.ListAccessBindingsRequest) (*access.ListAccessBindingsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAccessBindings", arg0, arg1)
	ret0, _ := ret[0].(*access.ListAccessBindingsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAccessBindings indicates an expected call of ListAccessBindings
func (mr *MockApiGatewayServiceServerMockRecorder) ListAccessBindings(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAccessBindings", reflect.TypeOf((*MockApiGatewayServiceServer)(nil).ListAccessBindings), arg0, arg1)
}

// ListOperations mocks base method
func (m *MockApiGatewayServiceServer) ListOperations(arg0 context.Context, arg1 *apigateway.ListOperationsRequest) (*apigateway.ListOperationsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListOperations", arg0, arg1)
	ret0, _ := ret[0].(*apigateway.ListOperationsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListOperations indicates an expected call of ListOperations
func (mr *MockApiGatewayServiceServerMockRecorder) ListOperations(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListOperations", reflect.TypeOf((*MockApiGatewayServiceServer)(nil).ListOperations), arg0, arg1)
}

// RemoveDomain mocks base method
func (m *MockApiGatewayServiceServer) RemoveDomain(arg0 context.Context, arg1 *apigateway.RemoveDomainRequest) (*operation.Operation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveDomain", arg0, arg1)
	ret0, _ := ret[0].(*operation.Operation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoveDomain indicates an expected call of RemoveDomain
func (mr *MockApiGatewayServiceServerMockRecorder) RemoveDomain(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveDomain", reflect.TypeOf((*MockApiGatewayServiceServer)(nil).RemoveDomain), arg0, arg1)
}

// SetAccessBindings mocks base method
func (m *MockApiGatewayServiceServer) SetAccessBindings(arg0 context.Context, arg1 *access.SetAccessBindingsRequest) (*operation.Operation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetAccessBindings", arg0, arg1)
	ret0, _ := ret[0].(*operation.Operation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetAccessBindings indicates an expected call of SetAccessBindings
func (mr *MockApiGatewayServiceServerMockRecorder) SetAccessBindings(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetAccessBindings", reflect.TypeOf((*MockApiGatewayServiceServer)(nil).SetAccessBindings), arg0, arg1)
}

// Update mocks base method
func (m *MockApiGatewayServiceServer) Update(arg0 context.Context, arg1 *apigateway.UpdateApiGatewayRequest) (*operation.Operation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(*operation.Operation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update
func (mr *MockApiGatewayServiceServerMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockApiGatewayServiceServer)(nil).Update), arg0, arg1)
}

// UpdateAccessBindings mocks base method
func (m *MockApiGatewayServiceServer) UpdateAccessBindings(arg0 context.Context, arg1 *access.UpdateAccessBindingsRequest) (*operation.Operation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAccessBindings", arg0, arg1)
	ret0, _ := ret[0].(*operation.Operation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAccessBindings indicates an expected call of UpdateAccessBindings
func (mr *MockApiGatewayServiceServerMockRecorder) UpdateAccessBindings(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAccessBindings", reflect.TypeOf((*MockApiGatewayServiceServer)(nil).UpdateAccessBindings), arg0, arg1)
}
