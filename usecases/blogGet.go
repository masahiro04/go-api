package uc

import (
	"go-api/domains"
)

type GetBlogUseCase struct {
	OutputPort PresenterRepository
	InputPort  GetBlogParams
}

type GetBlogParams struct {
	Id int
}

func (rp Repository) BlogGet(uc GetBlogUseCase) {
	blog, err := rp.blogDao.GetById(uc.InputPort.Id)
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
