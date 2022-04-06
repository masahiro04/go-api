//go:build !netgo
// +build !netgo

package mock

import (
	"github.com/golang/mock/gomock"
)

// Interactor : is used in order to update its properties accordingly to each test conditions
type Driver struct {
	Logger    *MockLogger
	Presenter *MockPresenter
	BlogDao   *MockBlogDao
	UserDao   *MockUserDao
	// Validator *MockValidator
	// DBTransaction *MockDBTransaction
}

type Tester struct {
	Calls      func(*Driver)
	ShouldPass bool
}

// type SimpleLogger struct{}

// func (SimpleLogger) Log(logs ...interface{}) {
// 	log.Println(logs...)
// }

//NewMockedInteractor : the Interactor constructor
func NewMockDriver(mockCtrl *gomock.Controller) Driver {
	return Driver{
		Logger:    NewMockLogger(mockCtrl),
		Presenter: NewMockPresenter(mockCtrl),
		BlogDao:   NewMockBlogDao(mockCtrl),
		UserDao:   NewMockUserDao(mockCtrl),
		// Validator: NewMockValidator(mockCtrl),
		// DBTransaction: NewMockDBTransaction(mockCtrl),
	}
}

//GetUCHandler : returns a uc.interactor in order to call its methods aka the use cases to test
// func (d Driver) GetUCHandler() usecases.Handler {
// 	return usecases.HandlerConstructor{
// 		Logger:    i.Logger,
// 		Presenter: i.Presenter,
// 		BlogDao:   i.BlogDao,
// 		UserDao:   i.UserDao,
// 		// Validator: i.Validator,
// 		// DBTransaction: i.DBTransaction,
// 	}.New()
// }
