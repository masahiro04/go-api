package usecases

import (
	"context"
	"go-api/domains"
	"go-api/domains/models"
)

// NOTE(okubo): InputPort
type GetUsersParams struct {
	Limit  int
	Offset int
}

// NOTE(okubo): OutputPort
type getUsersUseCase struct {
	Ctx        context.Context
	Logger     domains.Logger
	OutputPort domains.PresenterRepository
	UserDao    domains.UserRepository
}

func NewGetUsersUseCase(
	ctx context.Context,
	logger domains.Logger,
	outputPort domains.PresenterRepository,
	userDao domains.UserRepository,
) *getUsersUseCase {
	return &getUsersUseCase{
		Ctx:        ctx,
		Logger:     logger,
		OutputPort: outputPort,
		UserDao:    userDao,
	}
}

// NOTE(okubo): OutputPort(出力) と InputPort(入力) を結びつける = interactor
func (uc getUsersUseCase) UserGetAll(params GetUsersParams) {
	users, err := uc.UserDao.GetAll()
	if err != nil {
		uc.Logger.Errorf(uc.Ctx, err.Error())
		uc.OutputPort.Raise(models.BadRequest, err)
		return
	}

	users.ApplyLimitAndOffset(params.Limit, params.Offset)

	uc.OutputPort.GetUsers(users)

}
