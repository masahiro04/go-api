package usecases

import (
	"context"
	"errors"
	"go-api/domains"
	"go-api/domains/models"
	"go-api/domains/models/blog"
)

// NOTE(okubo): InputPort
type EditBlogParams struct {
	ID    int
	Title string
	Body  string
}

// NOTE(okubo): OutputPort
type editBlogUseCase struct {
	Ctx        context.Context
	Logger     domains.Logger
	OutputPort domains.PresenterRepository
	BlogDao    domains.BlogRepository
}

func NewEditBlogUseCase(
	ctx context.Context, logger domains.Logger,
	outputPort domains.PresenterRepository, blogDao domains.BlogRepository,
) *editBlogUseCase {
	return &editBlogUseCase{
		Ctx:        ctx,
		Logger:     logger,
		OutputPort: outputPort,
		BlogDao:    blogDao,
	}
}

// NOTE(okubo): OutputPort(出力) と InputPort(入力) を結びつける = interactor
func (uc editBlogUseCase) BlogEdit(params EditBlogParams) {
	var newBlog *models.Blog
	var err error

	newBlog, err = uc.BlogDao.GetById(params.ID)
	if err != nil {
		uc.Logger.Errorf(uc.Ctx, err.Error())
		uc.OutputPort.Raise(models.BadRequest, err)
		return
	}

	if newBlog == nil {
		uc.OutputPort.Raise(models.NotFound, errors.New("note found"))
		return
	}

	// NOTE(okubo): input portで検索している -> どう考えてもerrは起きない
	id, _ := blog.NewId(params.ID)

	title, err := blog.UpdateTitle(&params.Title)
	if err != nil {
		uc.Logger.Errorf(uc.Ctx, err.Error())
		uc.OutputPort.Raise(models.UnprocessableEntity, err)
		return
	}

	body, err := blog.UpdateBody(&params.Body)
	if err != nil {
		uc.Logger.Errorf(uc.Ctx, err.Error())
		uc.OutputPort.Raise(models.UnprocessableEntity, err)
		return
	}

	updatedBlog, err := uc.BlogDao.Update(
		params.ID, models.BuildBlog(id, *title, *body, newBlog.CreatedAt, newBlog.UpdatedAt),
	)
	if err != nil {
		uc.Logger.Errorf(uc.Ctx, err.Error())
		uc.OutputPort.Raise(models.UnprocessableEntity, err)
		return
	}

	uc.OutputPort.GetBlog(updatedBlog)
}
