// Code generated by MockGen. DO NOT EDIT.
// Source: ./usecases/HANDLER.go

// Package mock is a generated GoMock package.
package mock

import (
	uc "clean_architecture/golang/usecases"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockHandler is a mock of Handler interface.
type MockHandler struct {
	ctrl     *gomock.Controller
	recorder *MockHandlerMockRecorder
}

// MockHandlerMockRecorder is the mock recorder for MockHandler.
type MockHandlerMockRecorder struct {
	mock *MockHandler
}

// NewMockHandler creates a new mock instance.
func NewMockHandler(ctrl *gomock.Controller) *MockHandler {
	mock := &MockHandler{ctrl: ctrl}
	mock.recorder = &MockHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHandler) EXPECT() *MockHandlerMockRecorder {
	return m.recorder
}

// BlogGet mocks base method.
func (m *MockHandler) BlogGet(uc uc.GetBlogUseCase) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "BlogGet", uc)
}

// BlogGet indicates an expected call of BlogGet.
func (mr *MockHandlerMockRecorder) BlogGet(uc interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlogGet", reflect.TypeOf((*MockHandler)(nil).BlogGet), uc)
}

// BlogGetAll mocks base method.
func (m *MockHandler) BlogGetAll(uc uc.GetBlogsUseCase) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "BlogGetAll", uc)
}

// BlogGetAll indicates an expected call of BlogGetAll.
func (mr *MockHandlerMockRecorder) BlogGetAll(uc interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlogGetAll", reflect.TypeOf((*MockHandler)(nil).BlogGetAll), uc)
}

// CompanyCreate mocks base method.
func (m *MockHandler) CompanyCreate(uc uc.CreateCompanyUseCase) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "CompanyCreate", uc)
}

// CompanyCreate indicates an expected call of CompanyCreate.
func (mr *MockHandlerMockRecorder) CompanyCreate(uc interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CompanyCreate", reflect.TypeOf((*MockHandler)(nil).CompanyCreate), uc)
}

// CompanyDelete mocks base method.
func (m *MockHandler) CompanyDelete(uc uc.DeleteCompanyUseCase) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "CompanyDelete", uc)
}

// CompanyDelete indicates an expected call of CompanyDelete.
func (mr *MockHandlerMockRecorder) CompanyDelete(uc interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CompanyDelete", reflect.TypeOf((*MockHandler)(nil).CompanyDelete), uc)
}

// CompanyEdit mocks base method.
func (m *MockHandler) CompanyEdit(uc uc.EditCompanyUseCase) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "CompanyEdit", uc)
}

// CompanyEdit indicates an expected call of CompanyEdit.
func (mr *MockHandlerMockRecorder) CompanyEdit(uc interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CompanyEdit", reflect.TypeOf((*MockHandler)(nil).CompanyEdit), uc)
}

// CompanyGet mocks base method.
func (m *MockHandler) CompanyGet(uc uc.GetCompanyUseCase) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "CompanyGet", uc)
}

// CompanyGet indicates an expected call of CompanyGet.
func (mr *MockHandlerMockRecorder) CompanyGet(uc interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CompanyGet", reflect.TypeOf((*MockHandler)(nil).CompanyGet), uc)
}

// CompanyGetAll mocks base method.
func (m *MockHandler) CompanyGetAll(uc uc.GetCompaniesUseCase) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "CompanyGetAll", uc)
}

// CompanyGetAll indicates an expected call of CompanyGetAll.
func (mr *MockHandlerMockRecorder) CompanyGetAll(uc interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CompanyGetAll", reflect.TypeOf((*MockHandler)(nil).CompanyGetAll), uc)
}

// MockCompanyLogic is a mock of CompanyLogic interface.
type MockCompanyLogic struct {
	ctrl     *gomock.Controller
	recorder *MockCompanyLogicMockRecorder
}

// MockCompanyLogicMockRecorder is the mock recorder for MockCompanyLogic.
type MockCompanyLogicMockRecorder struct {
	mock *MockCompanyLogic
}

// NewMockCompanyLogic creates a new mock instance.
func NewMockCompanyLogic(ctrl *gomock.Controller) *MockCompanyLogic {
	mock := &MockCompanyLogic{ctrl: ctrl}
	mock.recorder = &MockCompanyLogicMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCompanyLogic) EXPECT() *MockCompanyLogicMockRecorder {
	return m.recorder
}

// CompanyCreate mocks base method.
func (m *MockCompanyLogic) CompanyCreate(uc uc.CreateCompanyUseCase) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "CompanyCreate", uc)
}

// CompanyCreate indicates an expected call of CompanyCreate.
func (mr *MockCompanyLogicMockRecorder) CompanyCreate(uc interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CompanyCreate", reflect.TypeOf((*MockCompanyLogic)(nil).CompanyCreate), uc)
}

// CompanyDelete mocks base method.
func (m *MockCompanyLogic) CompanyDelete(uc uc.DeleteCompanyUseCase) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "CompanyDelete", uc)
}

// CompanyDelete indicates an expected call of CompanyDelete.
func (mr *MockCompanyLogicMockRecorder) CompanyDelete(uc interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CompanyDelete", reflect.TypeOf((*MockCompanyLogic)(nil).CompanyDelete), uc)
}

// CompanyEdit mocks base method.
func (m *MockCompanyLogic) CompanyEdit(uc uc.EditCompanyUseCase) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "CompanyEdit", uc)
}

// CompanyEdit indicates an expected call of CompanyEdit.
func (mr *MockCompanyLogicMockRecorder) CompanyEdit(uc interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CompanyEdit", reflect.TypeOf((*MockCompanyLogic)(nil).CompanyEdit), uc)
}

// CompanyGet mocks base method.
func (m *MockCompanyLogic) CompanyGet(uc uc.GetCompanyUseCase) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "CompanyGet", uc)
}

// CompanyGet indicates an expected call of CompanyGet.
func (mr *MockCompanyLogicMockRecorder) CompanyGet(uc interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CompanyGet", reflect.TypeOf((*MockCompanyLogic)(nil).CompanyGet), uc)
}

// CompanyGetAll mocks base method.
func (m *MockCompanyLogic) CompanyGetAll(uc uc.GetCompaniesUseCase) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "CompanyGetAll", uc)
}

// CompanyGetAll indicates an expected call of CompanyGetAll.
func (mr *MockCompanyLogicMockRecorder) CompanyGetAll(uc interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CompanyGetAll", reflect.TypeOf((*MockCompanyLogic)(nil).CompanyGetAll), uc)
}

// MockBlogLogic is a mock of BlogLogic interface.
type MockBlogLogic struct {
	ctrl     *gomock.Controller
	recorder *MockBlogLogicMockRecorder
}

// MockBlogLogicMockRecorder is the mock recorder for MockBlogLogic.
type MockBlogLogicMockRecorder struct {
	mock *MockBlogLogic
}

// NewMockBlogLogic creates a new mock instance.
func NewMockBlogLogic(ctrl *gomock.Controller) *MockBlogLogic {
	mock := &MockBlogLogic{ctrl: ctrl}
	mock.recorder = &MockBlogLogicMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBlogLogic) EXPECT() *MockBlogLogicMockRecorder {
	return m.recorder
}

// BlogGet mocks base method.
func (m *MockBlogLogic) BlogGet(uc uc.GetBlogUseCase) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "BlogGet", uc)
}

// BlogGet indicates an expected call of BlogGet.
func (mr *MockBlogLogicMockRecorder) BlogGet(uc interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlogGet", reflect.TypeOf((*MockBlogLogic)(nil).BlogGet), uc)
}

// BlogGetAll mocks base method.
func (m *MockBlogLogic) BlogGetAll(uc uc.GetBlogsUseCase) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "BlogGetAll", uc)
}

// BlogGetAll indicates an expected call of BlogGetAll.
func (mr *MockBlogLogicMockRecorder) BlogGetAll(uc interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlogGetAll", reflect.TypeOf((*MockBlogLogic)(nil).BlogGetAll), uc)
}
