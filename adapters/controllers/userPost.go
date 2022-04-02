package controllers

import (
	"net/http"

	"go-api/adapters/presenters"
	"go-api/adapters/presenters/json"
	"go-api/domains/usecases"

	"github.com/gin-gonic/gin"
)

type UserRequest struct {
	User struct {
		Name  *string `json:"name" binding:"required"`
		Email *string `json:"email"`
	} `json:"user" binding:"required"`
}

func (rH RouterHandler) userPost(c *gin.Context) {
	log := rH.log(rH.MethodAndPath(c))
	req := &UserRequest{}

	if err := c.BindJSON(req); err != nil {
		log(err)
		c.Status(http.StatusBadRequest)
		return
	}

	useCase := usecases.CreateUserUseCase{
		OutputPort: json.NewPresenter(presenters.New(c)),
		UserDao:    rH.drivers.UserDao,
	}

	useCase.UserCreate(usecases.CreateUserParams{
		Name:  *req.User.Name,
		Email: *req.User.Email,
	})
}
