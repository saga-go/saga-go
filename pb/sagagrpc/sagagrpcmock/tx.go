// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/saga-go/saga-go/pb/sagagrpc (interfaces: TxEventServiceClient)

// Package sagagrpcmock is a generated GoMock package.
package sagagrpcmock

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	sagagrpc "github.com/saga-go/saga-go/pb/sagagrpc"
	grpc "google.golang.org/grpc"
	reflect "reflect"
)

// MockTxEventServiceClient is a mock of TxEventServiceClient interface
type MockTxEventServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockTxEventServiceClientMockRecorder
}

// MockTxEventServiceClientMockRecorder is the mock recorder for MockTxEventServiceClient
type MockTxEventServiceClientMockRecorder struct {
	mock *MockTxEventServiceClient
}

// NewMockTxEventServiceClient creates a new mock instance
func NewMockTxEventServiceClient(ctrl *gomock.Controller) *MockTxEventServiceClient {
	mock := &MockTxEventServiceClient{ctrl: ctrl}
	mock.recorder = &MockTxEventServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTxEventServiceClient) EXPECT() *MockTxEventServiceClientMockRecorder {
	return m.recorder
}

// OnConnected mocks base method
func (m *MockTxEventServiceClient) OnConnected(arg0 context.Context, arg1 ...grpc.CallOption) (sagagrpc.TxEventService_OnConnectedClient, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "OnConnected", varargs...)
	ret0, _ := ret[0].(sagagrpc.TxEventService_OnConnectedClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OnConnected indicates an expected call of OnConnected
func (mr *MockTxEventServiceClientMockRecorder) OnConnected(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnConnected", reflect.TypeOf((*MockTxEventServiceClient)(nil).OnConnected), varargs...)
}

// OnDisconnected mocks base method
func (m *MockTxEventServiceClient) OnDisconnected(arg0 context.Context, arg1 *sagagrpc.GrpcServiceConfig, arg2 ...grpc.CallOption) (*sagagrpc.GrpcAck, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "OnDisconnected", varargs...)
	ret0, _ := ret[0].(*sagagrpc.GrpcAck)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OnDisconnected indicates an expected call of OnDisconnected
func (mr *MockTxEventServiceClientMockRecorder) OnDisconnected(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnDisconnected", reflect.TypeOf((*MockTxEventServiceClient)(nil).OnDisconnected), varargs...)
}

// OnGetServerMeta mocks base method
func (m *MockTxEventServiceClient) OnGetServerMeta(arg0 context.Context, arg1 *sagagrpc.GrpcServiceConfig, arg2 ...grpc.CallOption) (*sagagrpc.ServerMeta, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "OnGetServerMeta", varargs...)
	ret0, _ := ret[0].(*sagagrpc.ServerMeta)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OnGetServerMeta indicates an expected call of OnGetServerMeta
func (mr *MockTxEventServiceClientMockRecorder) OnGetServerMeta(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnGetServerMeta", reflect.TypeOf((*MockTxEventServiceClient)(nil).OnGetServerMeta), varargs...)
}

// OnTxEvent mocks base method
func (m *MockTxEventServiceClient) OnTxEvent(arg0 context.Context, arg1 *sagagrpc.GrpcTxEvent, arg2 ...grpc.CallOption) (*sagagrpc.GrpcAck, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "OnTxEvent", varargs...)
	ret0, _ := ret[0].(*sagagrpc.GrpcAck)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OnTxEvent indicates an expected call of OnTxEvent
func (mr *MockTxEventServiceClientMockRecorder) OnTxEvent(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnTxEvent", reflect.TypeOf((*MockTxEventServiceClient)(nil).OnTxEvent), varargs...)
}
