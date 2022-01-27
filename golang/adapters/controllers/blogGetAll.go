package controllers

import (
	"strconv"

	"clean_architecture/golang/adapters/presenters"
	"clean_architecture/golang/adapters/presenters/json"
	uc "clean_architecture/golang/usecases"

	"github.com/gin-gonic/gin"
)

const (
	defaultLimit  = 20
	defaultOffset = 0
)

func (rH RouterHandler) blogsGetAll(c *gin.Context) {
	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		limit = defaultLimit
	}

	offset, err := strconv.Atoi(c.Query("offset"))
	if err != nil {
		offset = defaultOffset
	}

	useCase := uc.GetBlogsUseCase{
		OutputPort: json.NewPresenter(presenters.New(c)),
		InputPort:  uc.GetBlogsParams{Limit: limit, Offset: offset},
	}

	rH.ucHandler.BlogGetAll(useCase)
}
