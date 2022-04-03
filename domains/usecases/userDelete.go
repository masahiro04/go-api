package usecases

import (
	"context"
	"go-api/domains"
	"go-api/domains/models"
)

// NOTE(okubo): InputPort
type DeleteUserParams struct {
	ID int
}

// NOTE(okubo): OutputPort
type deleteUserUseCase struct {
	Ctx        context.Context
	Logger     domains.Logger
	OutputPort domains.PresenterRepository
	UserDao    domains.UserRepository
}

func NewDeleteUserUseCase(
	ctx context.Context,
	logger domains.Logger,
	outputPort domains.PresenterRepository,
	userDao domains.UserRepository,
) *deleteUserUseCase {
	return &deleteUserUseCase{
		Ctx:        ctx,
		Logger:     logger,
		OutputPort: outputPort,
		UserDao:    userDao,
	}
}

// NOTE(okubo): OutputPort(出力) と InputPort(入力) を結びつける = interactor
func (uc deleteUserUseCase) UserDelete(params DeleteUserParams) {
	if err := uc.UserDao.Delete(params.ID); err != nil {
		uc.Logger.Errorf(uc.Ctx, err.Error())
		uc.OutputPort.Raise(models.BadRequest, err)
		return
	}
}
