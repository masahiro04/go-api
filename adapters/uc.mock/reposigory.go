// Code generated by MockGen. DO NOT EDIT.
// Source: ./domains/repository.go

// Package mock is a generated GoMock package.
package mock

import (
	models "go-api/domains/models"
	reflect "reflect"
	time "time"

	auth "firebase.google.com/go/auth"
	gomock "github.com/golang/mock/gomock"
	gorm "gorm.io/gorm"
)

// MockPresenterRepository is a mock of PresenterRepository interface.
type MockPresenterRepository struct {
	ctrl     *gomock.Controller
	recorder *MockPresenterRepositoryMockRecorder
}

// MockPresenterRepositoryMockRecorder is the mock recorder for MockPresenterRepository.
type MockPresenterRepositoryMockRecorder struct {
	mock *MockPresenterRepository
}

// NewMockPresenterRepository creates a new mock instance.
func NewMockPresenterRepository(ctrl *gomock.Controller) *MockPresenterRepository {
	mock := &MockPresenterRepository{ctrl: ctrl}
	mock.recorder = &MockPresenterRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPresenterRepository) EXPECT() *MockPresenterRepositoryMockRecorder {
	return m.recorder
}

// CreateBlog mocks base method.
func (m *MockPresenterRepository) CreateBlog(blog *models.Blog) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "CreateBlog", blog)
}

// CreateBlog indicates an expected call of CreateBlog.
func (mr *MockPresenterRepositoryMockRecorder) CreateBlog(blog interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBlog", reflect.TypeOf((*MockPresenterRepository)(nil).CreateBlog), blog)
}

// CreateSignUp mocks base method.
func (m *MockPresenterRepository) CreateSignUp(user *models.User) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "CreateSignUp", user)
}

// CreateSignUp indicates an expected call of CreateSignUp.
func (mr *MockPresenterRepositoryMockRecorder) CreateSignUp(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSignUp", reflect.TypeOf((*MockPresenterRepository)(nil).CreateSignUp), user)
}

// GetBlog mocks base method.
func (m *MockPresenterRepository) GetBlog(blog *models.Blog) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "GetBlog", blog)
}

// GetBlog indicates an expected call of GetBlog.
func (mr *MockPresenterRepositoryMockRecorder) GetBlog(blog interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlog", reflect.TypeOf((*MockPresenterRepository)(nil).GetBlog), blog)
}

// GetBlogs mocks base method.
func (m *MockPresenterRepository) GetBlogs(blogs *models.Blogs) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "GetBlogs", blogs)
}

// GetBlogs indicates an expected call of GetBlogs.
func (mr *MockPresenterRepositoryMockRecorder) GetBlogs(blogs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlogs", reflect.TypeOf((*MockPresenterRepository)(nil).GetBlogs), blogs)
}

// GetUser mocks base method.
func (m *MockPresenterRepository) GetUser(user *models.User) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "GetUser", user)
}

// GetUser indicates an expected call of GetUser.
func (mr *MockPresenterRepositoryMockRecorder) GetUser(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockPresenterRepository)(nil).GetUser), user)
}

// GetUsers mocks base method.
func (m *MockPresenterRepository) GetUsers(users *models.Users) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "GetUsers", users)
}

// GetUsers indicates an expected call of GetUsers.
func (mr *MockPresenterRepositoryMockRecorder) GetUsers(users interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsers", reflect.TypeOf((*MockPresenterRepository)(nil).GetUsers), users)
}

// Present mocks base method.
func (m *MockPresenterRepository) Present() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Present")
	ret0, _ := ret[0].(error)
	return ret0
}

// Present indicates an expected call of Present.
func (mr *MockPresenterRepositoryMockRecorder) Present() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Present", reflect.TypeOf((*MockPresenterRepository)(nil).Present))
}

// Raise mocks base method.
func (m *MockPresenterRepository) Raise(errorKind models.ErrorKinds, err error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Raise", errorKind, err)
}

// Raise indicates an expected call of Raise.
func (mr *MockPresenterRepositoryMockRecorder) Raise(errorKind, err interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Raise", reflect.TypeOf((*MockPresenterRepository)(nil).Raise), errorKind, err)
}

// MockValidatorRepository is a mock of ValidatorRepository interface.
type MockValidatorRepository struct {
	ctrl     *gomock.Controller
	recorder *MockValidatorRepositoryMockRecorder
}

// MockValidatorRepositoryMockRecorder is the mock recorder for MockValidatorRepository.
type MockValidatorRepositoryMockRecorder struct {
	mock *MockValidatorRepository
}

// NewMockValidatorRepository creates a new mock instance.
func NewMockValidatorRepository(ctrl *gomock.Controller) *MockValidatorRepository {
	mock := &MockValidatorRepository{ctrl: ctrl}
	mock.recorder = &MockValidatorRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockValidatorRepository) EXPECT() *MockValidatorRepositoryMockRecorder {
	return m.recorder
}

// Validate mocks base method.
func (m *MockValidatorRepository) Validate(targetStruct interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Validate", targetStruct)
	ret0, _ := ret[0].(error)
	return ret0
}

// Validate indicates an expected call of Validate.
func (mr *MockValidatorRepositoryMockRecorder) Validate(targetStruct interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Validate", reflect.TypeOf((*MockValidatorRepository)(nil).Validate), targetStruct)
}

// MockDBTransactionRepository is a mock of DBTransactionRepository interface.
type MockDBTransactionRepository struct {
	ctrl     *gomock.Controller
	recorder *MockDBTransactionRepositoryMockRecorder
}

// MockDBTransactionRepositoryMockRecorder is the mock recorder for MockDBTransactionRepository.
type MockDBTransactionRepositoryMockRecorder struct {
	mock *MockDBTransactionRepository
}

// NewMockDBTransactionRepository creates a new mock instance.
func NewMockDBTransactionRepository(ctrl *gomock.Controller) *MockDBTransactionRepository {
	mock := &MockDBTransactionRepository{ctrl: ctrl}
	mock.recorder = &MockDBTransactionRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDBTransactionRepository) EXPECT() *MockDBTransactionRepositoryMockRecorder {
	return m.recorder
}

// WithTx mocks base method.
func (m *MockDBTransactionRepository) WithTx(runner func(*gorm.DB) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithTx", runner)
	ret0, _ := ret[0].(error)
	return ret0
}

// WithTx indicates an expected call of WithTx.
func (mr *MockDBTransactionRepositoryMockRecorder) WithTx(runner interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithTx", reflect.TypeOf((*MockDBTransactionRepository)(nil).WithTx), runner)
}

// MockFirebaseHandlerRepository is a mock of FirebaseHandlerRepository interface.
type MockFirebaseHandlerRepository struct {
	ctrl     *gomock.Controller
	recorder *MockFirebaseHandlerRepositoryMockRecorder
}

// MockFirebaseHandlerRepositoryMockRecorder is the mock recorder for MockFirebaseHandlerRepository.
type MockFirebaseHandlerRepositoryMockRecorder struct {
	mock *MockFirebaseHandlerRepository
}

// NewMockFirebaseHandlerRepository creates a new mock instance.
func NewMockFirebaseHandlerRepository(ctrl *gomock.Controller) *MockFirebaseHandlerRepository {
	mock := &MockFirebaseHandlerRepository{ctrl: ctrl}
	mock.recorder = &MockFirebaseHandlerRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFirebaseHandlerRepository) EXPECT() *MockFirebaseHandlerRepositoryMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockFirebaseHandlerRepository) CreateUser(user models.User) (*string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", user)
	ret0, _ := ret[0].(*string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockFirebaseHandlerRepositoryMockRecorder) CreateUser(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockFirebaseHandlerRepository)(nil).CreateUser), user)
}

// DeleteUser mocks base method.
func (m *MockFirebaseHandlerRepository) DeleteUser(uuId string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", uuId)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockFirebaseHandlerRepositoryMockRecorder) DeleteUser(uuId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockFirebaseHandlerRepository)(nil).DeleteUser), uuId)
}

// EmailSignInLink mocks base method.
func (m *MockFirebaseHandlerRepository) EmailSignInLink(email string) (*string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EmailSignInLink", email)
	ret0, _ := ret[0].(*string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EmailSignInLink indicates an expected call of EmailSignInLink.
func (mr *MockFirebaseHandlerRepositoryMockRecorder) EmailSignInLink(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EmailSignInLink", reflect.TypeOf((*MockFirebaseHandlerRepository)(nil).EmailSignInLink), email)
}

// EmailVerificationLinkWithSettings mocks base method.
func (m *MockFirebaseHandlerRepository) EmailVerificationLinkWithSettings(email string) (*string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EmailVerificationLinkWithSettings", email)
	ret0, _ := ret[0].(*string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EmailVerificationLinkWithSettings indicates an expected call of EmailVerificationLinkWithSettings.
func (mr *MockFirebaseHandlerRepositoryMockRecorder) EmailVerificationLinkWithSettings(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EmailVerificationLinkWithSettings", reflect.TypeOf((*MockFirebaseHandlerRepository)(nil).EmailVerificationLinkWithSettings), email)
}

// RevokeRefreshTokens mocks base method.
func (m *MockFirebaseHandlerRepository) RevokeRefreshTokens(uuId string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RevokeRefreshTokens", uuId)
	ret0, _ := ret[0].(error)
	return ret0
}

// RevokeRefreshTokens indicates an expected call of RevokeRefreshTokens.
func (mr *MockFirebaseHandlerRepositoryMockRecorder) RevokeRefreshTokens(uuId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RevokeRefreshTokens", reflect.TypeOf((*MockFirebaseHandlerRepository)(nil).RevokeRefreshTokens), uuId)
}

// SessionCookie mocks base method.
func (m *MockFirebaseHandlerRepository) SessionCookie(idToken string, expiresIn time.Duration) (*string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SessionCookie", idToken, expiresIn)
	ret0, _ := ret[0].(*string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SessionCookie indicates an expected call of SessionCookie.
func (mr *MockFirebaseHandlerRepositoryMockRecorder) SessionCookie(idToken, expiresIn interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SessionCookie", reflect.TypeOf((*MockFirebaseHandlerRepository)(nil).SessionCookie), idToken, expiresIn)
}

// VerifyIDToken mocks base method.
func (m *MockFirebaseHandlerRepository) VerifyIDToken(idToken string) (*auth.Token, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyIDToken", idToken)
	ret0, _ := ret[0].(*auth.Token)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VerifyIDToken indicates an expected call of VerifyIDToken.
func (mr *MockFirebaseHandlerRepositoryMockRecorder) VerifyIDToken(idToken interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyIDToken", reflect.TypeOf((*MockFirebaseHandlerRepository)(nil).VerifyIDToken), idToken)
}

// VerifySessionCookieAndCheckRevoked mocks base method.
func (m *MockFirebaseHandlerRepository) VerifySessionCookieAndCheckRevoked(sessionCookie string) (*string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifySessionCookieAndCheckRevoked", sessionCookie)
	ret0, _ := ret[0].(*string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VerifySessionCookieAndCheckRevoked indicates an expected call of VerifySessionCookieAndCheckRevoked.
func (mr *MockFirebaseHandlerRepositoryMockRecorder) VerifySessionCookieAndCheckRevoked(sessionCookie interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifySessionCookieAndCheckRevoked", reflect.TypeOf((*MockFirebaseHandlerRepository)(nil).VerifySessionCookieAndCheckRevoked), sessionCookie)
}

// MockBlogRepository is a mock of BlogRepository interface.
type MockBlogRepository struct {
	ctrl     *gomock.Controller
	recorder *MockBlogRepositoryMockRecorder
}

// MockBlogRepositoryMockRecorder is the mock recorder for MockBlogRepository.
type MockBlogRepositoryMockRecorder struct {
	mock *MockBlogRepository
}

// NewMockBlogRepository creates a new mock instance.
func NewMockBlogRepository(ctrl *gomock.Controller) *MockBlogRepository {
	mock := &MockBlogRepository{ctrl: ctrl}
	mock.recorder = &MockBlogRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBlogRepository) EXPECT() *MockBlogRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockBlogRepository) Create(blog models.Blog) (*models.Blog, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", blog)
	ret0, _ := ret[0].(*models.Blog)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockBlogRepositoryMockRecorder) Create(blog interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockBlogRepository)(nil).Create), blog)
}

// Delete mocks base method.
func (m *MockBlogRepository) Delete(id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockBlogRepositoryMockRecorder) Delete(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockBlogRepository)(nil).Delete), id)
}

// GetAll mocks base method.
func (m *MockBlogRepository) GetAll() (*models.Blogs, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].(*models.Blogs)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockBlogRepositoryMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockBlogRepository)(nil).GetAll))
}

// GetById mocks base method.
func (m *MockBlogRepository) GetById(id int) (*models.Blog, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", id)
	ret0, _ := ret[0].(*models.Blog)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockBlogRepositoryMockRecorder) GetById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockBlogRepository)(nil).GetById), id)
}

// Update mocks base method.
func (m *MockBlogRepository) Update(id int, company models.Blog) (*models.Blog, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", id, company)
	ret0, _ := ret[0].(*models.Blog)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockBlogRepositoryMockRecorder) Update(id, company interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockBlogRepository)(nil).Update), id, company)
}

// MockUserRepository is a mock of UserRepository interface.
type MockUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepositoryMockRecorder
}

// MockUserRepositoryMockRecorder is the mock recorder for MockUserRepository.
type MockUserRepositoryMockRecorder struct {
	mock *MockUserRepository
}

// NewMockUserRepository creates a new mock instance.
func NewMockUserRepository(ctrl *gomock.Controller) *MockUserRepository {
	mock := &MockUserRepository{ctrl: ctrl}
	mock.recorder = &MockUserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRepository) EXPECT() *MockUserRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockUserRepository) Create(user models.User) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", user)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockUserRepositoryMockRecorder) Create(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockUserRepository)(nil).Create), user)
}

// CreateTx mocks base method.
func (m *MockUserRepository) CreateTx(user models.User, tx *gorm.DB) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTx", user, tx)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTx indicates an expected call of CreateTx.
func (mr *MockUserRepositoryMockRecorder) CreateTx(user, tx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTx", reflect.TypeOf((*MockUserRepository)(nil).CreateTx), user, tx)
}

// Delete mocks base method.
func (m *MockUserRepository) Delete(id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockUserRepositoryMockRecorder) Delete(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockUserRepository)(nil).Delete), id)
}

// GetAll mocks base method.
func (m *MockUserRepository) GetAll() (*models.Users, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].(*models.Users)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockUserRepositoryMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockUserRepository)(nil).GetAll))
}

// GetById mocks base method.
func (m *MockUserRepository) GetById(id int) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", id)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockUserRepositoryMockRecorder) GetById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockUserRepository)(nil).GetById), id)
}

// Update mocks base method.
func (m *MockUserRepository) Update(id int, user models.User) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", id, user)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockUserRepositoryMockRecorder) Update(id, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockUserRepository)(nil).Update), id, user)
}