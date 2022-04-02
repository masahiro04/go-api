package uc

import (
	"go-api/domains"
)

type GetUserUseCase struct {
	OutputPort PresenterRepository
	InputPort  GetUserParams
}

type GetUserParams struct {
	ID int
}

func (rp Repository) UserGet(uc GetUserUseCase) {
	user, err := rp.userDao.GetById(uc.InputPort.ID)
	if err != nil {
		uc.OutputPort.Raise(domains.BadRequest, err)
		return
	}

	if user == nil {
		uc.OutputPort.Raise(domains.NotFound, errNotFound)
		return
	}

	uc.OutputPort.GetUser(user)
}
