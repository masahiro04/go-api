package json

import (
	"go-api/domains"
)

type getSignUpResponse struct {
	User getUserResponseItem `json:"user"`
}

func (presenter ResponsePresenter) CreateSignUp(user *domains.User) {
	response := getSignUpResponse{User: UserItem(user)}
	presenter.Presenter.Created(response)
}
