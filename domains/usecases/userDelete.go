package usecases

import (
	"go-api/domains"
	"go-api/domains/models"
)

type DeleteUserUseCase struct {
	OutputPort domains.PresenterRepository
	UserDao    domains.UserRepository
}

type DeleteUserParams struct {
	ID int
}

func (uc DeleteUserUseCase) UserDelete(params DeleteUserParams) {
	if err := uc.UserDao.Delete(params.ID); err != nil {
		uc.OutputPort.Raise(models.BadRequest, err)
		return
	}
}
