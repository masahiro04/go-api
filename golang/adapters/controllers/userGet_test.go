package controllers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	controllers "clean_architecture/golang/adapters/controllers"
	mock "clean_architecture/golang/adapters/uc.mock"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"gopkg.in/h2non/baloo.v3"
)

func TestUsersGetSuccess(t *testing.T) {
	var userGetPath = "/api/users/1"
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	ucHandler := mock.NewMockHandler(mockCtrl)
	ucHandler.EXPECT().UserGet(gomock.Any()).Times(1)
	gE := gin.Default()

	controllers.NewRouter(ucHandler).SetRoutes(gE)

	ts := httptest.NewServer(gE)
	defer ts.Close()

	if err := baloo.New(ts.URL).
		Get(userGetPath).
		Expect(t).
		//JSONSchema(testData.CompanySingleRespDefinition).
		Status(http.StatusOK).
		Done(); err != nil {
		t.Error(err)
	}
}
