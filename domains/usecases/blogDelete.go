package usecases

import (
	"context"
	"go-api/domains"
	"go-api/domains/models"
)

// NOTE(okubo): InputPort
type DeleteBlogParams struct {
	ID int
}

// NOTE(okubo): OutputPort
type deleteBlogUseCase struct {
	Ctx        context.Context
	Logger     domains.Logger
	OutputPort domains.PresenterRepository
	BlogDao    domains.BlogRepository
}

func NewDeleteBlogUseCase(
	ctx context.Context, logger domains.Logger,
	outputPort domains.PresenterRepository, blogDao domains.BlogRepository,
) *deleteBlogUseCase {
	return &deleteBlogUseCase{
		Ctx:        ctx,
		Logger:     logger,
		OutputPort: outputPort,
		BlogDao:    blogDao,
	}
}

// NOTE(okubo): OutputPort(出力) と InputPort(入力) を結びつける = interactor
func (uc deleteBlogUseCase) BlogDelete(params DeleteBlogParams) {
	if err := uc.BlogDao.Delete(params.ID); err != nil {
		uc.Logger.Errorf(uc.Ctx, err.Error())
		uc.OutputPort.Raise(models.BadRequest, err)
		return
	}
}
