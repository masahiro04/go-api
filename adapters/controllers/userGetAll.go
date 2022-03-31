package controllers

import (
	"strconv"

	"go-api/adapters/presenters"
	"go-api/adapters/presenters/json"
	uc "go-api/usecases"

	"github.com/gin-gonic/gin"
)

func (rH RouterHandler) userGetAll(c *gin.Context) {
	log := rH.log(rH.MethodAndPath(c))
	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		log(err)
		limit = defaultLimit
	}

	offset, err := strconv.Atoi(c.Query("offset"))
	if err != nil {
		log(err)
		offset = defaultOffset
	}

	useCase := uc.GetUsersUseCase{
		OutputPort: json.NewPresenter(presenters.New(c)),
		InputPort:  uc.GetUsersParams{Limit: limit, Offset: offset},
	}

	rH.ucHandler.UserGetAll(useCase)
}
