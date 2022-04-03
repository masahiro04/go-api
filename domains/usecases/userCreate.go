package usecases

import (
	"context"
	"go-api/domains"
	"go-api/domains/models"
	"go-api/domains/models/user"
)

// NOTE(okubo): InputPort
type CreateUserParams struct {
	Name     string
	Email    string
	Password string
}

// NOTE(okubo): OutputPort
type createUserUseCase struct {
	Ctx        context.Context
	Logger     domains.Logger
	OutputPort domains.PresenterRepository
	UserDao    domains.UserRepository
}

func NewCreateUserUseCase(
	ctx context.Context,
	logger domains.Logger,
	outputPort domains.PresenterRepository,
	userDao domains.UserRepository,
) *createUserUseCase {
	return &createUserUseCase{
		Ctx:        ctx,
		Logger:     logger,
		OutputPort: outputPort,
		UserDao:    userDao,
	}
}

// NOTE(okubo): OutputPort(出力) と InputPort(入力) を結びつける = interactor
func (uc createUserUseCase) UserCreate(params CreateUserParams) {
	name, err := user.NewName(params.Name)
	if err != nil {
		uc.Logger.Errorf(uc.Ctx, err.Error())
		uc.OutputPort.Raise(models.UnprocessableEntity, err)
		return
	}
	uuid, err := user.NewUUID("dummy")
	if err != nil {
		uc.Logger.Errorf(uc.Ctx, err.Error())
		uc.OutputPort.Raise(models.UnprocessableEntity, err)
		return
	}

	email, err := user.NewEmail(params.Email)
	if err != nil {
		uc.Logger.Errorf(uc.Ctx, err.Error())
		uc.OutputPort.Raise(models.UnprocessableEntity, err)
		return
	}

	password, err := user.NewPassword(params.Password)
	if err != nil {
		uc.Logger.Errorf(uc.Ctx, err.Error())
		uc.OutputPort.Raise(models.UnprocessableEntity, err)
		return
	}

	newUser := models.NewUser(uuid, name, email, password)

	createdUser, err := uc.UserDao.Create(newUser)
	if err != nil {
		uc.Logger.Errorf(uc.Ctx, err.Error())
		uc.OutputPort.Raise(models.UnprocessableEntity, err)
		return
	}

	uc.OutputPort.GetUser(createdUser)
}
