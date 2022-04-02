package usecases

import (
	"go-api/domains"
	"go-api/domains/models"
	"go-api/domains/models/user"
)

// NOTE(okubo): OutputPort
type CreateUserUseCase struct {
	OutputPort domains.PresenterRepository
	UserDao    domains.UserRepository
}

// NOTE(okubo): InputPort
type CreateUserParams struct {
	Name     string
	Email    string
	Password string
}

// NOTE(okubo): OutputPort(出力) と InputPort(入力) を結びつける = interactor
func (uc CreateUserUseCase) UserCreate(params CreateUserParams) {
	name, err := user.NewName(params.Name)
	if err != nil {
		uc.OutputPort.Raise(models.UnprocessableEntity, err)
		return
	}
	uuid, err := user.NewUUID("dummy")
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

	newUser := models.NewUser(uuid, name, email, password)

	createdUser, err := uc.UserDao.Create(newUser)
	if err != nil {
		uc.OutputPort.Raise(models.UnprocessableEntity, err)
		return
	}

	uc.OutputPort.GetUser(createdUser)
}
