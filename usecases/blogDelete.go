package uc

import (
	"go-api/domains"
)

type DeleteBlogUseCase struct {
	OutputPort PresenterRepository
	InputPort  DeleteBlogParams
}

type DeleteBlogParams struct {
	Id int
}

func (rp Repository) BlogDelete(uc DeleteBlogUseCase) {
	if err := rp.blogDao.Delete(uc.InputPort.Id); err != nil {
		uc.OutputPort.Raise(domains.BadRequest, err)
		return
	}
}
