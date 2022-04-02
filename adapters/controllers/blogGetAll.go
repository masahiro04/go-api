package controllers

import (
	"fmt"
	"strconv"

	"go-api/adapters/presenters"
	"go-api/adapters/presenters/json"
	"go-api/domains/usecases"

	"github.com/gin-gonic/gin"
)

const (
	defaultLimit  = 20
	defaultOffset = 0
)

func (rH RouterHandler) blogsGetAll(c *gin.Context) {
	log := rH.log(rH.MethodAndPath(c))
	fmt.Println("inside controller")
	fmt.Println(rH.driver.BlogDao)
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

	// TODO(okubo): からでもエラー起きないので、そこは直したい
	useCase := usecases.GetBlogsUseCase{
		OutputPort: json.NewPresenter(presenters.New(c)),
		BlogDao:    rH.driver.BlogDao,
	}

	useCase.BlogGetAll(usecases.GetBlogsParams{Limit: limit, Offset: offset})
}
