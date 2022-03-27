package controllers_test

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	controllers "go-api/adapters/controllers"
	"go-api/adapters/loggers"
	mock "go-api/adapters/uc.mock"

	factories "go-api/test/factories"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"gopkg.in/h2non/baloo.v3"
)

func TestDeleteUser_Success(t *testing.T) {
	var userDeletePath = "/api/users/"
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	user := factories.User()
	ucHandler := mock.NewMockHandler(mockCtrl)
	ucHandler.EXPECT().UserDelete(gomock.Any()).Times(1)
	gE := gin.Default()

	router := controllers.NewRouter(ucHandler)
	router.Logger = loggers.SimpleLogger{}
	router.SetRoutes(gE)

	ts := httptest.NewServer(gE)
	defer ts.Close()

	if err := baloo.New(ts.URL).
		Delete(userDeletePath + strconv.Itoa(user.ID.Value)).
		Expect(t).
		Status(http.StatusOK).
		Done(); err != nil {
		t.Error(err)
	}
}
