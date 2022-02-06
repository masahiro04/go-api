package controllers_test

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	controllers "clean_architecture/golang/adapters/controllers"
	"clean_architecture/golang/adapters/loggers"
	mock "clean_architecture/golang/adapters/uc.mock"
	"clean_architecture/golang/testData"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"gopkg.in/h2non/baloo.v3"
)

func TestUserPatchSuccess(t *testing.T) {
	var userPatchPath = "/api/users/"
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	user := testData.User()
	ucHandler := mock.NewMockHandler(mockCtrl)
	ucHandler.EXPECT().UserEdit(gomock.Any()).Times(1)

	gE := gin.Default()

	router := controllers.NewRouter(ucHandler)
	router.Logger = loggers.SimpleLogger{}
	router.SetRoutes(gE)

	ts := httptest.NewServer(gE)
	defer ts.Close()

	if err := baloo.New(ts.URL).
		Put(userPatchPath + strconv.Itoa(user.ID.Value)).
		BodyString(`{
  			"user": {
    			"id": "` + strconv.Itoa(user.ID.Value) + `",
    			"name": "` + user.Name.Value + `",
    			"email": "` + user.Name.Value + `"
  			}
		}`).
		Expect(t).
		Status(http.StatusOK).
		//JSONSchema(testData.CompanySingleRespDefinition).
		Done(); err != nil {
		t.Error(err)
	}
}
