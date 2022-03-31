package controllers

import (
	"net/http"

	"go-api/adapters/presenters"
	"go-api/adapters/presenters/json"
	uc "go-api/usecases"

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

	useCase := uc.CreateUserUseCase{
		OutputPort: json.NewPresenter(presenters.New(c)),
		InputPort: uc.CreateUserParams{
			Name:  *req.User.Name,
			Email: *req.User.Email,
		},
	}

	rH.ucHandler.UserCreate(useCase)
}
