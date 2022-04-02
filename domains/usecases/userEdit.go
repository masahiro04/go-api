package usecases

import (
	"go-api/domains"
	"go-api/domains/models"
	"go-api/domains/models/user"
)

type EditUserUseCase struct {
	OutputPort domains.PresenterRepository
	UserDao    domains.UserRepository
}

type EditUserParams struct {
	ID    int
	Name  string
	Email string
}

func (uc EditUserUseCase) UserEdit(params EditUserParams) {
	newUser, err := uc.UserDao.GetById(params.ID)
	if err != nil {
		uc.OutputPort.Raise(models.BadRequest, err)
		return
	}

	if newUser == nil {
		uc.OutputPort.Raise(models.NotFound, errNotFound)
		return
	}

	// NOTE(okubo): input portで検索している -> どう考えてもerrは起きない
	id, _ := user.NewId(params.ID)
	uuid, err := user.NewUUID(newUser.UUID.Value)
	if err != nil {
		uc.OutputPort.Raise(models.UnprocessableEntity, err)
		return
	}

	name, err := user.UpdateName(&params.Name)
	if err != nil {
		uc.OutputPort.Raise(models.UnprocessableEntity, err)
		return
	}

	email, err := user.UpdateEmail(&params.Email)
	if err != nil {
		uc.OutputPort.Raise(models.UnprocessableEntity, err)
		return
	}

	updatedUser, err := uc.UserDao.Update(
		params.ID, models.BuildUser(id, uuid, *name, *email, newUser.CreatedAt, newUser.UpdatedAt),
	)
	if err != nil {
		uc.OutputPort.Raise(models.UnprocessableEntity, err)
		return
	}

	uc.OutputPort.GetUser(updatedUser)
}
