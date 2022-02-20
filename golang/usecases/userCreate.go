package uc

import (
	"clean_architecture/golang/domains"
	"clean_architecture/golang/domains/user"
)

type CreateUserUseCase struct {
	OutputPort Presenter
	InputPort  CreateUserParams
}

type CreateUserParams struct {
	Name     string
	Email    string
	Password string
}

func (i interactor) UserCreate(uc CreateUserUseCase) {
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

	newUser := domains.NewUser(name, email, password)

	createdUser, err := i.userDao.Create(newUser)
	if err != nil {
		uc.OutputPort.Raise(domains.UnprocessableEntity, err)
		return
	}

	uc.OutputPort.GetUser(createdUser)
}
