package usecases

import (
	"go-api/domains"
	"time"

	"firebase.google.com/go/auth"
	"gorm.io/gorm"
	// "firebase.google.com/go/auth"
)

// interactor : the struct that will have as properties all the IMPLEMENTED interfaces
// in order to provide them to its methods : the use cases and implement the Handler interface
type Repository struct {
	logger          LoggerRepository
	presenter       PresenterRepository
	blogDao         BlogDaoRepository
	userDao         UserDaoRepository
	firebaseHandler FirebaseHandlerRepository
	// validator     Validator
	dbTransaction DBTransactionRepository
}

// Logger : only used to log stuff
type LoggerRepository interface {
	Log(...interface{})
}

type PresenterRepository interface {
	Raise(errorKind domains.ErrorKinds, err error)
	Present() error

	CreateSignUp(user *domains.User)
	GetBlog(blog *domains.Blog)
	CreateBlog(blog *domains.Blog)
	GetBlogs(blogs *domains.Blogs)

	GetUsers(users *domains.Users)
	GetUser(user *domains.User)
}

type ValidatorRepository interface {
	Validate(targetStruct interface{}) error
}

type DBTransactionRepository interface {
	WithTx(runner func(tx *gorm.DB) error) error
}

type FirebaseHandlerRepository interface {
	VerifyIDToken(idToken string) (token *auth.Token, err error)
	// GetUser(uuId string) (user *domain.User, err error)
	CreateUser(user domains.User) (uuId *string, err error)
	// UpdateUser(uuId string, updateParams *domain.UserUpdatableProperty) error
	// ActivateUser(uuId string) error
	// DisableUser(uuId string) error
	DeleteUser(uuId string) error
	EmailVerificationLinkWithSettings(email string) (*string, error)
	EmailSignInLink(email string) (*string, error)
	SessionCookie(idToken string, expiresIn time.Duration) (*string, error)
	VerifySessionCookieAndCheckRevoked(sessionCookie string) (uuid *string, err error)
	RevokeRefreshTokens(uuId string) error
}

type BlogDaoRepository interface {
	GetAll() (*domains.Blogs, error)
	GetById(id int) (*domains.Blog, error)
	Create(blog domains.Blog) (*domains.Blog, error)
	// CreateTx(company domains.Blog, tx *sql.Tx) (*domains.Blog, error)
	Update(id int, company domains.Blog) (*domains.Blog, error)
	Delete(id int) error
}

type UserDaoRepository interface {
	GetAll() (*domains.Users, error)
	GetById(id int) (*domains.User, error)
	Create(user domains.User) (*domains.User, error)
	CreateTx(user domains.User, tx *gorm.DB) (*domains.User, error)
	Update(id int, user domains.User) (*domains.User, error)
	Delete(id int) error
}
