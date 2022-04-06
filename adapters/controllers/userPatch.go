package controllers

import (
	"net/http"
	"strconv"

	"go-api/adapters/presenters"
	"go-api/adapters/presenters/json"
	"go-api/domains/usecases"

	"github.com/gin-gonic/gin"
)

func (rH RouterHandler) userPatch(ctx *gin.Context) {
	req := &UserRequest{}
	if err := ctx.BindJSON(req); err != nil {
		rH.drivers.Logger.Errorf(ctx, err.Error())
		ctx.Status(http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		rH.drivers.Logger.Errorf(ctx, err.Error())
		ctx.Status(http.StatusBadRequest)
		return
	}

	useCase := usecases.NewEditUserUseCase(
		ctx,
		rH.drivers.Logger,
		json.NewPresenter(presenters.New(ctx)),
		rH.drivers.UserDao,
	)

	useCase.UserEdit(usecases.EditUserParams{
		ID:    id,
		Name:  *req.User.Name,
		Email: *req.User.Email,
	})
}
