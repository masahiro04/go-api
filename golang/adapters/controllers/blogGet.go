package controllers

import (
	"net/http"
	"strconv"

	uc "clean_architecture/golang/usecases"

	formatter "clean_architecture/golang/adapters/json.formatter"
	"clean_architecture/golang/adapters/presenters"

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
		OutputPort: formatter.NewPresenter(presenters.New(c)),
		InputPort:  uc.GetBlogParams{Id: id},
	}
	rH.ucHandler.BlogGet(useCase)
}
