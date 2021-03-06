package usecases

import (
	"context"
	"go-api/domains/models"
)

// NOTE(okubo): InputPort
type GetBlogParams struct {
	ID int
}

// NOTE(okubo): OutputPort
type getBlogUseCase struct {
	Ctx        context.Context
	Logger     Logger
	OutputPort PresenterRepository
	BlogDao    BlogRepository
}

func NewGetBlogUseCase(
	ctx context.Context,
	logger Logger,
	outputPort PresenterRepository,
	blogDao BlogRepository,
) IBlogGet {
	return &getBlogUseCase{
		Ctx:        ctx,
		Logger:     logger,
		OutputPort: outputPort,
		BlogDao:    blogDao,
	}
}

// NOTE(okubo): OutputPort(出力) と InputPort(入力) を結びつける = interactor
func (uc getBlogUseCase) BlogGet(params GetBlogParams) {
	blog, err := uc.BlogDao.GetById(params.ID)
	if err != nil {
		uc.Logger.Errorf(uc.Ctx, err.Error())
		uc.OutputPort.Raise(models.BadRequest, err)
		return
	}

	if blog == nil {
		uc.Logger.Warnf(uc.Ctx, err.Error())
		uc.OutputPort.Raise(models.NotFound, errNotFound)
		return
	}

	uc.OutputPort.GetBlog(blog)
}
