package usecases

import (
	"go-api/domains"
	"go-api/domains/models"
	"go-api/domains/models/blog"
)

type CreateBlogUseCase struct {
	OutputPort domains.PresenterRepository
	BlogDao    domains.BlogRepository
}

type CreateBlogParams struct {
	Title string
	Body  string
}

func (uc CreateBlogUseCase) BlogCreate(params CreateBlogParams) {
	title, err := blog.NewTitle(params.Title)
	if err != nil {
		uc.OutputPort.Raise(models.UnprocessableEntity, err)
		return
	}

	body, err := blog.NewBody(params.Body)
	if err != nil {
		uc.OutputPort.Raise(models.UnprocessableEntity, err)
		return
	}

	newBlog := models.NewBlog(title, body)

	createdBlog, err := uc.BlogDao.Create(newBlog)
	if err != nil {
		uc.OutputPort.Raise(models.UnprocessableEntity, err)
		return
	}

	uc.OutputPort.CreateBlog(createdBlog)
}
