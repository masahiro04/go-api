package controllers

import (
	"strconv"

	"go-api/adapters/presenters"
	"go-api/adapters/presenters/json"
	"go-api/usecases"

	"github.com/gin-gonic/gin"
)

const (
	defaultLimit  = 20
	defaultOffset = 0
)

func (rH RouterHandler) blogsGetAll(ctx *gin.Context) {
	limit, err := strconv.Atoi(ctx.Query("limit"))
	if err != nil {
		rH.drivers.Logger.Warnf(ctx, "blogsGetAll", err)
		limit = defaultLimit
	}

	offset, err := strconv.Atoi(ctx.Query("offset"))
	if err != nil {
		rH.drivers.Logger.Warnf(ctx, "blogsGetAll", err)
		offset = defaultOffset
	}

	useCase := usecases.NewGetBlogsUseCase(
		ctx,
		rH.drivers.Logger,
		json.NewPresenter(presenters.New(ctx)),
		rH.drivers.BlogDao,
	)

	useCase.BlogGetAll(usecases.GetBlogsParams{Limit: limit, Offset: offset})
}
