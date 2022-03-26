package uc

import (
	"go-api/domains"
)

type GetBlogsUseCase struct {
	OutputPort Presenter
	InputPort  GetBlogsParams
}

type GetBlogsParams struct {
	Limit  int
	Offset int
}

func (i interactor) BlogGetAll(uc GetBlogsUseCase) {
	blogs, err := i.blogDao.GetAll()
	if err != nil {
		uc.OutputPort.Raise(domains.BadRequest, err)
		return
	}

	blogs.ApplyLimitAndOffset(uc.InputPort.Limit, uc.InputPort.Offset)

	uc.OutputPort.GetBlogs(blogs)
}
