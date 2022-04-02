package usecases

import (
	"go-api/domains"
	"go-api/domains/models"
)

// NOTE(okubo): OutputPort
type DeleteUserUseCase struct {
	OutputPort domains.PresenterRepository
	UserDao    domains.UserRepository
}

// NOTE(okubo): InputPort
type DeleteUserParams struct {
	ID int
}

// NOTE(okubo): OutputPort(出力) と InputPort(入力) を結びつける = interactor
func (uc DeleteUserUseCase) UserDelete(params DeleteUserParams) {
	if err := uc.UserDao.Delete(params.ID); err != nil {
		uc.OutputPort.Raise(models.BadRequest, err)
		return
	}
}
