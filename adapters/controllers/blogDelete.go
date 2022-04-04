package controllers

import (
	"net/http"
	"strconv"

	"go-api/adapters/presenters"
	"go-api/adapters/presenters/json"
	"go-api/usecases"

	"github.com/gin-gonic/gin"
)

func (rH RouterHandler) blogDelete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		rH.drivers.Logger.Errorf(ctx, err.Error())
		ctx.Status(http.StatusBadRequest)
		return
	}

	useCase := usecases.NewDeleteBlogUseCase(
		ctx,
		rH.drivers.Logger,
		json.NewPresenter(presenters.New(ctx)),
		rH.drivers.BlogDao,
	)

	useCase.BlogDelete(usecases.DeleteBlogParams{ID: id})
}
