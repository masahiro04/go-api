package uc

import (
	"clean_architecture/golang/domains"
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
		uc.OutputPort.Raise(domains.BadRequest, err)
		return
	}

	if blog == nil {
		uc.OutputPort.Raise(domains.NotFound, errNotFound)
		return
	}

	uc.OutputPort.GetBlog(blog)
}
