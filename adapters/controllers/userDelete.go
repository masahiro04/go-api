package controllers

import (
	"net/http"
	"strconv"

	"go-api/adapters/presenters"
	"go-api/adapters/presenters/json"
	uc "go-api/usecases"

	"github.com/gin-gonic/gin"
)

func (rH RouterHandler) userDelete(c *gin.Context) {
	log := rH.log(rH.MethodAndPath(c))

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log(err)
		c.Status(http.StatusBadRequest)
		return
	}

	useCase := uc.DeleteUserUseCase{
		OutputPort: json.NewPresenter(presenters.New(c), log),
		InputPort: uc.DeleteUserParams{
			ID: id,
		},
	}
	rH.ucHandler.UserDelete(useCase)
}
