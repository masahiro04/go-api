package uc

import (
	"go-api/domains"
	userModel "go-api/domains/user"
)

type EditUserUseCase struct {
	OutputPort PresenterRepository
	InputPort  EditUserParams
}

type EditUserParams struct {
	ID    int
	Name  string
	Email string
}

func (rp Repository) UserEdit(uc EditUserUseCase) {
	var user *domains.User
	var err error

	user, err = rp.userDao.GetById(uc.InputPort.ID)
	if err != nil {
		uc.OutputPort.Raise(domains.BadRequest, err)
		return
	}

	if user == nil {
		uc.OutputPort.Raise(domains.NotFound, errNotFound)
		return
	}

	// NOTE(okubo): input portで検索している -> どう考えてもerrは起きない
	id, _ := userModel.NewId(uc.InputPort.ID)
	uuid, err := userModel.NewUUID(user.UUID.Value)
	if err != nil {
		uc.OutputPort.Raise(domains.UnprocessableEntity, err)
		return
	}

	name, err := userModel.UpdateName(&uc.InputPort.Name)
	if err != nil {
		uc.OutputPort.Raise(domains.UnprocessableEntity, err)
		return
	}

	email, err := userModel.UpdateEmail(&uc.InputPort.Email)
	if err != nil {
		uc.OutputPort.Raise(domains.UnprocessableEntity, err)
		return
	}

	updatedUser, err := rp.userDao.Update(
		uc.InputPort.ID, domains.BuildUser(id, uuid, *name, *email, user.CreatedAt, user.UpdatedAt),
	)
	if err != nil {
		uc.OutputPort.Raise(domains.UnprocessableEntity, err)
		return
	}

	uc.OutputPort.GetUser(updatedUser)
}
