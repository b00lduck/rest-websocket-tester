// Code generated by MockGen. DO NOT EDIT.
// Source: interfaces.go

// Package mock_ws is a generated GoMock package.
package mock_ws

import (
	dto "github.com/b00lduck/rest-websocket-tester/internal/dto"
	ws "github.com/b00lduck/rest-websocket-tester/internal/ws"
	gomock "github.com/golang/mock/gomock"
	io "io"
	net "net"
	reflect "reflect"
	time "time"
)

// MockClient is a mock of Client interface
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// Send mocks base method
func (m *MockClient) Send() chan dto.Message {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send")
	ret0, _ := ret[0].(chan dto.Message)
	return ret0
}

// Send indicates an expected call of Send
func (mr *MockClientMockRecorder) Send() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockClient)(nil).Send))
}

// Close mocks base method
func (m *MockClient) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close
func (mr *MockClientMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockClient)(nil).Close))
}

// Lock mocks base method
func (m *MockClient) Lock() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Lock")
}

// Lock indicates an expected call of Lock
func (mr *MockClientMockRecorder) Lock() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Lock", reflect.TypeOf((*MockClient)(nil).Lock))
}

// Unlock mocks base method
func (m *MockClient) Unlock() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Unlock")
}

// Unlock indicates an expected call of Unlock
func (mr *MockClientMockRecorder) Unlock() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unlock", reflect.TypeOf((*MockClient)(nil).Unlock))
}

// MockHub is a mock of Hub interface
type MockHub struct {
	ctrl     *gomock.Controller
	recorder *MockHubMockRecorder
}

// MockHubMockRecorder is the mock recorder for MockHub
type MockHubMockRecorder struct {
	mock *MockHub
}

// NewMockHub creates a new mock instance
func NewMockHub(ctrl *gomock.Controller) *MockHub {
	mock := &MockHub{ctrl: ctrl}
	mock.recorder = &MockHubMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockHub) EXPECT() *MockHubMockRecorder {
	return m.recorder
}

// Run mocks base method
func (m *MockHub) Run() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Run")
}

// Run indicates an expected call of Run
func (mr *MockHubMockRecorder) Run() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockHub)(nil).Run))
}

// Unregister mocks base method
func (m *MockHub) Unregister(c ws.Client) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Unregister", c)
}

// Unregister indicates an expected call of Unregister
func (mr *MockHubMockRecorder) Unregister(c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unregister", reflect.TypeOf((*MockHub)(nil).Unregister), c)
}

// Register mocks base method
func (m *MockHub) Register(c ws.Client) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Register", c)
}

// Register indicates an expected call of Register
func (mr *MockHubMockRecorder) Register(c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockHub)(nil).Register), c)
}

// Broadcast mocks base method
func (m *MockHub) Broadcast(message dto.Message) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Broadcast", message)
}

// Broadcast indicates an expected call of Broadcast
func (mr *MockHubMockRecorder) Broadcast(message interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Broadcast", reflect.TypeOf((*MockHub)(nil).Broadcast), message)
}

// NumClients mocks base method
func (m *MockHub) NumClients() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NumClients")
	ret0, _ := ret[0].(int)
	return ret0
}

// NumClients indicates an expected call of NumClients
func (mr *MockHubMockRecorder) NumClients() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NumClients", reflect.TypeOf((*MockHub)(nil).NumClients))
}

// MockConnI is a mock of ConnI interface
type MockConnI struct {
	ctrl     *gomock.Controller
	recorder *MockConnIMockRecorder
}

// MockConnIMockRecorder is the mock recorder for MockConnI
type MockConnIMockRecorder struct {
	mock *MockConnI
}

// NewMockConnI creates a new mock instance
func NewMockConnI(ctrl *gomock.Controller) *MockConnI {
	mock := &MockConnI{ctrl: ctrl}
	mock.recorder = &MockConnIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockConnI) EXPECT() *MockConnIMockRecorder {
	return m.recorder
}

// RemoteAddr mocks base method
func (m *MockConnI) RemoteAddr() net.Addr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoteAddr")
	ret0, _ := ret[0].(net.Addr)
	return ret0
}

// RemoteAddr indicates an expected call of RemoteAddr
func (mr *MockConnIMockRecorder) RemoteAddr() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoteAddr", reflect.TypeOf((*MockConnI)(nil).RemoteAddr))
}

// Close mocks base method
func (m *MockConnI) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockConnIMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockConnI)(nil).Close))
}

// SetReadLimit mocks base method
func (m *MockConnI) SetReadLimit(arg0 int64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetReadLimit", arg0)
}

// SetReadLimit indicates an expected call of SetReadLimit
func (mr *MockConnIMockRecorder) SetReadLimit(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetReadLimit", reflect.TypeOf((*MockConnI)(nil).SetReadLimit), arg0)
}

// SetReadDeadline mocks base method
func (m *MockConnI) SetReadDeadline(arg0 time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetReadDeadline", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetReadDeadline indicates an expected call of SetReadDeadline
func (mr *MockConnIMockRecorder) SetReadDeadline(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetReadDeadline", reflect.TypeOf((*MockConnI)(nil).SetReadDeadline), arg0)
}

// SetWriteDeadline mocks base method
func (m *MockConnI) SetWriteDeadline(arg0 time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetWriteDeadline", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetWriteDeadline indicates an expected call of SetWriteDeadline
func (mr *MockConnIMockRecorder) SetWriteDeadline(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetWriteDeadline", reflect.TypeOf((*MockConnI)(nil).SetWriteDeadline), arg0)
}

// SetPongHandler mocks base method
func (m *MockConnI) SetPongHandler(arg0 func(string) error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetPongHandler", arg0)
}

// SetPongHandler indicates an expected call of SetPongHandler
func (mr *MockConnIMockRecorder) SetPongHandler(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetPongHandler", reflect.TypeOf((*MockConnI)(nil).SetPongHandler), arg0)
}

// ReadMessage mocks base method
func (m *MockConnI) ReadMessage() (int, []byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadMessage")
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].([]byte)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ReadMessage indicates an expected call of ReadMessage
func (mr *MockConnIMockRecorder) ReadMessage() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadMessage", reflect.TypeOf((*MockConnI)(nil).ReadMessage))
}

// WriteMessage mocks base method
func (m *MockConnI) WriteMessage(arg0 int, arg1 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteMessage", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteMessage indicates an expected call of WriteMessage
func (mr *MockConnIMockRecorder) WriteMessage(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteMessage", reflect.TypeOf((*MockConnI)(nil).WriteMessage), arg0, arg1)
}

// NextWriter mocks base method
func (m *MockConnI) NextWriter(arg0 int) (io.WriteCloser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NextWriter", arg0)
	ret0, _ := ret[0].(io.WriteCloser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NextWriter indicates an expected call of NextWriter
func (mr *MockConnIMockRecorder) NextWriter(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NextWriter", reflect.TypeOf((*MockConnI)(nil).NextWriter), arg0)
}