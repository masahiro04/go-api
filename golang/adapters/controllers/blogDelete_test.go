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

var blogDeletePath = "/api/blogs/"

func TestBlogDeleteSuccess(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	blog := testData.Blog()
	ucHandler := mock.NewMockHandler(mockCtrl)
	ucHandler.EXPECT().BlogDelete(gomock.Any()).Times(1)
	gE := gin.Default()

	router := controllers.NewRouter(ucHandler)
	router.Logger = loggers.SimpleLogger{}
	router.SetRoutes(gE)

	ts := httptest.NewServer(gE)
	defer ts.Close()

	if err := baloo.New(ts.URL).
		Delete(blogDeletePath + strconv.Itoa(blog.ID.Value)).
		Expect(t).
		Status(http.StatusOK).
		Done(); err != nil {
		t.Error(err)
	}
}
