package server_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	server "clean_architecture/golang/adapters/gin.server"
	mock "clean_architecture/golang/adapters/uc.mock"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"gopkg.in/h2non/baloo.v3"
)

var blogGetAllPath = "/api/blogs"

func TestCompanyGetAllSuccess(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	ucHandler := mock.NewMockHandler(mockCtrl)
	ucHandler.EXPECT().BlogGetAll(gomock.Any()).Times(1)
	gE := gin.Default()

	server.NewRouter(ucHandler).SetRoutes(gE)

	ts := httptest.NewServer(gE)
	defer ts.Close()

	if err := baloo.New(ts.URL).
		Get(blogGetAllPath).
		Expect(t).
		//JSONSchema(testData.CompanyMultipleRespDefinition).
		Status(http.StatusOK).
		Done(); err != nil {
		t.Error(err)
	}
}
