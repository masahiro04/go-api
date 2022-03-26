package uc

import (
	"go-api/domains"
)

type GetUsersUseCase struct {
	OutputPort Presenter
	InputPort  GetUsersParams
}

type GetUsersParams struct {
	Limit  int
	Offset int
}

func (i interactor) UserGetAll(uc GetUsersUseCase) {
	users, err := i.userDao.GetAll()
	if err != nil {
		uc.OutputPort.Raise(domains.BadRequest, err)
		return
	}

	users.ApplyLimitAndOffset(uc.InputPort.Limit, uc.InputPort.Offset)

	uc.OutputPort.GetUsers(users)

}
