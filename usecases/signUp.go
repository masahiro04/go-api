package uc

import (
	"go-api/domains"

	"go-api/domains/user"
	// "gorm.io/gorm"
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

// type SignUpParams struct {
// 	Name     string
// 	Email    string
// 	Password string
// 	// User SignUpUserParams
// }

// type SignUpUserParams struct {
// 	Email    string
// 	Password string
// }
//
// type SignUpCompanyParams struct {
// 	Name string
// }
//
func (i interactor) SignUp(uc SignUpUseCase) {
	var err error

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

	user := domains.NewUser(name, email, password)
	// uuId, err := i.firebaseHandler.CreateUser(user)
	_, err = i.firebaseHandler.CreateUser(user)
	if err != nil {
		uc.OutputPort.Raise(domains.UnprocessableEntity, err)
		return
	}
	// TODO(okubo): rollback入れる
	// if err != nil {
	// 	if rollbackErr := tx.Rollback(); rollbackErr != nil {
	// 		i.logger.Log(err)
	// 	}
	// 	return err
	// }

	// user.UuId = *uuId
	_, err = i.userDao.Create(user)
	if err != nil {
		uc.OutputPort.Raise(domains.UnprocessableEntity, err)
		return
	}
	// TODO(okubo): rollback入れる

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

	uc.OutputPort.CreateSignUp(&user)
}
