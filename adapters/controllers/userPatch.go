package controllers

import (
	"net/http"
	"strconv"

	"go-api/adapters/presenters"
	"go-api/adapters/presenters/json"
	uc "go-api/usecases"

	"github.com/gin-gonic/gin"
)

func (rH RouterHandler) userPatch(c *gin.Context) {
	log := rH.log(rH.MethodAndPath(c))

	req := &UserRequest{}
	if err := c.BindJSON(req); err != nil {
		log(err)
		c.Status(http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log(err)
		c.Status(http.StatusBadRequest)
		return
	}

	useCase := uc.EditUserUseCase{
		OutputPort: json.NewPresenter(presenters.New(c)),
		InputPort: uc.EditUserParams{
			ID:    id,
			Name:  *req.User.Name,
			Email: *req.User.Email,
		},
	}
	rH.ucHandler.UserEdit(useCase)

}
