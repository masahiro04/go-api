package usecases

import (
	"go-api/domains"
	"go-api/domains/models"
)

// NOTE(okubo): OutputPort
type GetUserUseCase struct {
	OutputPort domains.PresenterRepository
	UserDao    domains.UserRepository
}

// NOTE(okubo): InputPort
type GetUserParams struct {
	ID int
}

// NOTE(okubo): OutputPort(出力) と InputPort(入力) を結びつける = interactor
func (uc GetUserUseCase) UserGet(params GetUserParams) {
	user, err := uc.UserDao.GetById(params.ID)
	if err != nil {
		uc.OutputPort.Raise(models.BadRequest, err)
		return
	}

	if user == nil {
		uc.OutputPort.Raise(models.NotFound, errNotFound)
		return
	}

	uc.OutputPort.GetUser(user)
}
