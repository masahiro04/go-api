package uc

import (
	"go-api/domains"
)

type GetBlogsUseCase struct {
	OutputPort PresenterRepository
	InputPort  GetBlogsParams
}

type GetBlogsParams struct {
	Limit  int
	Offset int
}

func (rp Repository) BlogGetAll(uc GetBlogsUseCase) {
	blogs, err := rp.blogDao.GetAll()
	if err != nil {
		uc.OutputPort.Raise(domains.BadRequest, err)
		return
	}

	blogs.ApplyLimitAndOffset(uc.InputPort.Limit, uc.InputPort.Offset)

	uc.OutputPort.GetBlogs(blogs)
}
