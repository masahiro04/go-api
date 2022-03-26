package controllers_test

import (
	"net/http/httptest"
	"testing"

	controllers "go-api/adapters/controllers"
	mock "go-api/adapters/uc.mock"
	"go-api/testData"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"gopkg.in/h2non/baloo.v3"
)

func TestUserPostSuccess(t *testing.T) {
	var userPostPath = "/api/users"
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	// testData
	user := testData.User()

	// ucMock
	ucHandler := mock.NewMockHandler(mockCtrl)
	ucHandler.EXPECT().UserCreate(gomock.Any()).Times(1)

	// server
	gE := gin.Default()
	controllers.NewRouter(ucHandler).SetRoutes(gE)
	ts := httptest.NewServer(gE)
	defer ts.Close()

	if err := baloo.New(ts.URL).
		Post(userPostPath).
		BodyString(`
		{
  			"user": {
    			"name": "` + user.Name.Value + `",
    			"email": "` + user.Email.Value + `"
  			}
		}`).
		Expect(t).
		Done(); err != nil {
		t.Error(err)
	}
}
