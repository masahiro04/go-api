package controllers

import (
	"net/http"
	"strconv"

	"go-api/adapters/presenters"
	"go-api/adapters/presenters/json"
	"go-api/domains"
	"go-api/usecases"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func blogDelete(ctx *gin.Context, logger domains.Logger, db *gorm.DB) {
	// log := rH.log(rH.MethodAndPath(c))

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		// log(err)
		ctx.Status(http.StatusBadRequest)
		return
	}

	useCase := usecases.DeleteBlogUseCase{
		OutputPort: json.NewPresenter(presenters.New(ctx)),
		InputPort: usecases.DeleteBlogParams{
			Id: id,
		},
	}
	usecases.BlogDelete(useCase)
}
