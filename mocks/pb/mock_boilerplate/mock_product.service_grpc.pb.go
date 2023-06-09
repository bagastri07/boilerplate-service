// Code generated by MockGen. DO NOT EDIT.
// Source: ./pb/boilerplate/product.service_grpc.pb.go

// Package mock_boilerplate is a generated GoMock package.
package mock_boilerplate

import (
	context "context"
	reflect "reflect"

	boilerplate "github.com/bagastri07/boilerplate-service/pb/boilerplate"
	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
)

// MockProductServiceClient is a mock of ProductServiceClient interface.
type MockProductServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockProductServiceClientMockRecorder
}

// MockProductServiceClientMockRecorder is the mock recorder for MockProductServiceClient.
type MockProductServiceClientMockRecorder struct {
	mock *MockProductServiceClient
}

// NewMockProductServiceClient creates a new mock instance.
func NewMockProductServiceClient(ctrl *gomock.Controller) *MockProductServiceClient {
	mock := &MockProductServiceClient{ctrl: ctrl}
	mock.recorder = &MockProductServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProductServiceClient) EXPECT() *MockProductServiceClientMockRecorder {
	return m.recorder
}

// CreateProduct mocks base method.
func (m *MockProductServiceClient) CreateProduct(ctx context.Context, in *boilerplate.CreateProductRequest, opts ...grpc.CallOption) (*boilerplate.CreateProductResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateProduct", varargs...)
	ret0, _ := ret[0].(*boilerplate.CreateProductResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateProduct indicates an expected call of CreateProduct.
func (mr *MockProductServiceClientMockRecorder) CreateProduct(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProduct", reflect.TypeOf((*MockProductServiceClient)(nil).CreateProduct), varargs...)
}

// FindByID mocks base method.
func (m *MockProductServiceClient) FindByID(ctx context.Context, in *boilerplate.FindByIDRequest, opts ...grpc.CallOption) (*boilerplate.Product, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindByID", varargs...)
	ret0, _ := ret[0].(*boilerplate.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByID indicates an expected call of FindByID.
func (mr *MockProductServiceClientMockRecorder) FindByID(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByID", reflect.TypeOf((*MockProductServiceClient)(nil).FindByID), varargs...)
}

// MockProductServiceServer is a mock of ProductServiceServer interface.
type MockProductServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockProductServiceServerMockRecorder
}

// MockProductServiceServerMockRecorder is the mock recorder for MockProductServiceServer.
type MockProductServiceServerMockRecorder struct {
	mock *MockProductServiceServer
}

// NewMockProductServiceServer creates a new mock instance.
func NewMockProductServiceServer(ctrl *gomock.Controller) *MockProductServiceServer {
	mock := &MockProductServiceServer{ctrl: ctrl}
	mock.recorder = &MockProductServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProductServiceServer) EXPECT() *MockProductServiceServerMockRecorder {
	return m.recorder
}

// CreateProduct mocks base method.
func (m *MockProductServiceServer) CreateProduct(arg0 context.Context, arg1 *boilerplate.CreateProductRequest) (*boilerplate.CreateProductResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateProduct", arg0, arg1)
	ret0, _ := ret[0].(*boilerplate.CreateProductResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateProduct indicates an expected call of CreateProduct.
func (mr *MockProductServiceServerMockRecorder) CreateProduct(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProduct", reflect.TypeOf((*MockProductServiceServer)(nil).CreateProduct), arg0, arg1)
}

// FindByID mocks base method.
func (m *MockProductServiceServer) FindByID(arg0 context.Context, arg1 *boilerplate.FindByIDRequest) (*boilerplate.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByID", arg0, arg1)
	ret0, _ := ret[0].(*boilerplate.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByID indicates an expected call of FindByID.
func (mr *MockProductServiceServerMockRecorder) FindByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByID", reflect.TypeOf((*MockProductServiceServer)(nil).FindByID), arg0, arg1)
}

// mustEmbedUnimplementedProductServiceServer mocks base method.
func (m *MockProductServiceServer) mustEmbedUnimplementedProductServiceServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedProductServiceServer")
}

// mustEmbedUnimplementedProductServiceServer indicates an expected call of mustEmbedUnimplementedProductServiceServer.
func (mr *MockProductServiceServerMockRecorder) mustEmbedUnimplementedProductServiceServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedProductServiceServer", reflect.TypeOf((*MockProductServiceServer)(nil).mustEmbedUnimplementedProductServiceServer))
}

// MockUnsafeProductServiceServer is a mock of UnsafeProductServiceServer interface.
type MockUnsafeProductServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockUnsafeProductServiceServerMockRecorder
}

// MockUnsafeProductServiceServerMockRecorder is the mock recorder for MockUnsafeProductServiceServer.
type MockUnsafeProductServiceServerMockRecorder struct {
	mock *MockUnsafeProductServiceServer
}

// NewMockUnsafeProductServiceServer creates a new mock instance.
func NewMockUnsafeProductServiceServer(ctrl *gomock.Controller) *MockUnsafeProductServiceServer {
	mock := &MockUnsafeProductServiceServer{ctrl: ctrl}
	mock.recorder = &MockUnsafeProductServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnsafeProductServiceServer) EXPECT() *MockUnsafeProductServiceServerMockRecorder {
	return m.recorder
}

// mustEmbedUnimplementedProductServiceServer mocks base method.
func (m *MockUnsafeProductServiceServer) mustEmbedUnimplementedProductServiceServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedProductServiceServer")
}

// mustEmbedUnimplementedProductServiceServer indicates an expected call of mustEmbedUnimplementedProductServiceServer.
func (mr *MockUnsafeProductServiceServerMockRecorder) mustEmbedUnimplementedProductServiceServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedProductServiceServer", reflect.TypeOf((*MockUnsafeProductServiceServer)(nil).mustEmbedUnimplementedProductServiceServer))
}
