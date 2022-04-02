package usecases

import (
	"go-api/domains"
	"go-api/domains/models"
)

// NOTE(okubo): OutputPort
type GetUsersUseCase struct {
	OutputPort domains.PresenterRepository
	UserDao    domains.UserRepository
}

// NOTE(okubo): InputPort
type GetUsersParams struct {
	Limit  int
	Offset int
}

// NOTE(okubo): OutputPort(出力) と InputPort(入力) を結びつける = interactor
func (uc GetUsersUseCase) UserGetAll(params GetUsersParams) {
	users, err := uc.UserDao.GetAll()
	if err != nil {
		uc.OutputPort.Raise(models.BadRequest, err)
		return
	}

	users.ApplyLimitAndOffset(params.Limit, params.Offset)

	uc.OutputPort.GetUsers(users)

}
