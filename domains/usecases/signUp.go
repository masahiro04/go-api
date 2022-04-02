package usecases

import (
	"go-api/domains"
	"go-api/domains/models"
	"go-api/domains/models/user"

	"gorm.io/gorm"
)

type SignUpUseCase struct {
	OutputPort      domains.PresenterRepository
	UserDao         domains.UserRepository
	DBTransaction   domains.DBTransactionRepository
	FirebaseHandler domains.FirebaseHandlerRepository
}

type SignUpParams struct {
	Name     string
	Email    string
	Password string
}

func (uc SignUpUseCase) SignUp(params SignUpParams) {
	// var err error
	var createdUser models.User

	name, err := user.NewName(params.Name)
	if err != nil {
		uc.OutputPort.Raise(models.UnprocessableEntity, err)
		return
	}

	email, err := user.NewEmail(params.Email)
	if err != nil {
		uc.OutputPort.Raise(models.UnprocessableEntity, err)
		return
	}

	password, err := user.NewPassword(params.Password)
	if err != nil {
		uc.OutputPort.Raise(models.UnprocessableEntity, err)
		return
	}

	dummyUUID, err := user.NewUUID("dummy")
	u := models.NewUser(dummyUUID, name, email, password)
	// uuId, err := i.firebaseHandler.CreateUser(user)

	err = uc.DBTransaction.WithTx(func(tx *gorm.DB) error {
		uuid, err := uc.FirebaseHandler.CreateUser(u)
		if err != nil {
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
		uc.OutputPort.Raise(models.UnprocessableEntity, err)
		return
	}
	uc.OutputPort.CreateSignUp(&createdUser)
}
