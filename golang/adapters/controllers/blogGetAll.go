package controllers

import (
	"strconv"

	formatter "clean_architecture/golang/adapters/json.formatter"
	"clean_architecture/golang/adapters/presenters"
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
		OutputPort: formatter.NewPresenter(presenters.New(c)),
		InputPort:  uc.GetBlogsParams{Limit: limit, Offset: offset},
	}

	rH.ucHandler.BlogGetAll(useCase)
}
