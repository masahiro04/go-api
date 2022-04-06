package controllers

import (
	"net/http"

	"go-api/adapters/presenters"
	"go-api/adapters/presenters/json"
	"go-api/usecases"

	"github.com/gin-gonic/gin"
)

type UserRequest struct {
	User struct {
		Name  *string `json:"name" binding:"required"`
		Email *string `json:"email"`
	} `json:"user" binding:"required"`
}

func (rH RouterHandler) userPost(ctx *gin.Context) {
	req := &UserRequest{}
	if err := ctx.BindJSON(req); err != nil {
		rH.drivers.Logger.Errorf(ctx, err.Error())
		ctx.Status(http.StatusBadRequest)
		return
	}

	useCase := usecases.NewCreateUserUseCase(
		ctx,
		rH.drivers.Logger,
		json.NewPresenter(presenters.New(ctx)),
		rH.drivers.UserDao,
	)

	useCase.UserCreate(usecases.CreateUserParams{
		Name:  *req.User.Name,
		Email: *req.User.Email,
	})
}
