package uc

import (
	"go-api/domains"
	"go-api/domains/blog"
)

type CreateBlogUseCase struct {
	OutputPort PresenterRepository
	InputPort  CreateBlogParams
}

type CreateBlogParams struct {
	Title string
	Body  string
}

func (rp Repository) BlogCreate(uc CreateBlogUseCase) {
	title, err := blog.NewTitle(uc.InputPort.Title)
	if err != nil {
		uc.OutputPort.Raise(domains.UnprocessableEntity, err)
		return
	}

	body, err := blog.NewBody(uc.InputPort.Body)
	if err != nil {
		uc.OutputPort.Raise(domains.UnprocessableEntity, err)
		return
	}

	newBlog := domains.NewBlog(title, body)

	createdBlog, err := rp.blogDao.Create(newBlog)
	if err != nil {
		uc.OutputPort.Raise(domains.UnprocessableEntity, err)
		return
	}

	uc.OutputPort.CreateBlog(createdBlog)
}
