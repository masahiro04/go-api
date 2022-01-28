package uc

import (
	"clean_architecture/golang/domains"
	blog2 "clean_architecture/golang/domains/blog"
	"fmt"
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

// func (req EditBlogParams) getEditableFields() map[domains.BlogUpdatableProperty]*string {
// 	return map[domains.BlogUpdatableProperty]*string{
// 		domains.BlogTitle: &req.Title,
// 		domains.BlogBody:  &req.Body,
// 	}
// }

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

	id, _ := blog2.NewId(uc.InputPort.Id)
	// NOTE(okubo): input portで検索している -> どう考えてもerrは起きない
	// if blog.ID != id {
	// 	uc.OutputPort.Raise(domains.UnprocessableEntity, errWrongCompany)
	// 	return
	// }

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
	fmt.Println("---------")
	fmt.Println(uc.InputPort)
	updatedBlog, err := i.blogDao.Update(uc.InputPort.Id, domains.BuildBlog(id, *title, *body))
	fmt.Println(err)
	fmt.Println("---------")
	if err != nil {
		fmt.Println("error occurrred")
		uc.OutputPort.Raise(domains.UnprocessableEntity, err)
		return
	}

	uc.OutputPort.GetBlog(updatedBlog)
}
