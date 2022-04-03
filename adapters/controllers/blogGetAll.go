package controllers

import (
	"strconv"

	"go-api/adapters/presenters"
	"go-api/adapters/presenters/json"
	"go-api/domains/usecases"

	"github.com/gin-gonic/gin"
)

const (
	defaultLimit  = 20
	defaultOffset = 0
)

func (rH RouterHandler) blogsGetAll(c *gin.Context) {
	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		rH.drivers.Logger.Warnf(c, "blogsGetAll", err)
		limit = defaultLimit
	}

	offset, err := strconv.Atoi(c.Query("offset"))
	if err != nil {
		rH.drivers.Logger.Warnf(c, "blogsGetAll", err)
		offset = defaultOffset
	}

	// TODO(okubo): からでもエラー起きないので、そこは直したい
	useCase := usecases.GetBlogsUseCase{
		OutputPort: json.NewPresenter(presenters.New(c)),
		BlogDao:    rH.drivers.BlogDao,
	}

	useCase.BlogGetAll(usecases.GetBlogsParams{Limit: limit, Offset: offset})
}
