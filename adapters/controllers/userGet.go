package controllers

import (
	"net/http"
	"strconv"

	uc "go-api/usecases"

	"go-api/adapters/presenters"
	"go-api/adapters/presenters/json"

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
