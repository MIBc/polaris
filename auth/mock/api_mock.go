// Code generated by MockGen. DO NOT EDIT.
// Source: api.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	auth "github.com/polarismesh/polaris-server/auth"
	cache "github.com/polarismesh/polaris-server/cache"
	v1 "github.com/polarismesh/polaris-server/common/api/v1"
	model "github.com/polarismesh/polaris-server/common/model"
	store "github.com/polarismesh/polaris-server/store"
	reflect "reflect"
)

// MockAuthServer is a mock of AuthServer interface
type MockAuthServer struct {
	ctrl     *gomock.Controller
	recorder *MockAuthServerMockRecorder
}

// MockAuthServerMockRecorder is the mock recorder for MockAuthServer
type MockAuthServerMockRecorder struct {
	mock *MockAuthServer
}

// NewMockAuthServer creates a new mock instance
func NewMockAuthServer(ctrl *gomock.Controller) *MockAuthServer {
	mock := &MockAuthServer{ctrl: ctrl}
	mock.recorder = &MockAuthServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAuthServer) EXPECT() *MockAuthServerMockRecorder {
	return m.recorder
}

// Initialize mocks base method
func (m *MockAuthServer) Initialize(authOpt *auth.Config, storage store.Store, cacheMgn *cache.NamingCache) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Initialize", authOpt, storage, cacheMgn)
	ret0, _ := ret[0].(error)
	return ret0
}

// Initialize indicates an expected call of Initialize
func (mr *MockAuthServerMockRecorder) Initialize(authOpt, storage, cacheMgn interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Initialize", reflect.TypeOf((*MockAuthServer)(nil).Initialize), authOpt, storage, cacheMgn)
}

// Name mocks base method
func (m *MockAuthServer) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name
func (mr *MockAuthServerMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockAuthServer)(nil).Name))
}

// GetAuthChecker mocks base method
func (m *MockAuthServer) GetAuthChecker() auth.AuthChecker {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAuthChecker")
	ret0, _ := ret[0].(auth.AuthChecker)
	return ret0
}

// GetAuthChecker indicates an expected call of GetAuthChecker
func (mr *MockAuthServerMockRecorder) GetAuthChecker() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAuthChecker", reflect.TypeOf((*MockAuthServer)(nil).GetAuthChecker))
}

// AfterResourceOperation mocks base method
func (m *MockAuthServer) AfterResourceOperation(afterCtx *model.AcquireContext) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AfterResourceOperation", afterCtx)
}

// AfterResourceOperation indicates an expected call of AfterResourceOperation
func (mr *MockAuthServerMockRecorder) AfterResourceOperation(afterCtx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AfterResourceOperation", reflect.TypeOf((*MockAuthServer)(nil).AfterResourceOperation), afterCtx)
}

// Login mocks base method
func (m *MockAuthServer) Login(req *v1.LoginRequest) *v1.Response {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", req)
	ret0, _ := ret[0].(*v1.Response)
	return ret0
}

// Login indicates an expected call of Login
func (mr *MockAuthServerMockRecorder) Login(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockAuthServer)(nil).Login), req)
}

// CreateUsers mocks base method
func (m *MockAuthServer) CreateUsers(ctx context.Context, users []*v1.User) *v1.BatchWriteResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUsers", ctx, users)
	ret0, _ := ret[0].(*v1.BatchWriteResponse)
	return ret0
}

// CreateUsers indicates an expected call of CreateUsers
func (mr *MockAuthServerMockRecorder) CreateUsers(ctx, users interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUsers", reflect.TypeOf((*MockAuthServer)(nil).CreateUsers), ctx, users)
}

// UpdateUser mocks base method
func (m *MockAuthServer) UpdateUser(ctx context.Context, user *v1.User) *v1.Response {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", ctx, user)
	ret0, _ := ret[0].(*v1.Response)
	return ret0
}

// UpdateUser indicates an expected call of UpdateUser
func (mr *MockAuthServerMockRecorder) UpdateUser(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockAuthServer)(nil).UpdateUser), ctx, user)
}

// DeleteUsers mocks base method
func (m *MockAuthServer) DeleteUsers(ctx context.Context, users []*v1.User) *v1.BatchWriteResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUsers", ctx, users)
	ret0, _ := ret[0].(*v1.BatchWriteResponse)
	return ret0
}

// DeleteUsers indicates an expected call of DeleteUsers
func (mr *MockAuthServerMockRecorder) DeleteUsers(ctx, users interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUsers", reflect.TypeOf((*MockAuthServer)(nil).DeleteUsers), ctx, users)
}

// GetUsers mocks base method
func (m *MockAuthServer) GetUsers(ctx context.Context, query map[string]string) *v1.BatchQueryResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUsers", ctx, query)
	ret0, _ := ret[0].(*v1.BatchQueryResponse)
	return ret0
}

// GetUsers indicates an expected call of GetUsers
func (mr *MockAuthServerMockRecorder) GetUsers(ctx, query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsers", reflect.TypeOf((*MockAuthServer)(nil).GetUsers), ctx, query)
}

// GetUserToken mocks base method
func (m *MockAuthServer) GetUserToken(ctx context.Context, user *v1.User) *v1.Response {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserToken", ctx, user)
	ret0, _ := ret[0].(*v1.Response)
	return ret0
}

// GetUserToken indicates an expected call of GetUserToken
func (mr *MockAuthServerMockRecorder) GetUserToken(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserToken", reflect.TypeOf((*MockAuthServer)(nil).GetUserToken), ctx, user)
}

// UpdateUserToken mocks base method
func (m *MockAuthServer) UpdateUserToken(ctx context.Context, user *v1.User) *v1.Response {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserToken", ctx, user)
	ret0, _ := ret[0].(*v1.Response)
	return ret0
}

// UpdateUserToken indicates an expected call of UpdateUserToken
func (mr *MockAuthServerMockRecorder) UpdateUserToken(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserToken", reflect.TypeOf((*MockAuthServer)(nil).UpdateUserToken), ctx, user)
}

// ResetUserToken mocks base method
func (m *MockAuthServer) ResetUserToken(ctx context.Context, user *v1.User) *v1.Response {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResetUserToken", ctx, user)
	ret0, _ := ret[0].(*v1.Response)
	return ret0
}

// ResetUserToken indicates an expected call of ResetUserToken
func (mr *MockAuthServerMockRecorder) ResetUserToken(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResetUserToken", reflect.TypeOf((*MockAuthServer)(nil).ResetUserToken), ctx, user)
}

// CreateGroup mocks base method
func (m *MockAuthServer) CreateGroup(ctx context.Context, group *v1.UserGroup) *v1.Response {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateGroup", ctx, group)
	ret0, _ := ret[0].(*v1.Response)
	return ret0
}

// CreateGroup indicates an expected call of CreateGroup
func (mr *MockAuthServerMockRecorder) CreateGroup(ctx, group interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateGroup", reflect.TypeOf((*MockAuthServer)(nil).CreateGroup), ctx, group)
}

// UpdateGroup mocks base method
func (m *MockAuthServer) UpdateGroup(ctx context.Context, group *v1.ModifyUserGroup) *v1.Response {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateGroup", ctx, group)
	ret0, _ := ret[0].(*v1.Response)
	return ret0
}

// UpdateGroup indicates an expected call of UpdateGroup
func (mr *MockAuthServerMockRecorder) UpdateGroup(ctx, group interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateGroup", reflect.TypeOf((*MockAuthServer)(nil).UpdateGroup), ctx, group)
}

// DeleteGroups mocks base method
func (m *MockAuthServer) DeleteGroups(ctx context.Context, group []*v1.UserGroup) *v1.BatchWriteResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteGroups", ctx, group)
	ret0, _ := ret[0].(*v1.BatchWriteResponse)
	return ret0
}

// DeleteGroups indicates an expected call of DeleteGroups
func (mr *MockAuthServerMockRecorder) DeleteGroups(ctx, group interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteGroups", reflect.TypeOf((*MockAuthServer)(nil).DeleteGroups), ctx, group)
}

// GetGroups mocks base method
func (m *MockAuthServer) GetGroups(ctx context.Context, query map[string]string) *v1.BatchQueryResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGroups", ctx, query)
	ret0, _ := ret[0].(*v1.BatchQueryResponse)
	return ret0
}

// GetGroups indicates an expected call of GetGroups
func (mr *MockAuthServerMockRecorder) GetGroups(ctx, query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGroups", reflect.TypeOf((*MockAuthServer)(nil).GetGroups), ctx, query)
}

// GetGroup mocks base method
func (m *MockAuthServer) GetGroup(ctx context.Context, req *v1.UserGroup) *v1.Response {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGroup", ctx, req)
	ret0, _ := ret[0].(*v1.Response)
	return ret0
}

// GetGroup indicates an expected call of GetGroup
func (mr *MockAuthServerMockRecorder) GetGroup(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGroup", reflect.TypeOf((*MockAuthServer)(nil).GetGroup), ctx, req)
}

// GetGroupToken mocks base method
func (m *MockAuthServer) GetGroupToken(ctx context.Context, group *v1.UserGroup) *v1.Response {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGroupToken", ctx, group)
	ret0, _ := ret[0].(*v1.Response)
	return ret0
}

// GetGroupToken indicates an expected call of GetGroupToken
func (mr *MockAuthServerMockRecorder) GetGroupToken(ctx, group interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGroupToken", reflect.TypeOf((*MockAuthServer)(nil).GetGroupToken), ctx, group)
}

// UpdateGroupToken mocks base method
func (m *MockAuthServer) UpdateGroupToken(ctx context.Context, group *v1.UserGroup) *v1.Response {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateGroupToken", ctx, group)
	ret0, _ := ret[0].(*v1.Response)
	return ret0
}

// UpdateGroupToken indicates an expected call of UpdateGroupToken
func (mr *MockAuthServerMockRecorder) UpdateGroupToken(ctx, group interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateGroupToken", reflect.TypeOf((*MockAuthServer)(nil).UpdateGroupToken), ctx, group)
}

// ResetGroupToken mocks base method
func (m *MockAuthServer) ResetGroupToken(ctx context.Context, group *v1.UserGroup) *v1.Response {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResetGroupToken", ctx, group)
	ret0, _ := ret[0].(*v1.Response)
	return ret0
}

// ResetGroupToken indicates an expected call of ResetGroupToken
func (mr *MockAuthServerMockRecorder) ResetGroupToken(ctx, group interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResetGroupToken", reflect.TypeOf((*MockAuthServer)(nil).ResetGroupToken), ctx, group)
}

// CreateStrategy mocks base method
func (m *MockAuthServer) CreateStrategy(ctx context.Context, strategy *v1.AuthStrategy) *v1.Response {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateStrategy", ctx, strategy)
	ret0, _ := ret[0].(*v1.Response)
	return ret0
}

// CreateStrategy indicates an expected call of CreateStrategy
func (mr *MockAuthServerMockRecorder) CreateStrategy(ctx, strategy interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateStrategy", reflect.TypeOf((*MockAuthServer)(nil).CreateStrategy), ctx, strategy)
}

// UpdateStrategy mocks base method
func (m *MockAuthServer) UpdateStrategy(ctx context.Context, strategy *v1.ModifyAuthStrategy) *v1.Response {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateStrategy", ctx, strategy)
	ret0, _ := ret[0].(*v1.Response)
	return ret0
}

// UpdateStrategy indicates an expected call of UpdateStrategy
func (mr *MockAuthServerMockRecorder) UpdateStrategy(ctx, strategy interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStrategy", reflect.TypeOf((*MockAuthServer)(nil).UpdateStrategy), ctx, strategy)
}

// DeleteStrategies mocks base method
func (m *MockAuthServer) DeleteStrategies(ctx context.Context, reqs []*v1.AuthStrategy) *v1.BatchWriteResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteStrategies", ctx, reqs)
	ret0, _ := ret[0].(*v1.BatchWriteResponse)
	return ret0
}

// DeleteStrategies indicates an expected call of DeleteStrategies
func (mr *MockAuthServerMockRecorder) DeleteStrategies(ctx, reqs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteStrategies", reflect.TypeOf((*MockAuthServer)(nil).DeleteStrategies), ctx, reqs)
}

// GetStrategies mocks base method
func (m *MockAuthServer) GetStrategies(ctx context.Context, query map[string]string) *v1.BatchQueryResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStrategies", ctx, query)
	ret0, _ := ret[0].(*v1.BatchQueryResponse)
	return ret0
}

// GetStrategies indicates an expected call of GetStrategies
func (mr *MockAuthServerMockRecorder) GetStrategies(ctx, query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStrategies", reflect.TypeOf((*MockAuthServer)(nil).GetStrategies), ctx, query)
}

// GetStrategy mocks base method
func (m *MockAuthServer) GetStrategy(ctx context.Context, strategy *v1.AuthStrategy) *v1.Response {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStrategy", ctx, strategy)
	ret0, _ := ret[0].(*v1.Response)
	return ret0
}

// GetStrategy indicates an expected call of GetStrategy
func (mr *MockAuthServerMockRecorder) GetStrategy(ctx, strategy interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStrategy", reflect.TypeOf((*MockAuthServer)(nil).GetStrategy), ctx, strategy)
}

// MockAuthChecker is a mock of AuthChecker interface
type MockAuthChecker struct {
	ctrl     *gomock.Controller
	recorder *MockAuthCheckerMockRecorder
}

// MockAuthCheckerMockRecorder is the mock recorder for MockAuthChecker
type MockAuthCheckerMockRecorder struct {
	mock *MockAuthChecker
}

// NewMockAuthChecker creates a new mock instance
func NewMockAuthChecker(ctrl *gomock.Controller) *MockAuthChecker {
	mock := &MockAuthChecker{ctrl: ctrl}
	mock.recorder = &MockAuthCheckerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAuthChecker) EXPECT() *MockAuthCheckerMockRecorder {
	return m.recorder
}

// Initialize mocks base method
func (m *MockAuthChecker) Initialize(options *auth.Config, cacheMgn *cache.NamingCache) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Initialize", options, cacheMgn)
	ret0, _ := ret[0].(error)
	return ret0
}

// Initialize indicates an expected call of Initialize
func (mr *MockAuthCheckerMockRecorder) Initialize(options, cacheMgn interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Initialize", reflect.TypeOf((*MockAuthChecker)(nil).Initialize), options, cacheMgn)
}

// VerifyToken mocks base method
func (m *MockAuthChecker) VerifyToken(preCtx *model.AcquireContext) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyToken", preCtx)
	ret0, _ := ret[0].(error)
	return ret0
}

// VerifyToken indicates an expected call of VerifyToken
func (mr *MockAuthCheckerMockRecorder) VerifyToken(preCtx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyToken", reflect.TypeOf((*MockAuthChecker)(nil).VerifyToken), preCtx)
}

// CheckPermission mocks base method
func (m *MockAuthChecker) CheckPermission(preCtx *model.AcquireContext) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckPermission", preCtx)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckPermission indicates an expected call of CheckPermission
func (mr *MockAuthCheckerMockRecorder) CheckPermission(preCtx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckPermission", reflect.TypeOf((*MockAuthChecker)(nil).CheckPermission), preCtx)
}

// IsOpenAuth mocks base method
func (m *MockAuthChecker) IsOpenAuth() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsOpenAuth")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsOpenAuth indicates an expected call of IsOpenAuth
func (mr *MockAuthCheckerMockRecorder) IsOpenAuth() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsOpenAuth", reflect.TypeOf((*MockAuthChecker)(nil).IsOpenAuth))
}

// MockUserOperator is a mock of UserOperator interface
type MockUserOperator struct {
	ctrl     *gomock.Controller
	recorder *MockUserOperatorMockRecorder
}

// MockUserOperatorMockRecorder is the mock recorder for MockUserOperator
type MockUserOperatorMockRecorder struct {
	mock *MockUserOperator
}

// NewMockUserOperator creates a new mock instance
func NewMockUserOperator(ctrl *gomock.Controller) *MockUserOperator {
	mock := &MockUserOperator{ctrl: ctrl}
	mock.recorder = &MockUserOperatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUserOperator) EXPECT() *MockUserOperatorMockRecorder {
	return m.recorder
}

// CreateUsers mocks base method
func (m *MockUserOperator) CreateUsers(ctx context.Context, users []*v1.User) *v1.BatchWriteResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUsers", ctx, users)
	ret0, _ := ret[0].(*v1.BatchWriteResponse)
	return ret0
}

// CreateUsers indicates an expected call of CreateUsers
func (mr *MockUserOperatorMockRecorder) CreateUsers(ctx, users interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUsers", reflect.TypeOf((*MockUserOperator)(nil).CreateUsers), ctx, users)
}

// UpdateUser mocks base method
func (m *MockUserOperator) UpdateUser(ctx context.Context, user *v1.User) *v1.Response {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", ctx, user)
	ret0, _ := ret[0].(*v1.Response)
	return ret0
}

// UpdateUser indicates an expected call of UpdateUser
func (mr *MockUserOperatorMockRecorder) UpdateUser(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockUserOperator)(nil).UpdateUser), ctx, user)
}

// DeleteUsers mocks base method
func (m *MockUserOperator) DeleteUsers(ctx context.Context, users []*v1.User) *v1.BatchWriteResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUsers", ctx, users)
	ret0, _ := ret[0].(*v1.BatchWriteResponse)
	return ret0
}

// DeleteUsers indicates an expected call of DeleteUsers
func (mr *MockUserOperatorMockRecorder) DeleteUsers(ctx, users interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUsers", reflect.TypeOf((*MockUserOperator)(nil).DeleteUsers), ctx, users)
}

// GetUsers mocks base method
func (m *MockUserOperator) GetUsers(ctx context.Context, query map[string]string) *v1.BatchQueryResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUsers", ctx, query)
	ret0, _ := ret[0].(*v1.BatchQueryResponse)
	return ret0
}

// GetUsers indicates an expected call of GetUsers
func (mr *MockUserOperatorMockRecorder) GetUsers(ctx, query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsers", reflect.TypeOf((*MockUserOperator)(nil).GetUsers), ctx, query)
}

// GetUserToken mocks base method
func (m *MockUserOperator) GetUserToken(ctx context.Context, user *v1.User) *v1.Response {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserToken", ctx, user)
	ret0, _ := ret[0].(*v1.Response)
	return ret0
}

// GetUserToken indicates an expected call of GetUserToken
func (mr *MockUserOperatorMockRecorder) GetUserToken(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserToken", reflect.TypeOf((*MockUserOperator)(nil).GetUserToken), ctx, user)
}

// UpdateUserToken mocks base method
func (m *MockUserOperator) UpdateUserToken(ctx context.Context, user *v1.User) *v1.Response {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserToken", ctx, user)
	ret0, _ := ret[0].(*v1.Response)
	return ret0
}

// UpdateUserToken indicates an expected call of UpdateUserToken
func (mr *MockUserOperatorMockRecorder) UpdateUserToken(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserToken", reflect.TypeOf((*MockUserOperator)(nil).UpdateUserToken), ctx, user)
}

// ResetUserToken mocks base method
func (m *MockUserOperator) ResetUserToken(ctx context.Context, user *v1.User) *v1.Response {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResetUserToken", ctx, user)
	ret0, _ := ret[0].(*v1.Response)
	return ret0
}

// ResetUserToken indicates an expected call of ResetUserToken
func (mr *MockUserOperatorMockRecorder) ResetUserToken(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResetUserToken", reflect.TypeOf((*MockUserOperator)(nil).ResetUserToken), ctx, user)
}

// MockGroupOperator is a mock of GroupOperator interface
type MockGroupOperator struct {
	ctrl     *gomock.Controller
	recorder *MockGroupOperatorMockRecorder
}

// MockGroupOperatorMockRecorder is the mock recorder for MockGroupOperator
type MockGroupOperatorMockRecorder struct {
	mock *MockGroupOperator
}

// NewMockGroupOperator creates a new mock instance
func NewMockGroupOperator(ctrl *gomock.Controller) *MockGroupOperator {
	mock := &MockGroupOperator{ctrl: ctrl}
	mock.recorder = &MockGroupOperatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockGroupOperator) EXPECT() *MockGroupOperatorMockRecorder {
	return m.recorder
}

// CreateGroup mocks base method
func (m *MockGroupOperator) CreateGroup(ctx context.Context, group *v1.UserGroup) *v1.Response {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateGroup", ctx, group)
	ret0, _ := ret[0].(*v1.Response)
	return ret0
}

// CreateGroup indicates an expected call of CreateGroup
func (mr *MockGroupOperatorMockRecorder) CreateGroup(ctx, group interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateGroup", reflect.TypeOf((*MockGroupOperator)(nil).CreateGroup), ctx, group)
}

// UpdateGroup mocks base method
func (m *MockGroupOperator) UpdateGroup(ctx context.Context, group *v1.ModifyUserGroup) *v1.Response {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateGroup", ctx, group)
	ret0, _ := ret[0].(*v1.Response)
	return ret0
}

// UpdateGroup indicates an expected call of UpdateGroup
func (mr *MockGroupOperatorMockRecorder) UpdateGroup(ctx, group interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateGroup", reflect.TypeOf((*MockGroupOperator)(nil).UpdateGroup), ctx, group)
}

// DeleteGroups mocks base method
func (m *MockGroupOperator) DeleteGroups(ctx context.Context, group []*v1.UserGroup) *v1.BatchWriteResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteGroups", ctx, group)
	ret0, _ := ret[0].(*v1.BatchWriteResponse)
	return ret0
}

// DeleteGroups indicates an expected call of DeleteGroups
func (mr *MockGroupOperatorMockRecorder) DeleteGroups(ctx, group interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteGroups", reflect.TypeOf((*MockGroupOperator)(nil).DeleteGroups), ctx, group)
}

// GetGroups mocks base method
func (m *MockGroupOperator) GetGroups(ctx context.Context, query map[string]string) *v1.BatchQueryResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGroups", ctx, query)
	ret0, _ := ret[0].(*v1.BatchQueryResponse)
	return ret0
}

// GetGroups indicates an expected call of GetGroups
func (mr *MockGroupOperatorMockRecorder) GetGroups(ctx, query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGroups", reflect.TypeOf((*MockGroupOperator)(nil).GetGroups), ctx, query)
}

// GetGroup mocks base method
func (m *MockGroupOperator) GetGroup(ctx context.Context, req *v1.UserGroup) *v1.Response {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGroup", ctx, req)
	ret0, _ := ret[0].(*v1.Response)
	return ret0
}

// GetGroup indicates an expected call of GetGroup
func (mr *MockGroupOperatorMockRecorder) GetGroup(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGroup", reflect.TypeOf((*MockGroupOperator)(nil).GetGroup), ctx, req)
}

// GetGroupToken mocks base method
func (m *MockGroupOperator) GetGroupToken(ctx context.Context, group *v1.UserGroup) *v1.Response {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGroupToken", ctx, group)
	ret0, _ := ret[0].(*v1.Response)
	return ret0
}

// GetGroupToken indicates an expected call of GetGroupToken
func (mr *MockGroupOperatorMockRecorder) GetGroupToken(ctx, group interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGroupToken", reflect.TypeOf((*MockGroupOperator)(nil).GetGroupToken), ctx, group)
}

// UpdateGroupToken mocks base method
func (m *MockGroupOperator) UpdateGroupToken(ctx context.Context, group *v1.UserGroup) *v1.Response {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateGroupToken", ctx, group)
	ret0, _ := ret[0].(*v1.Response)
	return ret0
}

// UpdateGroupToken indicates an expected call of UpdateGroupToken
func (mr *MockGroupOperatorMockRecorder) UpdateGroupToken(ctx, group interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateGroupToken", reflect.TypeOf((*MockGroupOperator)(nil).UpdateGroupToken), ctx, group)
}

// ResetGroupToken mocks base method
func (m *MockGroupOperator) ResetGroupToken(ctx context.Context, group *v1.UserGroup) *v1.Response {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResetGroupToken", ctx, group)
	ret0, _ := ret[0].(*v1.Response)
	return ret0
}

// ResetGroupToken indicates an expected call of ResetGroupToken
func (mr *MockGroupOperatorMockRecorder) ResetGroupToken(ctx, group interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResetGroupToken", reflect.TypeOf((*MockGroupOperator)(nil).ResetGroupToken), ctx, group)
}

// MockStrategyOperator is a mock of StrategyOperator interface
type MockStrategyOperator struct {
	ctrl     *gomock.Controller
	recorder *MockStrategyOperatorMockRecorder
}

// MockStrategyOperatorMockRecorder is the mock recorder for MockStrategyOperator
type MockStrategyOperatorMockRecorder struct {
	mock *MockStrategyOperator
}

// NewMockStrategyOperator creates a new mock instance
func NewMockStrategyOperator(ctrl *gomock.Controller) *MockStrategyOperator {
	mock := &MockStrategyOperator{ctrl: ctrl}
	mock.recorder = &MockStrategyOperatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStrategyOperator) EXPECT() *MockStrategyOperatorMockRecorder {
	return m.recorder
}

// CreateStrategy mocks base method
func (m *MockStrategyOperator) CreateStrategy(ctx context.Context, strategy *v1.AuthStrategy) *v1.Response {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateStrategy", ctx, strategy)
	ret0, _ := ret[0].(*v1.Response)
	return ret0
}

// CreateStrategy indicates an expected call of CreateStrategy
func (mr *MockStrategyOperatorMockRecorder) CreateStrategy(ctx, strategy interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateStrategy", reflect.TypeOf((*MockStrategyOperator)(nil).CreateStrategy), ctx, strategy)
}

// UpdateStrategy mocks base method
func (m *MockStrategyOperator) UpdateStrategy(ctx context.Context, strategy *v1.ModifyAuthStrategy) *v1.Response {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateStrategy", ctx, strategy)
	ret0, _ := ret[0].(*v1.Response)
	return ret0
}

// UpdateStrategy indicates an expected call of UpdateStrategy
func (mr *MockStrategyOperatorMockRecorder) UpdateStrategy(ctx, strategy interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStrategy", reflect.TypeOf((*MockStrategyOperator)(nil).UpdateStrategy), ctx, strategy)
}

// DeleteStrategies mocks base method
func (m *MockStrategyOperator) DeleteStrategies(ctx context.Context, reqs []*v1.AuthStrategy) *v1.BatchWriteResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteStrategies", ctx, reqs)
	ret0, _ := ret[0].(*v1.BatchWriteResponse)
	return ret0
}

// DeleteStrategies indicates an expected call of DeleteStrategies
func (mr *MockStrategyOperatorMockRecorder) DeleteStrategies(ctx, reqs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteStrategies", reflect.TypeOf((*MockStrategyOperator)(nil).DeleteStrategies), ctx, reqs)
}

// GetStrategies mocks base method
func (m *MockStrategyOperator) GetStrategies(ctx context.Context, query map[string]string) *v1.BatchQueryResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStrategies", ctx, query)
	ret0, _ := ret[0].(*v1.BatchQueryResponse)
	return ret0
}

// GetStrategies indicates an expected call of GetStrategies
func (mr *MockStrategyOperatorMockRecorder) GetStrategies(ctx, query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStrategies", reflect.TypeOf((*MockStrategyOperator)(nil).GetStrategies), ctx, query)
}

// GetStrategy mocks base method
func (m *MockStrategyOperator) GetStrategy(ctx context.Context, strategy *v1.AuthStrategy) *v1.Response {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStrategy", ctx, strategy)
	ret0, _ := ret[0].(*v1.Response)
	return ret0
}

// GetStrategy indicates an expected call of GetStrategy
func (mr *MockStrategyOperatorMockRecorder) GetStrategy(ctx, strategy interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStrategy", reflect.TypeOf((*MockStrategyOperator)(nil).GetStrategy), ctx, strategy)
}

// MockAuthority is a mock of Authority interface
type MockAuthority struct {
	ctrl     *gomock.Controller
	recorder *MockAuthorityMockRecorder
}

// MockAuthorityMockRecorder is the mock recorder for MockAuthority
type MockAuthorityMockRecorder struct {
	mock *MockAuthority
}

// NewMockAuthority creates a new mock instance
func NewMockAuthority(ctrl *gomock.Controller) *MockAuthority {
	mock := &MockAuthority{ctrl: ctrl}
	mock.recorder = &MockAuthorityMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAuthority) EXPECT() *MockAuthorityMockRecorder {
	return m.recorder
}

// VerifyToken mocks base method
func (m *MockAuthority) VerifyToken(actualToken string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyToken", actualToken)
	ret0, _ := ret[0].(bool)
	return ret0
}

// VerifyToken indicates an expected call of VerifyToken
func (mr *MockAuthorityMockRecorder) VerifyToken(actualToken interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyToken", reflect.TypeOf((*MockAuthority)(nil).VerifyToken), actualToken)
}

// VerifyNamespace mocks base method
func (m *MockAuthority) VerifyNamespace(expectToken, actualToken string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyNamespace", expectToken, actualToken)
	ret0, _ := ret[0].(bool)
	return ret0
}

// VerifyNamespace indicates an expected call of VerifyNamespace
func (mr *MockAuthorityMockRecorder) VerifyNamespace(expectToken, actualToken interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyNamespace", reflect.TypeOf((*MockAuthority)(nil).VerifyNamespace), expectToken, actualToken)
}

// VerifyService mocks base method
func (m *MockAuthority) VerifyService(expectToken, actualToken string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyService", expectToken, actualToken)
	ret0, _ := ret[0].(bool)
	return ret0
}

// VerifyService indicates an expected call of VerifyService
func (mr *MockAuthorityMockRecorder) VerifyService(expectToken, actualToken interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyService", reflect.TypeOf((*MockAuthority)(nil).VerifyService), expectToken, actualToken)
}

// VerifyInstance mocks base method
func (m *MockAuthority) VerifyInstance(expectToken, actualToken string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyInstance", expectToken, actualToken)
	ret0, _ := ret[0].(bool)
	return ret0
}

// VerifyInstance indicates an expected call of VerifyInstance
func (mr *MockAuthorityMockRecorder) VerifyInstance(expectToken, actualToken interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyInstance", reflect.TypeOf((*MockAuthority)(nil).VerifyInstance), expectToken, actualToken)
}

// VerifyRule mocks base method
func (m *MockAuthority) VerifyRule(expectToken, actualToken string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyRule", expectToken, actualToken)
	ret0, _ := ret[0].(bool)
	return ret0
}

// VerifyRule indicates an expected call of VerifyRule
func (mr *MockAuthorityMockRecorder) VerifyRule(expectToken, actualToken interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyRule", reflect.TypeOf((*MockAuthority)(nil).VerifyRule), expectToken, actualToken)
}

// VerifyPlatform mocks base method
func (m *MockAuthority) VerifyPlatform(expectToken, actualToken string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyPlatform", expectToken, actualToken)
	ret0, _ := ret[0].(bool)
	return ret0
}

// VerifyPlatform indicates an expected call of VerifyPlatform
func (mr *MockAuthorityMockRecorder) VerifyPlatform(expectToken, actualToken interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyPlatform", reflect.TypeOf((*MockAuthority)(nil).VerifyPlatform), expectToken, actualToken)
}

// VerifyMesh mocks base method
func (m *MockAuthority) VerifyMesh(expectToken, actualToken string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyMesh", expectToken, actualToken)
	ret0, _ := ret[0].(bool)
	return ret0
}

// VerifyMesh indicates an expected call of VerifyMesh
func (mr *MockAuthorityMockRecorder) VerifyMesh(expectToken, actualToken interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyMesh", reflect.TypeOf((*MockAuthority)(nil).VerifyMesh), expectToken, actualToken)
}
