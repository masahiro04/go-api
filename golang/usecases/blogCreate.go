package uc

import (
	"clean_architecture/golang/domains"
	"clean_architecture/golang/domains/blog"
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
	//var newBlog *domains.Blog
	//var err error

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

	//// TODO: newするタイミングでValidationを自動で発火させたい
	//blog = &domains.Blog{
	//	Title: uc.InputPort.Title,
	//	Body:  uc.InputPort.Body,
	//}

	//err = i.validator.Validate(blog)
	//if err != nil {
	//	uc.OutputPort.Raise(domains.UnprocessableEntity, err)
	//	return
	//}

	createdBlog, err := i.blogRW.Create(newBlog)
	if err != nil {
		uc.OutputPort.Raise(domains.UnprocessableEntity, err)
		return
	}

	uc.OutputPort.CreateBlog(createdBlog)
}
