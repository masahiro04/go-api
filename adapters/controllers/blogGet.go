package controllers

import (
	"net/http"
	"strconv"

	"go-api/adapters/presenters"
	"go-api/adapters/presenters/json"
	"go-api/domains/usecases"

	"github.com/gin-gonic/gin"
)

func (rH RouterHandler) blogGet(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		rH.drivers.Logger.Errorf(ctx, err.Error())
		ctx.Status(http.StatusBadRequest)
		return
	}

	useCase := usecases.NewGetBlogUseCase(
		ctx,
		rH.drivers.Logger,
		json.NewPresenter(presenters.New(ctx)),
		rH.drivers.BlogDao,
	)
	useCase.BlogGet(usecases.GetBlogParams{ID: id})
}
