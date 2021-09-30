package uc

import "clean_architecture/golang/domain"

type DeleteBlogUseCase struct {
	OutputPort Presenter
	InputPort  DeleteBlogParams
}

type DeleteBlogParams struct {
	Id int
}

func (i interactor) BlogDelete(uc DeleteBlogUseCase) {
	if err := i.blogRW.Delete(uc.InputPort.Id); err != nil {
		uc.OutputPort.Raise(domain.BadRequest, err)
		return
	}
}
