package usecases

import (
	"errors"
	"go-api/domains"
	"go-api/domains/models"
	"go-api/domains/models/blog"
)

// NOTE(okubo): OutputPort
type EditBlogUseCase struct {
	OutputPort domains.PresenterRepository
	BlogDao    domains.BlogRepository
}

// NOTE(okubo): InputPort
type EditBlogParams struct {
	Id    int
	Title string
	Body  string
}

// NOTE(okubo): OutputPort(出力) と InputPort(入力) を結びつける = interactor
func (uc EditBlogUseCase) BlogEdit(params EditBlogParams) {
	var newBlog *models.Blog
	var err error

	newBlog, err = uc.BlogDao.GetById(params.Id)
	if err != nil {
		uc.OutputPort.Raise(models.BadRequest, err)
		return
	}

	if newBlog == nil {
		uc.OutputPort.Raise(models.NotFound, errors.New("note found"))
		return
	}

	// NOTE(okubo): input portで検索している -> どう考えてもerrは起きない
	id, _ := blog.NewId(params.Id)

	title, err := blog.UpdateTitle(&params.Title)
	if err != nil {
		uc.OutputPort.Raise(models.UnprocessableEntity, err)
		return
	}

	body, err := blog.UpdateBody(&params.Body)
	if err != nil {
		uc.OutputPort.Raise(models.UnprocessableEntity, err)
		return
	}

	updatedBlog, err := uc.BlogDao.Update(
		params.Id, models.BuildBlog(id, *title, *body, newBlog.CreatedAt, newBlog.UpdatedAt),
	)
	if err != nil {
		uc.OutputPort.Raise(models.UnprocessableEntity, err)
		return
	}

	uc.OutputPort.GetBlog(updatedBlog)
}
