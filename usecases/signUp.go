package uc

import (
	"go-api/domains"

	"go-api/domains/user"

	"gorm.io/gorm"
)

type SignUpUseCase struct {
	OutputPort Presenter
	InputPort  SignUpParams
}

type SignUpParams struct {
	Name     string
	Email    string
	Password string
}

func (i interactor) SignUp(uc SignUpUseCase) {
	// var err error
	var createdUser domains.User

	name, err := user.NewName(uc.InputPort.Name)
	if err != nil {
		uc.OutputPort.Raise(domains.UnprocessableEntity, err)
		return
	}

	email, err := user.NewEmail(uc.InputPort.Email)
	if err != nil {
		uc.OutputPort.Raise(domains.UnprocessableEntity, err)
		return
	}

	password, err := user.NewPassword(uc.InputPort.Password)
	if err != nil {
		uc.OutputPort.Raise(domains.UnprocessableEntity, err)
		return
	}

	dummyUUID, err := user.NewUUID("dummy")
	u := domains.NewUser(dummyUUID, name, email, password)
	// uuId, err := i.firebaseHandler.CreateUser(user)

	err = i.dbTransaction.WithTx(func(tx *gorm.DB) error {
		uuid, err := i.firebaseHandler.CreateUser(u)
		if err != nil {
			i.logger.Log(err)
			return err
		}

		newUUID, err := user.NewUUID(*uuid)
		u.UUID = newUUID

		usr, err := i.userDao.CreateTx(u, tx)
		createdUser = *usr
		if err != nil {
			if rollbackErr := tx.Rollback().Error; rollbackErr != nil {
				if err = i.firebaseHandler.DeleteUser(createdUser.UUID.Value); err != nil {
					i.logger.Log(err)
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
		uc.OutputPort.Raise(domains.UnprocessableEntity, err)
		return
	}
	uc.OutputPort.CreateSignUp(&createdUser)
}
