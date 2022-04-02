package usecases

import (
	"fmt"
	"go-api/domains"
	"go-api/domains/models"
)

type GetBlogsUseCase struct {
	OutputPort domains.PresenterRepository
	BlogDao    domains.BlogRepository
}

type GetBlogsParams struct {
	Limit  int
	Offset int
}

func (uc GetBlogsUseCase) BlogGetAll(params GetBlogsParams) {
	fmt.Println("sentinel1")
	// TODO(okubo): BlogDaoがnilになっている。・・・
	fmt.Println(uc.BlogDao)
	blogs, err := uc.BlogDao.GetAll()
	if err != nil {
		uc.OutputPort.Raise(models.BadRequest, err)
		return
	}

	blogs.ApplyLimitAndOffset(params.Limit, params.Offset)

	uc.OutputPort.GetBlogs(blogs)
}
