package controllers

import (
	"net/http"
	"strconv"

	formatter "clean_architecture/golang/adapters/json.formatter"
	"clean_architecture/golang/adapters/presenters"
	uc "clean_architecture/golang/usecases"

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
		OutputPort: formatter.NewPresenter(presenters.New(c)),
		InputPort: uc.DeleteBlogParams{
			Id: id,
		},
	}
	rH.ucHandler.BlogDelete(useCase)
}
