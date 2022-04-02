package usecases

import (
	"go-api/domains"
	"go-api/domains/models"
)

type GetBlogUseCase struct {
	OutputPort domains.PresenterRepository
	BlogDao    domains.BlogRepository
}

type GetBlogParams struct {
	Id int
}

func (uc GetBlogUseCase) BlogGet(params GetBlogParams) {
	blog, err := uc.BlogDao.GetById(params.Id)
	if err != nil {
		uc.OutputPort.Raise(models.BadRequest, err)
		return
	}

	if blog == nil {
		uc.OutputPort.Raise(models.NotFound, errNotFound)
		return
	}

	uc.OutputPort.GetBlog(blog)
}
