package controllers

import (
	"net/http"
	"strconv"

	"go-api/adapters/presenters"
	"go-api/adapters/presenters/json"
	"go-api/domains/usecases"

	"github.com/gin-gonic/gin"
)

func (rH RouterHandler) blogPatch(c *gin.Context) {
	req := &BlogRequest{}
	if err := c.BindJSON(req); err != nil {
		rH.drivers.Logger.Errorf(c, err.Error())
		c.Status(http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		rH.drivers.Logger.Errorf(c, err.Error())
		c.Status(http.StatusBadRequest)
		return
	}

	useCase := usecases.EditBlogUseCase{
		OutputPort: json.NewPresenter(presenters.New(c)),
		BlogDao:    rH.drivers.BlogDao,
	}
	useCase.BlogEdit(usecases.EditBlogParams{
		Id:    id,
		Title: req.Blog.Title,
		Body:  req.Blog.Body,
	})

}
