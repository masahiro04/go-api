package usecases

import (
	"go-api/domains"
	"go-api/domains/models"
)

type DeleteBlogUseCase struct {
	OutputPort domains.PresenterRepository
	BlogDao    domains.BlogRepository
}

type DeleteBlogParams struct {
	Id int
}

func (uc DeleteBlogUseCase) BlogDelete(params DeleteBlogParams) {
	if err := uc.BlogDao.Delete(params.Id); err != nil {
		uc.OutputPort.Raise(models.BadRequest, err)
		return
	}
}
