package uc

import (
	"go-api/domains"
)

type GetUsersUseCase struct {
	OutputPort PresenterRepository
	InputPort  GetUsersParams
}

type GetUsersParams struct {
	Limit  int
	Offset int
}

func (rp Repository) UserGetAll(uc GetUsersUseCase) {
	users, err := rp.userDao.GetAll()
	if err != nil {
		uc.OutputPort.Raise(domains.BadRequest, err)
		return
	}

	users.ApplyLimitAndOffset(uc.InputPort.Limit, uc.InputPort.Offset)

	uc.OutputPort.GetUsers(users)

}
