package adapters

import "go-api/domains"

type Driver struct {
	Logger          domains.LoggerRepository
	Presenter       domains.PresenterRepository
	BlogDao         domains.BlogRepository
	UserDao         domains.UserRepository
	FirebaseHandler domains.FirebaseHandlerRepository
	DBTransaction   domains.DBTransactionRepository
}

func NewDriver(driver Driver) Driver {
	return Driver{
		Logger:          driver.Logger,
		Presenter:       driver.Presenter,
		BlogDao:         driver.BlogDao,
		UserDao:         driver.UserDao,
		FirebaseHandler: driver.FirebaseHandler,
		DBTransaction:   driver.DBTransaction,
	}
}
