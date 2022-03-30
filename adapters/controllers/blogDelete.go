package controllers

import (
	"net/http"
	"strconv"

	"go-api/adapters/presenters"
	"go-api/adapters/presenters/json"
	uc "go-api/usecases"

	"github.com/gin-gonic/gin"
)

func (rH RouterHandler) blogDelete(c *gin.Context) {
	log := rH.log(rH.MethodAndPath(c))

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log(err)
		c.Status(http.StatusBadRequest)
		return
	}

	useCase := uc.DeleteBlogUseCase{
		OutputPort: json.NewPresenter(presenters.New(c), log),
		InputPort: uc.DeleteBlogParams{
			Id: id,
		},
	}
	rH.ucHandler.BlogDelete(useCase)
}
