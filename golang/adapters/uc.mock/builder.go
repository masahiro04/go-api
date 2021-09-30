//go:build !netgo
// +build !netgo

package mock

import (
	"log"

	"clean_architecture/golang/usecases"
	"github.com/golang/mock/gomock"
)

// Interactor : is used in order to update its properties accordingly to each test conditions
type Interactor struct {
	Logger        *MockLogger
	Presenter     *MockPresenter
	BlogRW        *MockBlogRW
	Validator     *MockValidator
	DBTransaction *MockDBTransaction
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
		Logger:        NewMockLogger(mockCtrl),
		Presenter:     NewMockPresenter(mockCtrl),
		BlogRW:        NewMockBlogRW(mockCtrl),
		Validator:     NewMockValidator(mockCtrl),
		DBTransaction: NewMockDBTransaction(mockCtrl),
	}
}

//GetUCHandler : returns a uc.interactor in order to call its methods aka the use cases to test
func (i Interactor) GetUCHandler() uc.Handler {
	return uc.HandlerConstructor{
		Logger:        i.Logger,
		Presenter:     i.Presenter,
		BlogRW:        i.BlogRW,
		Validator:     i.Validator,
		DBTransaction: i.DBTransaction,
	}.New()
}
