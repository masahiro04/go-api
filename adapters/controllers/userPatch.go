package controllers

import (
	"net/http"
	"strconv"

	"go-api/adapters/presenters"
	"go-api/adapters/presenters/json"
	"go-api/domains/usecases"

	"github.com/gin-gonic/gin"
)

func (rH RouterHandler) userPatch(c *gin.Context) {
	req := &UserRequest{}
	if err := c.BindJSON(req); err != nil {
		rH.drivers.Logger.Errorf(c, err.Error())
		c.Status(http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		rH.drivers.Logger.Errorf(c, err.Error())
		c.Status(http.StatusBadRequest)
		return
	}

	useCase := usecases.EditUserUseCase{
		OutputPort: json.NewPresenter(presenters.New(c)),
		UserDao:    rH.drivers.UserDao,
	}
	useCase.UserEdit(usecases.EditUserParams{
		ID:    id,
		Name:  *req.User.Name,
		Email: *req.User.Email,
	})

}
