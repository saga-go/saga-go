// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/saga-go/saga-go/pb/sagagrpc (interfaces: TxEventService_OnConnectedClient)

// Package sagagrpcmock is a generated GoMock package.
package sagagrpcmock

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	sagagrpc "github.com/saga-go/saga-go/pb/sagagrpc"
	metadata "google.golang.org/grpc/metadata"
	reflect "reflect"
)

// MockTxEventService_OnConnectedClient is a mock of TxEventService_OnConnectedClient interface
type MockTxEventService_OnConnectedClient struct {
	ctrl     *gomock.Controller
	recorder *MockTxEventService_OnConnectedClientMockRecorder
}

// MockTxEventService_OnConnectedClientMockRecorder is the mock recorder for MockTxEventService_OnConnectedClient
type MockTxEventService_OnConnectedClientMockRecorder struct {
	mock *MockTxEventService_OnConnectedClient
}

// NewMockTxEventService_OnConnectedClient creates a new mock instance
func NewMockTxEventService_OnConnectedClient(ctrl *gomock.Controller) *MockTxEventService_OnConnectedClient {
	mock := &MockTxEventService_OnConnectedClient{ctrl: ctrl}
	mock.recorder = &MockTxEventService_OnConnectedClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTxEventService_OnConnectedClient) EXPECT() *MockTxEventService_OnConnectedClientMockRecorder {
	return m.recorder
}

// CloseSend mocks base method
func (m *MockTxEventService_OnConnectedClient) CloseSend() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseSend")
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseSend indicates an expected call of CloseSend
func (mr *MockTxEventService_OnConnectedClientMockRecorder) CloseSend() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseSend", reflect.TypeOf((*MockTxEventService_OnConnectedClient)(nil).CloseSend))
}

// Context mocks base method
func (m *MockTxEventService_OnConnectedClient) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context
func (mr *MockTxEventService_OnConnectedClientMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockTxEventService_OnConnectedClient)(nil).Context))
}

// Header mocks base method
func (m *MockTxEventService_OnConnectedClient) Header() (metadata.MD, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Header")
	ret0, _ := ret[0].(metadata.MD)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Header indicates an expected call of Header
func (mr *MockTxEventService_OnConnectedClientMockRecorder) Header() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Header", reflect.TypeOf((*MockTxEventService_OnConnectedClient)(nil).Header))
}

// Recv mocks base method
func (m *MockTxEventService_OnConnectedClient) Recv() (*sagagrpc.GrpcCompensateCommand, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Recv")
	ret0, _ := ret[0].(*sagagrpc.GrpcCompensateCommand)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Recv indicates an expected call of Recv
func (mr *MockTxEventService_OnConnectedClientMockRecorder) Recv() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Recv", reflect.TypeOf((*MockTxEventService_OnConnectedClient)(nil).Recv))
}

// RecvMsg mocks base method
func (m *MockTxEventService_OnConnectedClient) RecvMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecvMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg
func (mr *MockTxEventService_OnConnectedClientMockRecorder) RecvMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockTxEventService_OnConnectedClient)(nil).RecvMsg), arg0)
}

// Send mocks base method
func (m *MockTxEventService_OnConnectedClient) Send(arg0 *sagagrpc.GrpcServiceConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send
func (mr *MockTxEventService_OnConnectedClientMockRecorder) Send(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockTxEventService_OnConnectedClient)(nil).Send), arg0)
}

// SendMsg mocks base method
func (m *MockTxEventService_OnConnectedClient) SendMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg
func (mr *MockTxEventService_OnConnectedClientMockRecorder) SendMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockTxEventService_OnConnectedClient)(nil).SendMsg), arg0)
}

// Trailer mocks base method
func (m *MockTxEventService_OnConnectedClient) Trailer() metadata.MD {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Trailer")
	ret0, _ := ret[0].(metadata.MD)
	return ret0
}

// Trailer indicates an expected call of Trailer
func (mr *MockTxEventService_OnConnectedClientMockRecorder) Trailer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Trailer", reflect.TypeOf((*MockTxEventService_OnConnectedClient)(nil).Trailer))
}
