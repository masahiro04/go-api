package uc

import (
	"clean_architecture/golang/domains"
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
	blogs, err := i.blogRW.GetAll()
	if err != nil {
		uc.OutputPort.Raise(domains.BadRequest, err)
		return
	}

	blogs = domains.BlogCollection(blogs).ApplyLimitAndOffset(uc.InputPort.Limit, uc.InputPort.Offset)

	uc.OutputPort.GetBlogs(blogs)
}
