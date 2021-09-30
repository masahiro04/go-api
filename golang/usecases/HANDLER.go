package uc

import (
	"log"
)

type Handler interface {
	BlogLogic
}

type BlogLogic interface {
	BlogGet(uc GetBlogUseCase)
	BlogGetAll(uc GetBlogsUseCase)
}

type HandlerConstructor struct {
	Logger        Logger
	Presenter     Presenter
	BlogRW        BlogRW
	Validator     Validator
	DBTransaction DBTransaction
}

func (c HandlerConstructor) New() Handler {
	if c.Logger == nil {
		log.Fatal("missing Logger")
	}
	if c.BlogRW == nil {
		log.Fatal("missing CompanyRW")
	}
	if c.Validator == nil {
		log.Fatal("missing Validator")
	}
	if c.DBTransaction == nil {
		log.Fatal("missing DBTransaction")
	}

	return interactor{
		logger:        c.Logger,
		presenter:     c.Presenter,
		blogRW:        c.BlogRW,
		validator:     c.Validator,
		dbTransaction: c.DBTransaction,
	}
}
