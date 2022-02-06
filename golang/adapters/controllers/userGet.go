package controllers

import (
	"net/http"
	"strconv"

	uc "clean_architecture/golang/usecases"

	"clean_architecture/golang/adapters/presenters"
	"clean_architecture/golang/adapters/presenters/json"

	"github.com/gin-gonic/gin"
)

func (rH RouterHandler) userGet(c *gin.Context) {
	log := rH.log(rH.MethodAndPath(c))

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log(err)
		c.Status(http.StatusBadRequest)
		return
	}

	useCase := uc.GetUserUseCase{
		OutputPort: json.NewPresenter(presenters.New(c)),
		InputPort:  uc.GetUserParams{ID: id},
	}
	rH.ucHandler.UserGet(useCase)
}
