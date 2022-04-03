package controllers

import (
	"net/http"
	"strconv"

	"go-api/adapters/presenters"
	"go-api/adapters/presenters/json"
	"go-api/domains/usecases"

	"github.com/gin-gonic/gin"
)

func (rH RouterHandler) blogGet(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		rH.drivers.Logger.Errorf(c, err.Error())
		c.Status(http.StatusBadRequest)
		return
	}

	useCase := usecases.GetBlogUseCase{
		OutputPort: json.NewPresenter(presenters.New(c)),
		BlogDao:    rH.drivers.BlogDao,
	}
	useCase.BlogGet(usecases.GetBlogParams{Id: id})
}
