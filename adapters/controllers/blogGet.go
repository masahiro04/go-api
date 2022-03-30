package controllers

import (
	"net/http"
	"strconv"

	uc "go-api/usecases"

	"go-api/adapters/presenters"
	"go-api/adapters/presenters/json"

	"github.com/gin-gonic/gin"
)

func (rH RouterHandler) blogGet(c *gin.Context) {
	log := rH.log(rH.MethodAndPath(c))

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log(err)
		c.Status(http.StatusBadRequest)
		return
	}

	useCase := uc.GetBlogUseCase{
		OutputPort: json.NewPresenter(presenters.New(c), log),
		InputPort:  uc.GetBlogParams{Id: id},
	}
	rH.ucHandler.BlogGet(useCase)
}
