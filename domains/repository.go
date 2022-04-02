package domains

import (
	"go-api/domains/models"
	"time"

	"firebase.google.com/go/auth"
	"gorm.io/gorm"
	// "firebase.google.com/go/auth"
)

// interactor : the struct that will have as properties all the IMPLEMENTED interfaces
// in order to provide them to its methods : the use cases and implement the Handler interface
// type Repository struct {
// 	Logger          LoggerRepository
// 	Presenter       PresenterRepository
// 	BlogDao         BlogRepository
// 	UserDao         UserRepository
// 	FirebaseHandler FirebaseHandlerRepository
// 	// validator     Validator
// 	DBTransaction DBTransactionRepository
// }
//
// func (rp Repository) New(rpInstance Repository) Repository {
// 	return Repository{
// 		Logger:          rpInstance.Logger,
// 		Presenter:       rpInstance.Presenter,
// 		BlogDao:         rpInstance.BlogDao,
// 		UserDao:         rpInstance.UserDao,
// 		FirebaseHandler: rpInstance.FirebaseHandler,
// 		DBTransaction:   rpInstance.DBTransaction,
// 	}
// }

// Logger : only used to log stuff
type LoggerRepository interface {
	Log(...interface{})
}

type PresenterRepository interface {
	Raise(errorKind models.ErrorKinds, err error)
	Present() error

	CreateSignUp(user *models.User)
	GetBlog(blog *models.Blog)
	CreateBlog(blog *models.Blog)
	GetBlogs(blogs *models.Blogs)

	GetUsers(users *models.Users)
	GetUser(user *models.User)
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
	CreateUser(user models.User) (uuId *string, err error)
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

type BlogRepository interface {
	GetAll() (*models.Blogs, error)
	GetById(id int) (*models.Blog, error)
	Create(blog models.Blog) (*models.Blog, error)
	// CreateTx(company models.Blog, tx *sql.Tx) (*models.Blog, error)
	Update(id int, company models.Blog) (*models.Blog, error)
	Delete(id int) error
}

type UserRepository interface {
	GetAll() (*models.Users, error)
	GetById(id int) (*models.User, error)
	Create(user models.User) (*models.User, error)
	CreateTx(user models.User, tx *gorm.DB) (*models.User, error)
	Update(id int, user models.User) (*models.User, error)
	Delete(id int) error
}
