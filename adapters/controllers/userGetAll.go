package controllers

import (
	"strconv"

	"go-api/adapters/presenters"
	"go-api/adapters/presenters/json"
	"go-api/domains/usecases"

	"github.com/gin-gonic/gin"
)

func (rH RouterHandler) userGetAll(c *gin.Context) {
	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		rH.drivers.Logger.Warnf(c, err.Error())
		limit = defaultLimit
	}

	offset, err := strconv.Atoi(c.Query("offset"))
	if err != nil {
		rH.drivers.Logger.Warnf(c, err.Error())
		offset = defaultOffset
	}

	useCase := usecases.GetUsersUseCase{
		OutputPort: json.NewPresenter(presenters.New(c)),
		UserDao:    rH.drivers.UserDao,
	}

	useCase.UserGetAll(usecases.GetUsersParams{Limit: limit, Offset: offset})
}
