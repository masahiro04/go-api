package usecases

import (
	"context"
	"go-api/domains/models"
)

// NOTE(okubo): InputPort
type GetBlogsParams struct {
	Limit  int
	Offset int
}

// NOTE(okubo): OutputPort
type getBlogsUseCase struct {
	Ctx        context.Context
	Logger     Logger
	OutputPort PresenterRepository
	BlogDao    BlogRepository
}

func NewGetBlogsUseCase(ctx context.Context, logger Logger, outputPort PresenterRepository, blogDao BlogRepository) *getBlogsUseCase {
	return &getBlogsUseCase{
		Ctx:        ctx,
		Logger:     logger,
		OutputPort: outputPort,
		BlogDao:    blogDao,
	}
}

// NOTE(okubo): OutputPort(出力) と InputPort(入力) を結びつける = interactor
func (uc getBlogsUseCase) BlogGetAll(params GetBlogsParams) {
	blogs, err := uc.BlogDao.GetAll()
	if err != nil {
		uc.Logger.Errorf(uc.Ctx, err.Error())
		uc.OutputPort.Raise(models.BadRequest, err)
		return
	}
	blogs.ApplyLimitAndOffset(params.Limit, params.Offset)
	uc.OutputPort.GetBlogs(blogs)
}
