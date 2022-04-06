package controllers

import (
	"net/http"
	"strconv"

	"go-api/adapters/presenters"
	"go-api/adapters/presenters/json"
	"go-api/usecases"

	"github.com/gin-gonic/gin"
)

func (rH RouterHandler) blogPatch(ctx *gin.Context) {
	req := &BlogRequest{}
	if err := ctx.BindJSON(req); err != nil {
		rH.drivers.Logger.Errorf(ctx, err.Error())
		ctx.Status(http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		rH.drivers.Logger.Errorf(ctx, err.Error())
		ctx.Status(http.StatusBadRequest)
		return
	}

	useCase := usecases.NewEditBlogUseCase(
		ctx,
		rH.drivers.Logger,
		json.NewPresenter(presenters.New(ctx)),
		rH.drivers.BlogDao,
	)
	useCase.BlogEdit(usecases.EditBlogParams{
		ID:    id,
		Title: req.Blog.Title,
		Body:  req.Blog.Body,
	})

}
