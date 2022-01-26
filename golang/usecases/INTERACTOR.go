package uc

import (
	"clean_architecture/golang/domains"
	"database/sql"
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
	Raise(errorKind domains.ErrorKinds, err error)
	Present() error
	GetBlog(blog *domains.Blog)
	CreateBlog(blog *domains.Blog)
	GetBlogs(blogs domains.BlogCollection)
}

type Validator interface {
	Validate(targetStruct interface{}) error
}

type DBTransaction interface {
	WithTx(runner func(tx *sql.Tx) error) error
}

type BlogRW interface {
	GetAll() ([]*domains.Blog, error)
	GetById(id int) (*domains.Blog, error)
	Create(company domains.Blog) (*domains.Blog, error)
	CreateTx(company domains.Blog, tx *sql.Tx) (*domains.Blog, error)
	Update(id int, company domains.Blog) (*domains.Blog, error)
	Delete(id int) error
}
