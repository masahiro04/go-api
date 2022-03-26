package uc

import (
	"go-api/domains"
	blog2 "go-api/domains/blog"
)

type EditBlogUseCase struct {
	OutputPort Presenter
	InputPort  EditBlogParams
}

type EditBlogParams struct {
	Id    int
	Title string
	Body  string
}

func (i interactor) BlogEdit(uc EditBlogUseCase) {
	var blog *domains.Blog
	var err error

	blog, err = i.blogDao.GetById(uc.InputPort.Id)
	if err != nil {
		uc.OutputPort.Raise(domains.BadRequest, err)
		return
	}

	if blog == nil {
		uc.OutputPort.Raise(domains.NotFound, errNotFound)
		return
	}

	// NOTE(okubo): input portで検索している -> どう考えてもerrは起きない
	id, _ := blog2.NewId(uc.InputPort.Id)

	title, err := blog2.UpdateTitle(&uc.InputPort.Title)
	if err != nil {
		uc.OutputPort.Raise(domains.UnprocessableEntity, err)
		return
	}

	body, err := blog2.UpdateBody(&uc.InputPort.Body)
	if err != nil {
		uc.OutputPort.Raise(domains.UnprocessableEntity, err)
		return
	}

	updatedBlog, err := i.blogDao.Update(
		uc.InputPort.Id, domains.BuildBlog(id, *title, *body, blog.CreatedAt, blog.UpdatedAt),
	)
	if err != nil {
		uc.OutputPort.Raise(domains.UnprocessableEntity, err)
		return
	}

	uc.OutputPort.GetBlog(updatedBlog)
}
