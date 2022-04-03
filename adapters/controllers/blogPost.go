package controllers

import (
	"net/http"

	"go-api/adapters/presenters"
	"go-api/adapters/presenters/json"
	"go-api/domains/usecases"

	"github.com/gin-gonic/gin"
)

type BlogRequest struct {
	Blog struct {
		Title string `json:"title" binding:"required"`
		Body  string `json:"body"`
	} `json:"blog" binding:"required"`
}

func (rH RouterHandler) blogPost(c *gin.Context) {
	req := &BlogRequest{}
	if err := c.BindJSON(req); err != nil {
		rH.drivers.Logger.Errorf(c, err.Error())
		c.Status(http.StatusBadRequest)
		return
	}

	useCase := usecases.CreateBlogUseCase{
		OutputPort: json.NewPresenter(presenters.New(c)),
		BlogDao:    rH.drivers.BlogDao,
	}

	useCase.BlogCreate(usecases.CreateBlogParams{
		Title: req.Blog.Title,
		Body:  req.Blog.Body,
	})
}
