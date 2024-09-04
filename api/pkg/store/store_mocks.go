// Code generated by MockGen. DO NOT EDIT.
// Source: store.go

// Package store is a generated GoMock package.
package store

import (
	context "context"
	reflect "reflect"


	types "github.com/helixml/helix/api/pkg/types"
	gomock "go.uber.org/mock/gomock"
)

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// CreateAPIKey mocks base method.
func (m *MockStore) CreateAPIKey(ctx context.Context, apiKey *types.APIKey) (*types.APIKey, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAPIKey", ctx, apiKey)
	ret0, _ := ret[0].(*types.APIKey)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAPIKey indicates an expected call of CreateAPIKey.
func (mr *MockStoreMockRecorder) CreateAPIKey(ctx, apiKey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAPIKey", reflect.TypeOf((*MockStore)(nil).CreateAPIKey), ctx, apiKey)
}

// CreateApp mocks base method.
func (m *MockStore) CreateApp(ctx context.Context, tool *types.App) (*types.App, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateApp", ctx, tool)
	ret0, _ := ret[0].(*types.App)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateApp indicates an expected call of CreateApp.
func (mr *MockStoreMockRecorder) CreateApp(ctx, tool interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateApp", reflect.TypeOf((*MockStore)(nil).CreateApp), ctx, tool)
}

// CreateDataEntity mocks base method.
func (m *MockStore) CreateDataEntity(ctx context.Context, dataEntity *types.DataEntity) (*types.DataEntity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateDataEntity", ctx, dataEntity)
	ret0, _ := ret[0].(*types.DataEntity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateDataEntity indicates an expected call of CreateDataEntity.
func (mr *MockStoreMockRecorder) CreateDataEntity(ctx, dataEntity interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDataEntity", reflect.TypeOf((*MockStore)(nil).CreateDataEntity), ctx, dataEntity)
}

// CreateKnowledge mocks base method.
func (m *MockStore) CreateKnowledge(ctx context.Context, knowledge *types.Knowledge) (*types.Knowledge, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateKnowledge", ctx, knowledge)
	ret0, _ := ret[0].(*types.Knowledge)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateKnowledge indicates an expected call of CreateKnowledge.
func (mr *MockStoreMockRecorder) CreateKnowledge(ctx, knowledge interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateKnowledge", reflect.TypeOf((*MockStore)(nil).CreateKnowledge), ctx, knowledge)
}

// CreateKnowledgeVersion mocks base method.
func (m *MockStore) CreateKnowledgeVersion(ctx context.Context, version *types.KnowledgeVersion) (*types.KnowledgeVersion, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateKnowledgeVersion", ctx, version)
	ret0, _ := ret[0].(*types.KnowledgeVersion)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateKnowledgeVersion indicates an expected call of CreateKnowledgeVersion.
func (mr *MockStoreMockRecorder) CreateKnowledgeVersion(ctx, version interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateKnowledgeVersion", reflect.TypeOf((*MockStore)(nil).CreateKnowledgeVersion), ctx, version)
}

// CreateLLMCall mocks base method.
func (m *MockStore) CreateLLMCall(ctx context.Context, call *types.LLMCall) (*types.LLMCall, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateLLMCall", ctx, call)
	ret0, _ := ret[0].(*types.LLMCall)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateLLMCall indicates an expected call of CreateLLMCall.
func (mr *MockStoreMockRecorder) CreateLLMCall(ctx, call interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateLLMCall", reflect.TypeOf((*MockStore)(nil).CreateLLMCall), ctx, call)
}

// CreateScriptRun mocks base method.
func (m *MockStore) CreateScriptRun(ctx context.Context, task *types.ScriptRun) (*types.ScriptRun, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateScriptRun", ctx, task)
	ret0, _ := ret[0].(*types.ScriptRun)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateScriptRun indicates an expected call of CreateScriptRun.
func (mr *MockStoreMockRecorder) CreateScriptRun(ctx, task interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateScriptRun", reflect.TypeOf((*MockStore)(nil).CreateScriptRun), ctx, task)
}

// CreateSession mocks base method.
func (m *MockStore) CreateSession(ctx context.Context, session types.Session) (*types.Session, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSession", ctx, session)
	ret0, _ := ret[0].(*types.Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSession indicates an expected call of CreateSession.
func (mr *MockStoreMockRecorder) CreateSession(ctx, session interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSession", reflect.TypeOf((*MockStore)(nil).CreateSession), ctx, session)
}

// CreateSessionToolBinding mocks base method.
func (m *MockStore) CreateSessionToolBinding(ctx context.Context, sessionID, toolID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSessionToolBinding", ctx, sessionID, toolID)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateSessionToolBinding indicates an expected call of CreateSessionToolBinding.
func (mr *MockStoreMockRecorder) CreateSessionToolBinding(ctx, sessionID, toolID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSessionToolBinding", reflect.TypeOf((*MockStore)(nil).CreateSessionToolBinding), ctx, sessionID, toolID)
}

// CreateTool mocks base method.
func (m *MockStore) CreateTool(ctx context.Context, tool *types.Tool) (*types.Tool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTool", ctx, tool)
	ret0, _ := ret[0].(*types.Tool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTool indicates an expected call of CreateTool.
func (mr *MockStoreMockRecorder) CreateTool(ctx, tool interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTool", reflect.TypeOf((*MockStore)(nil).CreateTool), ctx, tool)
}

// CreateUserMeta mocks base method.
func (m *MockStore) CreateUserMeta(ctx context.Context, UserMeta types.UserMeta) (*types.UserMeta, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUserMeta", ctx, UserMeta)
	ret0, _ := ret[0].(*types.UserMeta)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUserMeta indicates an expected call of CreateUserMeta.
func (mr *MockStoreMockRecorder) CreateUserMeta(ctx, UserMeta interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUserMeta", reflect.TypeOf((*MockStore)(nil).CreateUserMeta), ctx, UserMeta)
}

// DeleteAPIKey mocks base method.
func (m *MockStore) DeleteAPIKey(ctx context.Context, apiKey string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAPIKey", ctx, apiKey)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAPIKey indicates an expected call of DeleteAPIKey.
func (mr *MockStoreMockRecorder) DeleteAPIKey(ctx, apiKey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAPIKey", reflect.TypeOf((*MockStore)(nil).DeleteAPIKey), ctx, apiKey)
}

// DeleteApp mocks base method.
func (m *MockStore) DeleteApp(ctx context.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteApp", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteApp indicates an expected call of DeleteApp.
func (mr *MockStoreMockRecorder) DeleteApp(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteApp", reflect.TypeOf((*MockStore)(nil).DeleteApp), ctx, id)
}

// DeleteDataEntity mocks base method.
func (m *MockStore) DeleteDataEntity(ctx context.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteDataEntity", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteDataEntity indicates an expected call of DeleteDataEntity.
func (mr *MockStoreMockRecorder) DeleteDataEntity(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteDataEntity", reflect.TypeOf((*MockStore)(nil).DeleteDataEntity), ctx, id)
}

// DeleteKnowledge mocks base method.
func (m *MockStore) DeleteKnowledge(ctx context.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteKnowledge", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteKnowledge indicates an expected call of DeleteKnowledge.
func (mr *MockStoreMockRecorder) DeleteKnowledge(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteKnowledge", reflect.TypeOf((*MockStore)(nil).DeleteKnowledge), ctx, id)
}

// DeleteKnowledgeVersion mocks base method.
func (m *MockStore) DeleteKnowledgeVersion(ctx context.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteKnowledgeVersion", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteKnowledgeVersion indicates an expected call of DeleteKnowledgeVersion.
func (mr *MockStoreMockRecorder) DeleteKnowledgeVersion(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteKnowledgeVersion", reflect.TypeOf((*MockStore)(nil).DeleteKnowledgeVersion), ctx, id)
}

// DeleteScriptRun mocks base method.
func (m *MockStore) DeleteScriptRun(ctx context.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteScriptRun", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteScriptRun indicates an expected call of DeleteScriptRun.
func (mr *MockStoreMockRecorder) DeleteScriptRun(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteScriptRun", reflect.TypeOf((*MockStore)(nil).DeleteScriptRun), ctx, id)
}

// DeleteSession mocks base method.
func (m *MockStore) DeleteSession(ctx context.Context, id string) (*types.Session, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSession", ctx, id)
	ret0, _ := ret[0].(*types.Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteSession indicates an expected call of DeleteSession.
func (mr *MockStoreMockRecorder) DeleteSession(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSession", reflect.TypeOf((*MockStore)(nil).DeleteSession), ctx, id)
}

// DeleteSessionToolBinding mocks base method.
func (m *MockStore) DeleteSessionToolBinding(ctx context.Context, sessionID, toolID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSessionToolBinding", ctx, sessionID, toolID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSessionToolBinding indicates an expected call of DeleteSessionToolBinding.
func (mr *MockStoreMockRecorder) DeleteSessionToolBinding(ctx, sessionID, toolID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSessionToolBinding", reflect.TypeOf((*MockStore)(nil).DeleteSessionToolBinding), ctx, sessionID, toolID)
}

// DeleteTool mocks base method.
func (m *MockStore) DeleteTool(ctx context.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTool", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTool indicates an expected call of DeleteTool.
func (mr *MockStoreMockRecorder) DeleteTool(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTool", reflect.TypeOf((*MockStore)(nil).DeleteTool), ctx, id)
}

// EnsureUserMeta mocks base method.
func (m *MockStore) EnsureUserMeta(ctx context.Context, UserMeta types.UserMeta) (*types.UserMeta, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnsureUserMeta", ctx, UserMeta)
	ret0, _ := ret[0].(*types.UserMeta)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EnsureUserMeta indicates an expected call of EnsureUserMeta.
func (mr *MockStoreMockRecorder) EnsureUserMeta(ctx, UserMeta interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureUserMeta", reflect.TypeOf((*MockStore)(nil).EnsureUserMeta), ctx, UserMeta)
}

// GetAPIKey mocks base method.
func (m *MockStore) GetAPIKey(ctx context.Context, apiKey string) (*types.APIKey, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAPIKey", ctx, apiKey)
	ret0, _ := ret[0].(*types.APIKey)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAPIKey indicates an expected call of GetAPIKey.
func (mr *MockStoreMockRecorder) GetAPIKey(ctx, apiKey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAPIKey", reflect.TypeOf((*MockStore)(nil).GetAPIKey), ctx, apiKey)
}

// GetApp mocks base method.
func (m *MockStore) GetApp(ctx context.Context, id string) (*types.App, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetApp", ctx, id)
	ret0, _ := ret[0].(*types.App)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetApp indicates an expected call of GetApp.
func (mr *MockStoreMockRecorder) GetApp(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetApp", reflect.TypeOf((*MockStore)(nil).GetApp), ctx, id)
}

// GetDataEntity mocks base method.
func (m *MockStore) GetDataEntity(ctx context.Context, id string) (*types.DataEntity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDataEntity", ctx, id)
	ret0, _ := ret[0].(*types.DataEntity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDataEntity indicates an expected call of GetDataEntity.
func (mr *MockStoreMockRecorder) GetDataEntity(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDataEntity", reflect.TypeOf((*MockStore)(nil).GetDataEntity), ctx, id)
}

// GetKnowledge mocks base method.
func (m *MockStore) GetKnowledge(ctx context.Context, id string) (*types.Knowledge, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetKnowledge", ctx, id)
	ret0, _ := ret[0].(*types.Knowledge)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetKnowledge indicates an expected call of GetKnowledge.
func (mr *MockStoreMockRecorder) GetKnowledge(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetKnowledge", reflect.TypeOf((*MockStore)(nil).GetKnowledge), ctx, id)
}

// GetKnowledgeVersion mocks base method.
func (m *MockStore) GetKnowledgeVersion(ctx context.Context, id string) (*types.KnowledgeVersion, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetKnowledgeVersion", ctx, id)
	ret0, _ := ret[0].(*types.KnowledgeVersion)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetKnowledgeVersion indicates an expected call of GetKnowledgeVersion.
func (mr *MockStoreMockRecorder) GetKnowledgeVersion(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetKnowledgeVersion", reflect.TypeOf((*MockStore)(nil).GetKnowledgeVersion), ctx, id)
}

// GetSession mocks base method.
func (m *MockStore) GetSession(ctx context.Context, id string) (*types.Session, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSession", ctx, id)
	ret0, _ := ret[0].(*types.Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSession indicates an expected call of GetSession.
func (mr *MockStoreMockRecorder) GetSession(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSession", reflect.TypeOf((*MockStore)(nil).GetSession), ctx, id)
}

// GetSessions mocks base method.
func (m *MockStore) GetSessions(ctx context.Context, query GetSessionsQuery) ([]*types.Session, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSessions", ctx, query)
	ret0, _ := ret[0].([]*types.Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSessions indicates an expected call of GetSessions.
func (mr *MockStoreMockRecorder) GetSessions(ctx, query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSessions", reflect.TypeOf((*MockStore)(nil).GetSessions), ctx, query)
}

// GetSessionsCounter mocks base method.
func (m *MockStore) GetSessionsCounter(ctx context.Context, query GetSessionsQuery) (*types.Counter, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSessionsCounter", ctx, query)
	ret0, _ := ret[0].(*types.Counter)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSessionsCounter indicates an expected call of GetSessionsCounter.
func (mr *MockStoreMockRecorder) GetSessionsCounter(ctx, query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSessionsCounter", reflect.TypeOf((*MockStore)(nil).GetSessionsCounter), ctx, query)
}

// GetTool mocks base method.
func (m *MockStore) GetTool(ctx context.Context, id string) (*types.Tool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTool", ctx, id)
	ret0, _ := ret[0].(*types.Tool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTool indicates an expected call of GetTool.
func (mr *MockStoreMockRecorder) GetTool(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTool", reflect.TypeOf((*MockStore)(nil).GetTool), ctx, id)
}

// GetUserMeta mocks base method.
func (m *MockStore) GetUserMeta(ctx context.Context, id string) (*types.UserMeta, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserMeta", ctx, id)
	ret0, _ := ret[0].(*types.UserMeta)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserMeta indicates an expected call of GetUserMeta.
func (mr *MockStoreMockRecorder) GetUserMeta(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserMeta", reflect.TypeOf((*MockStore)(nil).GetUserMeta), ctx, id)
}

// ListAPIKeys mocks base method.
func (m *MockStore) ListAPIKeys(ctx context.Context, query *ListApiKeysQuery) ([]*types.APIKey, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAPIKeys", ctx, query)
	ret0, _ := ret[0].([]*types.APIKey)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAPIKeys indicates an expected call of ListAPIKeys.
func (mr *MockStoreMockRecorder) ListAPIKeys(ctx, query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAPIKeys", reflect.TypeOf((*MockStore)(nil).ListAPIKeys), ctx, query)
}

// ListApps mocks base method.
func (m *MockStore) ListApps(ctx context.Context, q *ListAppsQuery) ([]*types.App, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListApps", ctx, q)
	ret0, _ := ret[0].([]*types.App)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListApps indicates an expected call of ListApps.
func (mr *MockStoreMockRecorder) ListApps(ctx, q interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListApps", reflect.TypeOf((*MockStore)(nil).ListApps), ctx, q)
}

// ListDataEntities mocks base method.
func (m *MockStore) ListDataEntities(ctx context.Context, q *ListDataEntitiesQuery) ([]*types.DataEntity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListDataEntities", ctx, q)
	ret0, _ := ret[0].([]*types.DataEntity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListDataEntities indicates an expected call of ListDataEntities.
func (mr *MockStoreMockRecorder) ListDataEntities(ctx, q interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListDataEntities", reflect.TypeOf((*MockStore)(nil).ListDataEntities), ctx, q)
}

// ListKnowledge mocks base method.
func (m *MockStore) ListKnowledge(ctx context.Context, q *ListKnowledgeQuery) ([]*types.Knowledge, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListKnowledge", ctx, q)
	ret0, _ := ret[0].([]*types.Knowledge)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListKnowledge indicates an expected call of ListKnowledge.
func (mr *MockStoreMockRecorder) ListKnowledge(ctx, q interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListKnowledge", reflect.TypeOf((*MockStore)(nil).ListKnowledge), ctx, q)
}

// ListKnowledgeVersions mocks base method.
func (m *MockStore) ListKnowledgeVersions(ctx context.Context, q *ListKnowledgeVersionQuery) ([]*types.KnowledgeVersion, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListKnowledgeVersions", ctx, q)
	ret0, _ := ret[0].([]*types.KnowledgeVersion)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListKnowledgeVersions indicates an expected call of ListKnowledgeVersions.
func (mr *MockStoreMockRecorder) ListKnowledgeVersions(ctx, q interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListKnowledgeVersions", reflect.TypeOf((*MockStore)(nil).ListKnowledgeVersions), ctx, q)
}

// ListLLMCalls mocks base method.
func (m *MockStore) ListLLMCalls(ctx context.Context, page, pageSize int, sessionFilter string) ([]*types.LLMCall, int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListLLMCalls", ctx, page, pageSize, sessionFilter)
	ret0, _ := ret[0].([]*types.LLMCall)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListLLMCalls indicates an expected call of ListLLMCalls.
func (mr *MockStoreMockRecorder) ListLLMCalls(ctx, page, pageSize, sessionFilter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListLLMCalls", reflect.TypeOf((*MockStore)(nil).ListLLMCalls), ctx, page, pageSize, sessionFilter)
}

// ListScriptRuns mocks base method.
func (m *MockStore) ListScriptRuns(ctx context.Context, q *types.GptScriptRunsQuery) ([]*types.ScriptRun, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListScriptRuns", ctx, q)
	ret0, _ := ret[0].([]*types.ScriptRun)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListScriptRuns indicates an expected call of ListScriptRuns.
func (mr *MockStoreMockRecorder) ListScriptRuns(ctx, q interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListScriptRuns", reflect.TypeOf((*MockStore)(nil).ListScriptRuns), ctx, q)
}

// ListSessionTools mocks base method.
func (m *MockStore) ListSessionTools(ctx context.Context, sessionID string) ([]*types.Tool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListSessionTools", ctx, sessionID)
	ret0, _ := ret[0].([]*types.Tool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSessionTools indicates an expected call of ListSessionTools.
func (mr *MockStoreMockRecorder) ListSessionTools(ctx, sessionID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSessionTools", reflect.TypeOf((*MockStore)(nil).ListSessionTools), ctx, sessionID)
}

// ListTools mocks base method.
func (m *MockStore) ListTools(ctx context.Context, q *ListToolsQuery) ([]*types.Tool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTools", ctx, q)
	ret0, _ := ret[0].([]*types.Tool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTools indicates an expected call of ListTools.
func (mr *MockStoreMockRecorder) ListTools(ctx, q interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTools", reflect.TypeOf((*MockStore)(nil).ListTools), ctx, q)
}

// LookupKnowledge mocks base method.
func (m *MockStore) LookupKnowledge(ctx context.Context, q *LookupKnowledgeQuery) (*types.Knowledge, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LookupKnowledge", ctx, q)
	ret0, _ := ret[0].(*types.Knowledge)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LookupKnowledge indicates an expected call of LookupKnowledge.
func (mr *MockStoreMockRecorder) LookupKnowledge(ctx, q interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LookupKnowledge", reflect.TypeOf((*MockStore)(nil).LookupKnowledge), ctx, q)
}

// UpdateApp mocks base method.
func (m *MockStore) UpdateApp(ctx context.Context, tool *types.App) (*types.App, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateApp", ctx, tool)
	ret0, _ := ret[0].(*types.App)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateApp indicates an expected call of UpdateApp.
func (mr *MockStoreMockRecorder) UpdateApp(ctx, tool interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateApp", reflect.TypeOf((*MockStore)(nil).UpdateApp), ctx, tool)
}

// UpdateDataEntity mocks base method.
func (m *MockStore) UpdateDataEntity(ctx context.Context, dataEntity *types.DataEntity) (*types.DataEntity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateDataEntity", ctx, dataEntity)
	ret0, _ := ret[0].(*types.DataEntity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateDataEntity indicates an expected call of UpdateDataEntity.
func (mr *MockStoreMockRecorder) UpdateDataEntity(ctx, dataEntity interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateDataEntity", reflect.TypeOf((*MockStore)(nil).UpdateDataEntity), ctx, dataEntity)
}

// UpdateKnowledge mocks base method.
func (m *MockStore) UpdateKnowledge(ctx context.Context, knowledge *types.Knowledge) (*types.Knowledge, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateKnowledge", ctx, knowledge)
	ret0, _ := ret[0].(*types.Knowledge)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateKnowledge indicates an expected call of UpdateKnowledge.
func (mr *MockStoreMockRecorder) UpdateKnowledge(ctx, knowledge interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateKnowledge", reflect.TypeOf((*MockStore)(nil).UpdateKnowledge), ctx, knowledge)
}

// UpdateSession mocks base method.
func (m *MockStore) UpdateSession(ctx context.Context, session types.Session) (*types.Session, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateSession", ctx, session)
	ret0, _ := ret[0].(*types.Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateSession indicates an expected call of UpdateSession.
func (mr *MockStoreMockRecorder) UpdateSession(ctx, session interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSession", reflect.TypeOf((*MockStore)(nil).UpdateSession), ctx, session)
}

// UpdateSessionMeta mocks base method.
func (m *MockStore) UpdateSessionMeta(ctx context.Context, data types.SessionMetaUpdate) (*types.Session, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateSessionMeta", ctx, data)
	ret0, _ := ret[0].(*types.Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateSessionMeta indicates an expected call of UpdateSessionMeta.
func (mr *MockStoreMockRecorder) UpdateSessionMeta(ctx, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSessionMeta", reflect.TypeOf((*MockStore)(nil).UpdateSessionMeta), ctx, data)
}

// UpdateTool mocks base method.
func (m *MockStore) UpdateTool(ctx context.Context, tool *types.Tool) (*types.Tool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTool", ctx, tool)
	ret0, _ := ret[0].(*types.Tool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateTool indicates an expected call of UpdateTool.
func (mr *MockStoreMockRecorder) UpdateTool(ctx, tool interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTool", reflect.TypeOf((*MockStore)(nil).UpdateTool), ctx, tool)
}

// UpdateUserMeta mocks base method.
func (m *MockStore) UpdateUserMeta(ctx context.Context, UserMeta types.UserMeta) (*types.UserMeta, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserMeta", ctx, UserMeta)
	ret0, _ := ret[0].(*types.UserMeta)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUserMeta indicates an expected call of UpdateUserMeta.
func (mr *MockStoreMockRecorder) UpdateUserMeta(ctx, UserMeta interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserMeta", reflect.TypeOf((*MockStore)(nil).UpdateUserMeta), ctx, UserMeta)
}
