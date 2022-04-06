package adapters

import "go-api/usecases"

type Driver struct {
	Logger usecases.Logger
	// Presenter       usecases.PresenterRepository
	BlogDao         usecases.BlogRepository
	UserDao         usecases.UserRepository
	FirebaseHandler usecases.FirebaseHandlerRepository
	DBTransaction   usecases.DBTransactionRepository
}

func NewDriver(logger usecases.Logger, blogDao usecases.BlogRepository, userDao usecases.UserRepository, firebaseHandler usecases.FirebaseHandlerRepository, dbTransaction usecases.DBTransactionRepository) *Driver {
	return &Driver{
		Logger: logger,
		// Presenter:       presenter,
		BlogDao:         blogDao,
		UserDao:         userDao,
		FirebaseHandler: firebaseHandler,
		DBTransaction:   dbTransaction,
	}
}
