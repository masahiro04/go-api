package uc

import (
	"clean_architecture/golang/domains"
	// "firebase.google.com/go/auth"
)

// interactor : the struct that will have as properties all the IMPLEMENTED interfaces
// in order to provide them to its methods : the use cases and implement the Handler interface
type interactor struct {
	logger    Logger
	presenter Presenter
	blogDao   BlogDao
	userDao   UserDao
	// validator     Validator
	// dbTransaction DBTransaction
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
	GetBlogs(blogs *domains.Blogs)

	GetUsers(users *domains.Users)
	GetUser(user *domains.User)
}

type Validator interface {
	Validate(targetStruct interface{}) error
}

type DBTransaction interface {
	// TODO(okubo): かたをちゃんと入れる
	WithTx(runner func(tx interface{}) error) error
}

type BlogDao interface {
	GetAll() (*domains.Blogs, error)
	GetById(id int) (*domains.Blog, error)
	Create(blog domains.Blog) (*domains.Blog, error)
	// CreateTx(company domains.Blog, tx *sql.Tx) (*domains.Blog, error)
	Update(id int, company domains.Blog) (*domains.Blog, error)
	Delete(id int) error
}

type UserDao interface {
	GetAll() (*domains.Users, error)
	GetById(id int) (*domains.User, error)
	Create(user domains.User) (*domains.User, error)
	// CreateTx(user domains.User, tx *sql.Tx) (*domains.User, error)
	Update(id int, user domains.User) (*domains.User, error)
	Delete(id int) error
}
