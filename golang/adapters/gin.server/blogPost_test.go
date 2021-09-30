package server_test

import (
	"net/http/httptest"
	"testing"

	server "clean_architecture/golang/adapters/gin.server"
	mock "clean_architecture/golang/adapters/uc.mock"
	"clean_architecture/golang/testData"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"gopkg.in/h2non/baloo.v3"
)

var blogPostPath = "/api/blogs"

func TestBlogPost_happyCase(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	// testData
	blog := testData.Blog()

	// ucMock
	ucHandler := mock.NewMockHandler(mockCtrl)
	ucHandler.EXPECT().BlogCreate(gomock.Any()).Times(1)

	// server
	gE := gin.Default()
	server.NewRouter(ucHandler).SetRoutes(gE)
	ts := httptest.NewServer(gE)
	defer ts.Close()

	if err := baloo.New(ts.URL).
		Post(blogPostPath).
		BodyString(`
		{
  			"blog": {
    			"title": "` + blog.Title + `",
    			"body": "` + blog.Body + `"
  			}
		}`).
		Expect(t).
		Done(); err != nil {
		t.Error(err)
	}
}
