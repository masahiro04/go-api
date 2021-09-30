package uc

import (
	"clean_architecture/golang/domain"
)

type GetBlogUseCase struct {
	OutputPort Presenter
	InputPort  GetBlogParams
}

type GetBlogParams struct {
	Id int
}

func (i interactor) BlogGet(uc GetBlogUseCase) {
	blog, err := i.blogRW.GetById(uc.InputPort.Id)
	if err != nil {
		uc.OutputPort.Raise(domain.BadRequest, err)
		return
	}

	if blog == nil {
		uc.OutputPort.Raise(domain.NotFound, errNotFound)
		return
	}

	uc.OutputPort.GetBlog(blog)
}
