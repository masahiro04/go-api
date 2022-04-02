package usecases

import (
	"go-api/domains"
	"go-api/domains/models"
)

// NOTE(okubo): OutputPort
type GetBlogUseCase struct {
	OutputPort domains.PresenterRepository
	BlogDao    domains.BlogRepository
}

// NOTE(okubo): InputPort
type GetBlogParams struct {
	Id int
}

// NOTE(okubo): OutputPort(出力) と InputPort(入力) を結びつける = interactor
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
