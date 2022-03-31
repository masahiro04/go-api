package controllers

import (
	"net/http"

	"go-api/adapters/presenters"
	"go-api/adapters/presenters/json"
	uc "go-api/usecases"

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

	params := uc.SignUpUseCase{
		OutputPort: json.NewPresenter(presenters.New(c)),
		InputPort: uc.SignUpParams{
			Name:     req.Name,
			Email:    req.Email,
			Password: *req.Password,
		},
	}

	rH.ucHandler.SignUp(params)
}

// import (
// 	"net/http"
// 	"strconv"
//
// 	"go-api/adapters/presenters"
// 	"go-api/adapters/presenters/json"
// 	uc "go-api/usecases"
//
// 	"github.com/gin-gonic/gin"
// )
//
// func (rH RouterHandler) userDelete(c *gin.Context) {
// 	log := rH.log(rH.MethodAndPath(c))
//
// 	id, err := strconv.Atoi(c.Param("id"))
// 	if err != nil {
// 		log(err)
// 		c.Status(http.StatusBadRequest)
// 		return
// 	}
//
// 	useCase := uc.DeleteUserUseCase{
// 		OutputPort: json.NewPresenter(presenters.New(c)),
// 		InputPort: uc.DeleteUserParams{
// 			ID: id,
// 		},
// 	}
// 	rH.ucHandler.UserDelete(useCase)
// }
//
