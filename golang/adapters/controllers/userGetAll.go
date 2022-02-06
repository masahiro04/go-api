package controllers

import (
	"strconv"

	"clean_architecture/golang/adapters/presenters"
	"clean_architecture/golang/adapters/presenters/json"
	uc "clean_architecture/golang/usecases"

	"github.com/gin-gonic/gin"
)

func (rH RouterHandler) userGetAll(c *gin.Context) {
	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		limit = defaultLimit
	}

	offset, err := strconv.Atoi(c.Query("offset"))
	if err != nil {
		offset = defaultOffset
	}

	useCase := uc.GetUsersUseCase{
		OutputPort: json.NewPresenter(presenters.New(c)),
		InputPort:  uc.GetUsersParams{Limit: limit, Offset: offset},
	}

	rH.ucHandler.UserGetAll(useCase)
}
