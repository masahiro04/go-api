package controllers

import (
	"strconv"

	"go-api/adapters/presenters"
	"go-api/adapters/presenters/json"
	"go-api/usecases"

	"github.com/gin-gonic/gin"
)

func (rH RouterHandler) userGetAll(ctx *gin.Context) {
	limit, err := strconv.Atoi(ctx.Query("limit"))
	if err != nil {
		rH.drivers.Logger.Warnf(ctx, err.Error())
		limit = defaultLimit
	}

	offset, err := strconv.Atoi(ctx.Query("offset"))
	if err != nil {
		rH.drivers.Logger.Warnf(ctx, err.Error())
		offset = defaultOffset
	}

	useCase := usecases.NewGetUsersUseCase(
		ctx,
		rH.drivers.Logger,
		json.NewPresenter(presenters.New(ctx)),
		rH.drivers.UserDao,
	)

	useCase.UserGetAll(usecases.GetUsersParams{Limit: limit, Offset: offset})
}
