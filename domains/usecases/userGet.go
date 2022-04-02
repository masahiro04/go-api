package usecases

import (
	"go-api/domains"
	"go-api/domains/models"
)

type GetUserUseCase struct {
	OutputPort domains.PresenterRepository
	UserDao    domains.UserRepository
}

type GetUserParams struct {
	ID int
}

func (uc GetUserUseCase) UserGet(params GetUserParams) {
	user, err := uc.UserDao.GetById(params.ID)
	if err != nil {
		uc.OutputPort.Raise(models.BadRequest, err)
		return
	}

	if user == nil {
		uc.OutputPort.Raise(models.NotFound, errNotFound)
		return
	}

	uc.OutputPort.GetUser(user)
}
