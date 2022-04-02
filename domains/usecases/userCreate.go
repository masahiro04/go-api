package usecases

import (
	"go-api/domains"
	"go-api/domains/user"
)

type CreateUserUseCase struct {
	OutputPort PresenterRepository
	InputPort  CreateUserParams
}

type CreateUserParams struct {
	Name     string
	Email    string
	Password string
}

func (rp Repository) UserCreate(uc CreateUserUseCase) {
	name, err := user.NewName(uc.InputPort.Name)
	if err != nil {
		uc.OutputPort.Raise(domains.UnprocessableEntity, err)
		return
	}
	uuid, err := user.NewUUID("dummy")
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

	newUser := domains.NewUser(uuid, name, email, password)

	createdUser, err := rp.userDao.Create(newUser)
	if err != nil {
		uc.OutputPort.Raise(domains.UnprocessableEntity, err)
		return
	}

	uc.OutputPort.GetUser(createdUser)
}
