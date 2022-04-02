package usecases

import (
	"go-api/domains"
	"go-api/domains/models"
)

// NOTE(okubo): OutputPort
type DeleteBlogUseCase struct {
	OutputPort domains.PresenterRepository
	BlogDao    domains.BlogRepository
}

// NOTE(okubo): InputPort
type DeleteBlogParams struct {
	Id int
}

// NOTE(okubo): OutputPort(出力) と InputPort(入力) を結びつける = interactor
func (uc DeleteBlogUseCase) BlogDelete(params DeleteBlogParams) {
	if err := uc.BlogDao.Delete(params.Id); err != nil {
		uc.OutputPort.Raise(models.BadRequest, err)
		return
	}
}
