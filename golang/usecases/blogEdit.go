package uc

import (
	"clean_architecture/golang/domain"
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

func (req EditBlogParams) getEditableFields() map[domain.BlogUpdatableProperty]*string {
	return map[domain.BlogUpdatableProperty]*string{
		domain.BlogTitle: &req.Title,
		domain.BlogBody:  &req.Body,
	}
}

func (i interactor) BlogEdit(uc EditBlogUseCase) {
	var blog *domain.Blog
	var err error

	blog, err = i.blogRW.GetById(uc.InputPort.Id)
	if err != nil {
		uc.OutputPort.Raise(domain.BadRequest, err)
		return
	}
	if blog == nil {
		uc.OutputPort.Raise(domain.NotFound, errNotFound)
		return
	}
	if blog.ID != uc.InputPort.Id {
		uc.OutputPort.Raise(domain.UnprocessableEntity, errWrongCompany)
		return
	}

	fieldsToUpdate := uc.InputPort.getEditableFields()
	domain.UpdateBlog(blog,
		domain.SetBlogTitle(fieldsToUpdate[domain.BlogTitle]),
		domain.SetBlogBody(fieldsToUpdate[domain.BlogBody]),
	)

	err = i.validator.Validate(*blog)
	if err != nil {
		uc.OutputPort.Raise(domain.UnprocessableEntity, err)
		return
	}

	updatedBlog, err := i.blogRW.Update(uc.InputPort.Id, *blog)
	if err != nil {
		uc.OutputPort.Raise(domain.UnprocessableEntity, err)
		return
	}

	uc.OutputPort.GetBlog(updatedBlog)
}
