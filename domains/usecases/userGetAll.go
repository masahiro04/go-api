package usecases

import (
	"go-api/domains"
	"go-api/domains/models"
)

type GetUsersUseCase struct {
	OutputPort domains.PresenterRepository
	UserDao    domains.UserRepository
}

type GetUsersParams struct {
	Limit  int
	Offset int
}

func (uc GetUsersUseCase) UserGetAll(params GetUsersParams) {
	users, err := uc.UserDao.GetAll()
	if err != nil {
		uc.OutputPort.Raise(models.BadRequest, err)
		return
	}

	users.ApplyLimitAndOffset(params.Limit, params.Offset)

	uc.OutputPort.GetUsers(users)

}
