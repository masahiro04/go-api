package uc

import (
	"log"
)

type Handler interface {
	BlogLogic
	UserLogic
}

type BlogLogic interface {
	BlogGetAll(uc GetBlogsUseCase)
	BlogGet(uc GetBlogUseCase)
	BlogCreate(uc CreateBlogUseCase)
	BlogEdit(uc EditBlogUseCase)
	BlogDelete(uc DeleteBlogUseCase)
}

type UserLogic interface {
	UserGetAll(uc GetUsersUseCase)
	UserGet(uc GetUserUseCase)
	UserCreate(uc CreateUserUseCase)
	UserEdit(uc EditUserUseCase)
	UserDelete(uc DeleteUserUseCase)

	SignUp(uc SignUpUseCase)
}

type HandlerConstructor struct {
	Logger          Logger
	Presenter       Presenter
	BlogDao         BlogDao
	UserDao         UserDao
	FirebaseHandler FirebaseHandler
	// Validator Validator
	// DBTransaction DBTransaction
}

func (c HandlerConstructor) New() Handler {
	if c.Logger == nil {
		log.Fatal("missing Logger")
	}
	if c.FirebaseHandler == nil {
		log.Fatal("missing FirebaseHandler")
	}
	if c.BlogDao == nil {
		log.Fatal("missing BlogDao")
	}
	if c.UserDao == nil {
		log.Fatal("missing UserDao")
	}
	// if c.Validator == nil {
	// 	log.Fatal("missing Validator")
	// }
	// if c.DBTransaction == nil {
	// 	log.Fatal("missing DBTransaction")
	// }

	return interactor{
		logger:          c.Logger,
		presenter:       c.Presenter,
		firebaseHandler: c.FirebaseHandler,
		blogDao:         c.BlogDao,
		userDao:         c.UserDao,
		// validator: c.Validator,
		// dbTransaction: c.DBTransaction,
	}
}
