package controllers

import (
	"net/http"

	"go-api/adapters/presenters"
	"go-api/adapters/presenters/json"
	"go-api/usecases"

	"github.com/gin-gonic/gin"
)

type SignUpRequest struct {
	Name     string  `json:"name"`
	Email    string  `json:"email"`
	Password *string `json:"password"`
}

func (rH RouterHandler) signUp(ctx *gin.Context) {
	req := &SignUpRequest{}
	if err := ctx.BindJSON(req); err != nil {
		rH.drivers.Logger.Errorf(ctx, err.Error())
		ctx.Status(http.StatusBadRequest)
		return
	}

	useCase := usecases.NewSignUpUseCase(
		ctx,
		rH.drivers.Logger,
		json.NewPresenter(presenters.New(ctx)),
		rH.drivers.UserDao,
		rH.drivers.DBTransaction,
		rH.drivers.FirebaseHandler,
	)

	useCase.SignUp(usecases.SignUpParams{
		Name:     req.Name,
		Email:    req.Email,
		Password: *req.Password,
	})
}
