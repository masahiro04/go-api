package usecases

import (
	"context"
	"go-api/domains/models"
	"go-api/domains/models/user"
)

// NOTE(okubo): InputPort
type EditUserParams struct {
	ID    int
	Name  string
	Email string
}

// NOTE(okubo): OutputPort
type editUserUseCase struct {
	Ctx        context.Context
	Logger     Logger
	OutputPort PresenterRepository
	UserDao    UserRepository
}

func NewEditUserUseCase(
	ctx context.Context,
	logger Logger,
	outputPort PresenterRepository,
	userDao UserRepository,
) IUserEdit {
	return &editUserUseCase{
		Ctx:        ctx,
		Logger:     logger,
		OutputPort: outputPort,
		UserDao:    userDao,
	}
}

// NOTE(okubo): OutputPort(出力) と InputPort(入力) を結びつける = interactor
func (uc editUserUseCase) UserEdit(params EditUserParams) {
	newUser, err := uc.UserDao.GetById(params.ID)
	if err != nil {
		uc.Logger.Errorf(uc.Ctx, err.Error())
		uc.OutputPort.Raise(models.BadRequest, err)
		return
	}

	if newUser == nil {
		uc.Logger.Errorf(uc.Ctx, err.Error())
		uc.OutputPort.Raise(models.NotFound, errNotFound)
		return
	}

	// NOTE(okubo): input portで検索している -> どう考えてもerrは起きない
	id, _ := user.NewId(params.ID)
	uuid, err := user.NewUUID(newUser.UUID.Value)
	if err != nil {
		uc.Logger.Errorf(uc.Ctx, err.Error())
		uc.OutputPort.Raise(models.UnprocessableEntity, err)
		return
	}

	name, err := user.UpdateName(&params.Name)
	if err != nil {
		uc.Logger.Errorf(uc.Ctx, err.Error())
		uc.OutputPort.Raise(models.UnprocessableEntity, err)
		return
	}

	email, err := user.UpdateEmail(&params.Email)
	if err != nil {
		uc.Logger.Errorf(uc.Ctx, err.Error())
		uc.OutputPort.Raise(models.UnprocessableEntity, err)
		return
	}

	updatedUser, err := uc.UserDao.Update(
		params.ID, models.BuildUser(id, uuid, *name, *email, newUser.CreatedAt, newUser.UpdatedAt),
	)
	if err != nil {
		uc.Logger.Errorf(uc.Ctx, err.Error())
		uc.OutputPort.Raise(models.UnprocessableEntity, err)
		return
	}

	uc.OutputPort.GetUser(updatedUser)
}
