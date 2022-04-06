package usecases

import (
	"context"
	"go-api/domains/models"
)

// NOTE(okubo): InputPort
type GetUserParams struct {
	ID int
}

// NOTE(okubo): OutputPort
type getUserUseCase struct {
	Ctx        context.Context
	Logger     Logger
	OutputPort PresenterRepository
	UserDao    UserRepository
}

func NewGetUserUseCase(
	ctx context.Context,
	logger Logger,
	outputPort PresenterRepository,
	userDao UserRepository,
) IUserGet {
	return &getUserUseCase{
		Ctx:        ctx,
		Logger:     logger,
		OutputPort: outputPort,
		UserDao:    userDao,
	}
}

// NOTE(okubo): OutputPort(出力) と InputPort(入力) を結びつける = interactor
func (uc getUserUseCase) UserGet(params GetUserParams) {
	user, err := uc.UserDao.GetById(params.ID)
	if err != nil {
		uc.Logger.Errorf(uc.Ctx, err.Error())
		uc.OutputPort.Raise(models.BadRequest, err)
		return
	}

	if user == nil {
		uc.Logger.Errorf(uc.Ctx, err.Error())
		uc.OutputPort.Raise(models.NotFound, errNotFound)
		return
	}

	uc.OutputPort.GetUser(user)
}
