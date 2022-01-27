package controllers

import (
	"net/http"
	"strconv"

	"clean_architecture/golang/adapters/presenters"
	"clean_architecture/golang/adapters/presenters/json"
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
		OutputPort: json.NewPresenter(presenters.New(c)),
		InputPort: uc.DeleteBlogParams{
			Id: id,
		},
	}
	rH.ucHandler.BlogDelete(useCase)
}
