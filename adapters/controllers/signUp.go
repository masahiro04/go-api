package controllers

import (
	"net/http"

	"go-api/adapters/presenters"
	"go-api/adapters/presenters/json"
	"go-api/domains/usecases"

	"github.com/gin-gonic/gin"
)

type SignUpRequest struct {
	Name     string  `json:"name"`
	Email    string  `json:"email"`
	Password *string `json:"password"`
}

func (rH RouterHandler) signUp(c *gin.Context) {
	log := rH.log(rH.MethodAndPath(c))

	req := &SignUpRequest{}
	if err := c.BindJSON(req); err != nil {
		log(err)
		c.Status(http.StatusBadRequest)
		return
	}

	useCase := usecases.SignUpUseCase{
		OutputPort:      json.NewPresenter(presenters.New(c)),
		UserDao:         rH.drivers.UserDao,
		DBTransaction:   rH.drivers.DBTransaction,
		FirebaseHandler: rH.drivers.FirebaseHandler,
	}

	useCase.SignUp(usecases.SignUpParams{
		Name:     req.Name,
		Email:    req.Email,
		Password: *req.Password,
	})
}
