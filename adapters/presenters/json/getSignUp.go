package json

import "go-api/domains/models"

type getSignUpResponse struct {
	User getUserResponseItem `json:"user"`
}

func (presenter ResponsePresenter) CreateSignUp(user *models.User) {
	response := getSignUpResponse{User: UserItem(user)}
	presenter.Presenter.Created(response)
}
