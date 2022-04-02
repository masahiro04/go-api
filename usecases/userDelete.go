package usecases

import (
	"go-api/domains"
)

type DeleteUserUseCase struct {
	OutputPort PresenterRepository
	InputPort  DeleteUserParams
}

type DeleteUserParams struct {
	ID int
}

func (rp Repository) UserDelete(uc DeleteUserUseCase) {
	if err := rp.userDao.Delete(uc.InputPort.ID); err != nil {
		uc.OutputPort.Raise(domains.BadRequest, err)
		return
	}
}
