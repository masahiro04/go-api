package uc

import (
	"go-api/domains"
)

type DeleteUserUseCase struct {
	OutputPort Presenter
	InputPort  DeleteUserParams
}

type DeleteUserParams struct {
	ID int
}

func (i interactor) UserDelete(uc DeleteUserUseCase) {
	if err := i.userDao.Delete(uc.InputPort.ID); err != nil {
		uc.OutputPort.Raise(domains.BadRequest, err)
		return
	}
}
