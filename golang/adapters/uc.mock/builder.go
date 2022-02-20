//go:build !netgo
// +build !netgo

package mock

import (
	"log"

	uc "clean_architecture/golang/usecases"

	"github.com/golang/mock/gomock"
)

// Interactor : is used in order to update its properties accordingly to each test conditions
type Interactor struct {
	Logger          *MockLogger
	Presenter       *MockPresenter
	BlogDao         *MockBlogDao
	UserDao         *MockUserDao
	FirebaseHandler *MockFirebaseHandler
	// Validator *MockValidator
	// DBTransaction *MockDBTransaction
}

type Tester struct {
	Calls      func(*Interactor)
	ShouldPass bool
}

type SimpleLogger struct{}

func (SimpleLogger) Log(logs ...interface{}) {
	log.Println(logs...)
}

//NewMockedInteractor : the Interactor constructor
func NewMockedInteractor(mockCtrl *gomock.Controller) Interactor {
	return Interactor{
		Logger:          NewMockLogger(mockCtrl),
		Presenter:       NewMockPresenter(mockCtrl),
		BlogDao:         NewMockBlogDao(mockCtrl),
		UserDao:         NewMockUserDao(mockCtrl),
		FirebaseHandler: NewMockFirebaseHandler(mockCtrl),
		// Validator: NewMockValidator(mockCtrl),
		// DBTransaction: NewMockDBTransaction(mockCtrl),
	}
}

//GetUCHandler : returns a uc.interactor in order to call its methods aka the use cases to test
func (i Interactor) GetUCHandler() uc.Handler {
	return uc.HandlerConstructor{
		Logger:          i.Logger,
		Presenter:       i.Presenter,
		BlogDao:         i.BlogDao,
		UserDao:         i.UserDao,
		FirebaseHandler: i.FirebaseHandler,
		// Validator: i.Validator,
		// DBTransaction: i.DBTransaction,
	}.New()
}
