package json

import "go-api/domains/models"

type getGetUsersResponse struct {
	Users []getUserResponseItem `json:"response"`
	Count int                   `json:"count"`
}

func (presenter ResponsePresenter) GetUsers(users *models.Users) {
	// nilではなく空配列でレスポンスを返せるようにする
	response := getGetUsersResponse{
		Users: []getUserResponseItem{},
		Count: users.Size(),
	}

	for _, user := range users.Value {
		response.Users = append(response.Users, UserItem(&user))
	}

	presenter.Presenter.StatusOK(response)
}
