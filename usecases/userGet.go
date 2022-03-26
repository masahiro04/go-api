package uc

import (
	"go-api/domains"
)

type GetUserUseCase struct {
	OutputPort Presenter
	InputPort  GetUserParams
}

type GetUserParams struct {
	ID int
}

func (i interactor) UserGet(uc GetUserUseCase) {
	user, err := i.userDao.GetById(uc.InputPort.ID)
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
