package uc

import (
	"clean_architecture/golang/domain"
)

type CreateBlogUseCase struct {
	OutputPort Presenter
	InputPort  CreateBlogParams
}

type CreateBlogParams struct {
	Title string
	Body  string
}

func (i interactor) BlogCreate(uc CreateBlogUseCase) {
	var blog *domain.Blog
	var err error

	err = i.validator.Validate(blog)
	if err != nil {
		uc.OutputPort.Raise(domain.UnprocessableEntity, err)
		return
	}

	blog = &domain.Blog{
		Title: uc.InputPort.Title,
		Body:  uc.InputPort.Body,
	}

	blog, err = i.blogRW.Create(*blog)
	if err != nil {
		uc.OutputPort.Raise(domain.UnprocessableEntity, err)
		return
	}

	uc.OutputPort.CreateBlog(blog)
}
