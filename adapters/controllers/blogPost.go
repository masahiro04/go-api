package controllers

import (
	"net/http"

	"go-api/adapters/presenters"
	"go-api/adapters/presenters/json"
	"go-api/usecases"

	"github.com/gin-gonic/gin"
)

type BlogRequest struct {
	Blog struct {
		Title string `json:"title" binding:"required"`
		Body  string `json:"body"`
	} `json:"blog" binding:"required"`
}

func (rH RouterHandler) blogPost(ctx *gin.Context) {
	req := &BlogRequest{}
	if err := ctx.BindJSON(req); err != nil {
		rH.drivers.Logger.Errorf(ctx, err.Error())
		ctx.Status(http.StatusBadRequest)
		return
	}

	useCase := usecases.NewCreateBlogUseCase(
		ctx,
		rH.drivers.Logger,
		json.NewPresenter(presenters.New(ctx)),
		rH.drivers.BlogDao,
	)

	useCase.BlogCreate(usecases.CreateBlogParams{
		Title: req.Blog.Title,
		Body:  req.Blog.Body,
	})
}
