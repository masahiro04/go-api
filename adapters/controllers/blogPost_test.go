package controllers_test

import (
	"net/http/httptest"
	"testing"

	controllers "go-api/adapters/controllers"
	mock "go-api/adapters/uc.mock"

	factories "go-api/test/factories"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"gopkg.in/h2non/baloo.v3"
)

var blogPostPath = "/api/blogs"

func TestPostBlog_Success(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	// factories
	blog := factories.Blog()

	// ucMock
	ucHandler := mock.NewMockHandler(mockCtrl)
	ucHandler.EXPECT().BlogCreate(gomock.Any()).Times(1)

	// server
	gE := gin.Default()
	controllers.NewRouter(ucHandler).SetRoutes(gE)
	ts := httptest.NewServer(gE)
	defer ts.Close()

	if err := baloo.New(ts.URL).
		Post(blogPostPath).
		BodyString(`
		{
  			"blog": {
    			"title": "` + blog.Title.Value + `",
    			"body": "` + blog.Body.Value + `"
  			}
		}`).
		Expect(t).
		Done(); err != nil {
		t.Error(err)
	}
}
