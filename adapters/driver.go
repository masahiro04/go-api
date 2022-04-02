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

func NewDriver(drivers Driver) Driver {
	return Driver{
		Logger:          drivers.Logger,
		Presenter:       drivers.Presenter,
		BlogDao:         drivers.BlogDao,
		UserDao:         drivers.UserDao,
		FirebaseHandler: drivers.FirebaseHandler,
		DBTransaction:   drivers.DBTransaction,
	}
}
