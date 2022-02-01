package json

import (
	"clean_architecture/golang/domains"
	"fmt"
)

type getUserResponse struct {
	User getUserResponseItem `json:"response"`
}

type getUserResponseItem struct {
	Id        interface{} `json:"id"`
	Name      interface{} `json:"name"`
	Email     interface{} `json:"email"`
	CreatedAt interface{} `json:"createdAt"`
	UpdatedAt interface{} `json:"updatedAt"`
}

func (presenter ResponsePresenter) GetUser(user *domains.User) {
	response := getUserResponse{User: UserItem(user)}
	presenter.Presenter.StatusOK(response)
}

func (presenter ResponsePresenter) CreateUser(user *domains.User) {
	response := getUserResponse{User: UserItem(user)}
	presenter.Presenter.Created(response)
}

func UserItem(user *domains.User) getUserResponseItem {
	fmt.Println(user.CreatedAt)
	fmt.Println("user.CreatedAt")
	return getUserResponseItem{
		Id:        user.ID.Value,
		Name:      user.Name.Value,
		Email:     user.Email.Value,
		CreatedAt: user.CreatedAt.UTC().Format(dateLayout),
		UpdatedAt: user.UpdatedAt.UTC().Format(dateLayout),
	}
}
