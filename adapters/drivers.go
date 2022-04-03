package adapters

import "go-api/domains"

type Driver struct {
	Logger domains.Logger
	// Presenter       domains.PresenterRepository
	BlogDao         domains.BlogRepository
	UserDao         domains.UserRepository
	FirebaseHandler domains.FirebaseHandlerRepository
	DBTransaction   domains.DBTransactionRepository
}

func NewDriver(logger domains.Logger, blogDao domains.BlogRepository, userDao domains.UserRepository, firebaseHandler domains.FirebaseHandlerRepository, dbTransaction domains.DBTransactionRepository) *Driver {
	return &Driver{
		Logger: logger,
		// Presenter:       presenter,
		BlogDao:         blogDao,
		UserDao:         userDao,
		FirebaseHandler: firebaseHandler,
		DBTransaction:   dbTransaction,
	}
}
