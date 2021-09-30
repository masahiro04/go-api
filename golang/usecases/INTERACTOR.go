package uc

import (
	"database/sql"

	"clean_architecture/golang/domain"
)

// interactor : the struct that will have as properties all the IMPLEMENTED interfaces
// in order to provide them to its methods : the use cases and implement the Handler interface
type interactor struct {
	logger        Logger
	presenter     Presenter
	blogRW        BlogRW
	validator     Validator
	dbTransaction DBTransaction
}

// Logger : only used to log stuff
type Logger interface {
	Log(...interface{})
}

type Presenter interface {
	Raise(errorKind domain.ErrorKinds, err error)
	Present() error
	GetBlog(blog *domain.Blog)
	CreateBlog(blog *domain.Blog)
	GetBlogs(blogs domain.BlogCollection)
}

type Validator interface {
	Validate(targetStruct interface{}) error
}

type DBTransaction interface {
	WithTx(runner func(tx *sql.Tx) error) error
}

type BlogRW interface {
	GetAll() ([]*domain.Blog, error)
	GetById(id int) (*domain.Blog, error)
	Create(company domain.Blog) (*domain.Blog, error)
	CreateTx(company domain.Blog, tx *sql.Tx) (*domain.Blog, error)
	Update(id int, company domain.Blog) (*domain.Blog, error)
	Delete(id int) error
}
