package uc

import (
	"go-api/domains"
)

type DeleteBlogUseCase struct {
	OutputPort Presenter
	InputPort  DeleteBlogParams
}

type DeleteBlogParams struct {
	Id int
}

func (i interactor) BlogDelete(uc DeleteBlogUseCase) {
	if err := i.blogDao.Delete(uc.InputPort.Id); err != nil {
		uc.OutputPort.Raise(domains.BadRequest, err)
		return
	}
}
