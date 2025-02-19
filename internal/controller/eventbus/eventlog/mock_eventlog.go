// Code generated by MockGen. DO NOT EDIT.
// Source: eventlog.go

// Package eventlog is a generated GoMock package.
package eventlog

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	metadata "github.com/linkall-labs/vanus/internal/controller/eventbus/metadata"
	kv "github.com/linkall-labs/vanus/internal/kv"
	vanus "github.com/linkall-labs/vanus/internal/primitive/vanus"
)

// MockManager is a mock of Manager interface.
type MockManager struct {
	ctrl     *gomock.Controller
	recorder *MockManagerMockRecorder
}

// MockManagerMockRecorder is the mock recorder for MockManager.
type MockManagerMockRecorder struct {
	mock *MockManager
}

// NewMockManager creates a new mock instance.
func NewMockManager(ctrl *gomock.Controller) *MockManager {
	mock := &MockManager{ctrl: ctrl}
	mock.recorder = &MockManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockManager) EXPECT() *MockManagerMockRecorder {
	return m.recorder
}

// AcquireEventLog mocks base method.
func (m *MockManager) AcquireEventLog(ctx context.Context, eventbusID vanus.ID, eventbusName string) (*metadata.Eventlog, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AcquireEventLog", ctx, eventbusID, eventbusName)
	ret0, _ := ret[0].(*metadata.Eventlog)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AcquireEventLog indicates an expected call of AcquireEventLog.
func (mr *MockManagerMockRecorder) AcquireEventLog(ctx, eventbusID, eventbusName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AcquireEventLog", reflect.TypeOf((*MockManager)(nil).AcquireEventLog), ctx, eventbusID, eventbusName)
}

// DeleteEventlog mocks base method.
func (m *MockManager) DeleteEventlog(ctx context.Context, id vanus.ID) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DeleteEventlog", ctx, id)
}

// DeleteEventlog indicates an expected call of DeleteEventlog.
func (mr *MockManagerMockRecorder) DeleteEventlog(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteEventlog", reflect.TypeOf((*MockManager)(nil).DeleteEventlog), ctx, id)
}

// GetAppendableSegment mocks base method.
func (m *MockManager) GetAppendableSegment(ctx context.Context, eli *metadata.Eventlog, num int) ([]Segment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAppendableSegment", ctx, eli, num)
	ret0, _ := ret[0].([]Segment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAppendableSegment indicates an expected call of GetAppendableSegment.
func (mr *MockManagerMockRecorder) GetAppendableSegment(ctx, eli, num interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAppendableSegment", reflect.TypeOf((*MockManager)(nil).GetAppendableSegment), ctx, eli, num)
}

// GetBlock mocks base method.
func (m *MockManager) GetBlock(id vanus.ID) *metadata.Block {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBlock", id)
	ret0, _ := ret[0].(*metadata.Block)
	return ret0
}

// GetBlock indicates an expected call of GetBlock.
func (mr *MockManagerMockRecorder) GetBlock(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlock", reflect.TypeOf((*MockManager)(nil).GetBlock), id)
}

// GetEventLog mocks base method.
func (m *MockManager) GetEventLog(ctx context.Context, id vanus.ID) *metadata.Eventlog {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEventLog", ctx, id)
	ret0, _ := ret[0].(*metadata.Eventlog)
	return ret0
}

// GetEventLog indicates an expected call of GetEventLog.
func (mr *MockManagerMockRecorder) GetEventLog(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEventLog", reflect.TypeOf((*MockManager)(nil).GetEventLog), ctx, id)
}

// GetEventLogSegmentList mocks base method.
func (m *MockManager) GetEventLogSegmentList(elID vanus.ID) []Segment {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEventLogSegmentList", elID)
	ret0, _ := ret[0].([]Segment)
	return ret0
}

// GetEventLogSegmentList indicates an expected call of GetEventLogSegmentList.
func (mr *MockManagerMockRecorder) GetEventLogSegmentList(elID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEventLogSegmentList", reflect.TypeOf((*MockManager)(nil).GetEventLogSegmentList), elID)
}

// GetSegmentByBlockID mocks base method.
func (m *MockManager) GetSegmentByBlockID(block *metadata.Block) (Segment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSegmentByBlockID", block)
	ret0, _ := ret[0].(Segment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSegmentByBlockID indicates an expected call of GetSegmentByBlockID.
func (mr *MockManagerMockRecorder) GetSegmentByBlockID(block interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSegmentByBlockID", reflect.TypeOf((*MockManager)(nil).GetSegmentByBlockID), block)
}

// Run mocks base method.
func (m *MockManager) Run(ctx context.Context, kvClient kv.Client, startTask bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Run", ctx, kvClient, startTask)
	ret0, _ := ret[0].(error)
	return ret0
}

// Run indicates an expected call of Run.
func (mr *MockManagerMockRecorder) Run(ctx, kvClient, startTask interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockManager)(nil).Run), ctx, kvClient, startTask)
}

// Stop mocks base method.
func (m *MockManager) Stop() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Stop")
}

// Stop indicates an expected call of Stop.
func (mr *MockManagerMockRecorder) Stop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockManager)(nil).Stop))
}

// UpdateSegment mocks base method.
func (m_2 *MockManager) UpdateSegment(ctx context.Context, m map[string][]Segment) {
	m_2.ctrl.T.Helper()
	m_2.ctrl.Call(m_2, "UpdateSegment", ctx, m)
}

// UpdateSegment indicates an expected call of UpdateSegment.
func (mr *MockManagerMockRecorder) UpdateSegment(ctx, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSegment", reflect.TypeOf((*MockManager)(nil).UpdateSegment), ctx, m)
}

// UpdateSegmentReplicas mocks base method.
func (m *MockManager) UpdateSegmentReplicas(ctx context.Context, segID vanus.ID, term uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateSegmentReplicas", ctx, segID, term)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateSegmentReplicas indicates an expected call of UpdateSegmentReplicas.
func (mr *MockManagerMockRecorder) UpdateSegmentReplicas(ctx, segID, term interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSegmentReplicas", reflect.TypeOf((*MockManager)(nil).UpdateSegmentReplicas), ctx, segID, term)
}
