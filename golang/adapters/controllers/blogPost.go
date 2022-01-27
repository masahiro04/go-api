package controllers

import (
	"net/http"

	formatter "clean_architecture/golang/adapters/json.formatter"
	"clean_architecture/golang/adapters/presenters"
	uc "clean_architecture/golang/usecases"

	"github.com/gin-gonic/gin"
)

type BlogRequest struct {
	Blog struct {
		Title *string `json:"title" binding:"required"`
		Body  *string `json:"body"`
	} `json:"blog" binding:"required"`
}

func (rH RouterHandler) blogPost(c *gin.Context) {
	log := rH.log(rH.MethodAndPath(c))
	req := &BlogRequest{}

	if err := c.BindJSON(req); err != nil {
		log(err)
		c.Status(http.StatusBadRequest)
		return
	}

	useCase := uc.CreateBlogUseCase{
		OutputPort: formatter.NewPresenter(presenters.New(c)),
		InputPort: uc.CreateBlogParams{
			Title: *req.Blog.Title,
			Body:  *req.Blog.Body,
		},
	}

	rH.ucHandler.BlogCreate(useCase)
}
