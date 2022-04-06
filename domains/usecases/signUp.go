package usecases

import (
	"context"
	"go-api/domains"
	"go-api/domains/models"
	"go-api/domains/models/user"

	"gorm.io/gorm"
)

// NOTE(okubo): InputPort
type SignUpParams struct {
	Name     string
	Email    string
	Password string
}

// NOTE(okubo): OutputPort
type signUpUseCase struct {
	Ctx             context.Context
	Logger          domains.Logger
	OutputPort      domains.PresenterRepository
	UserDao         domains.UserRepository
	DBTransaction   domains.DBTransactionRepository
	FirebaseHandler domains.FirebaseHandlerRepository
}

func NewSignUpUseCase(
	ctx context.Context,
	logger domains.Logger,
	outputPort domains.PresenterRepository,
	userDao domains.UserRepository,
	dbtransaction domains.DBTransactionRepository,
	firebaseHandler domains.FirebaseHandlerRepository,
) *signUpUseCase {
	return &signUpUseCase{
		Ctx:             ctx,
		Logger:          logger,
		OutputPort:      outputPort,
		UserDao:         userDao,
		DBTransaction:   dbtransaction,
		FirebaseHandler: firebaseHandler,
	}
}

// NOTE(okubo): InputPort
// NOTE(okubo): OutputPort(出力) と InputPort(入力) を結びつける = interactor
func (uc signUpUseCase) SignUp(params SignUpParams) {
	var createdUser models.User

	name, err := user.NewName(params.Name)
	if err != nil {
		uc.Logger.Errorf(uc.Ctx, err.Error())
		uc.OutputPort.Raise(models.UnprocessableEntity, err)
		return
	}

	email, err := user.NewEmail(params.Email)
	if err != nil {
		uc.Logger.Errorf(uc.Ctx, err.Error())
		uc.OutputPort.Raise(models.UnprocessableEntity, err)
		return
	}

	password, err := user.NewPassword(params.Password)
	if err != nil {
		uc.Logger.Errorf(uc.Ctx, err.Error())
		uc.OutputPort.Raise(models.UnprocessableEntity, err)
		return
	}

	dummyUUID, err := user.NewUUID("dummy")
	u := models.NewUser(dummyUUID, name, email, password)
	// uuId, err := i.firebaseHandler.CreateUser(user)

	err = uc.DBTransaction.WithTx(func(tx *gorm.DB) error {
		uuid, err := uc.FirebaseHandler.CreateUser(u)
		if err != nil {
			uc.Logger.Errorf(uc.Ctx, err.Error())
			// rp.logger.Log(err)
			return err
		}

		newUUID, err := user.NewUUID(*uuid)
		u.UUID = newUUID

		usr, err := uc.UserDao.CreateTx(u, tx)
		createdUser = *usr
		if err != nil {
			if rollbackErr := tx.Rollback().Error; rollbackErr != nil {
				if err = uc.FirebaseHandler.DeleteUser(createdUser.UUID.Value); err != nil {
					uc.Logger.Errorf(uc.Ctx, err.Error())
					// rp.logger.Log(err)
				}
			}
			return err
		}

		return nil
	})

	// link, err := i.firebaseHandler.EmailSignInLink(user.Email)
	// if err != nil {
	// 	presenter.Raise(domain.UnprocessableEntity, err)
	// 	return
	// }
	//
	// message, err := i.mailPresenter.SendMessage(
	// 	user.Email,
	// 	domain.MessageTitleForSignUp,
	// 	domain.MessageBodyForSignUp,
	// 	&domain.MessageTemplateData{Link: link})
	// if err != nil {
	// 	presenter.Raise(domain.UnprocessableEntity, err)
	// 	return
	// }
	//
	// err = i.awsSqs.SendMessage(*message)
	// if err != nil {
	// 	presenter.Raise(domain.UnprocessableEntity, err)
	// 	return
	// }
	if err != nil {
		uc.Logger.Errorf(uc.Ctx, err.Error())
		uc.OutputPort.Raise(models.UnprocessableEntity, err)
		return
	}
	uc.OutputPort.CreateSignUp(&createdUser)
}
