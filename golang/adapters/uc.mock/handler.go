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

// BlogCreate mocks base method.
func (m *MockHandler) BlogCreate(uc uc.CreateBlogUseCase) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "BlogCreate", uc)
}

// BlogCreate indicates an expected call of BlogCreate.
func (mr *MockHandlerMockRecorder) BlogCreate(uc interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlogCreate", reflect.TypeOf((*MockHandler)(nil).BlogCreate), uc)
}

// BlogDelete mocks base method.
func (m *MockHandler) BlogDelete(uc uc.DeleteBlogUseCase) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "BlogDelete", uc)
}

// BlogDelete indicates an expected call of BlogDelete.
func (mr *MockHandlerMockRecorder) BlogDelete(uc interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlogDelete", reflect.TypeOf((*MockHandler)(nil).BlogDelete), uc)
}

// BlogEdit mocks base method.
func (m *MockHandler) BlogEdit(uc uc.EditBlogUseCase) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "BlogEdit", uc)
}

// BlogEdit indicates an expected call of BlogEdit.
func (mr *MockHandlerMockRecorder) BlogEdit(uc interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlogEdit", reflect.TypeOf((*MockHandler)(nil).BlogEdit), uc)
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

// BlogCreate mocks base method.
func (m *MockBlogLogic) BlogCreate(uc uc.CreateBlogUseCase) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "BlogCreate", uc)
}

// BlogCreate indicates an expected call of BlogCreate.
func (mr *MockBlogLogicMockRecorder) BlogCreate(uc interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlogCreate", reflect.TypeOf((*MockBlogLogic)(nil).BlogCreate), uc)
}

// BlogDelete mocks base method.
func (m *MockBlogLogic) BlogDelete(uc uc.DeleteBlogUseCase) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "BlogDelete", uc)
}

// BlogDelete indicates an expected call of BlogDelete.
func (mr *MockBlogLogicMockRecorder) BlogDelete(uc interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlogDelete", reflect.TypeOf((*MockBlogLogic)(nil).BlogDelete), uc)
}

// BlogEdit mocks base method.
func (m *MockBlogLogic) BlogEdit(uc uc.EditBlogUseCase) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "BlogEdit", uc)
}

// BlogEdit indicates an expected call of BlogEdit.
func (mr *MockBlogLogicMockRecorder) BlogEdit(uc interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlogEdit", reflect.TypeOf((*MockBlogLogic)(nil).BlogEdit), uc)
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
