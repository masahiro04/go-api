package usecases

import (
	"context"
	"go-api/domains/models"
	"go-api/domains/models/blog"
)

// NOTE(okubo): InputPort
type CreateBlogParams struct {
	Title string
	Body  string
}

// NOTE(okubo): OutputPort
type createBlogUseCase struct {
	Ctx        context.Context
	Logger     Logger
	OutputPort PresenterRepository
	BlogDao    BlogRepository
}

func NewCreateBlogUseCase(ctx context.Context, logger Logger, outputPort PresenterRepository, blogDao BlogRepository) *createBlogUseCase {
	return &createBlogUseCase{
		Ctx:        ctx,
		Logger:     logger,
		OutputPort: outputPort,
		BlogDao:    blogDao,
	}
}

func (uc createBlogUseCase) BlogCreate(params CreateBlogParams) {
	title, err := blog.NewTitle(params.Title)
	if err != nil {
		uc.OutputPort.Raise(models.UnprocessableEntity, err)
		return
	}

	body, err := blog.NewBody(params.Body)
	if err != nil {
		uc.Logger.Warnf(uc.Ctx, err.Error())
		uc.OutputPort.Raise(models.UnprocessableEntity, err)
		return
	}

	newBlog := models.NewBlog(title, body)

	createdBlog, err := uc.BlogDao.Create(newBlog)
	if err != nil {
		uc.Logger.Warnf(uc.Ctx, err.Error())
		uc.OutputPort.Raise(models.UnprocessableEntity, err)
		return
	}

	uc.OutputPort.CreateBlog(createdBlog)
}
