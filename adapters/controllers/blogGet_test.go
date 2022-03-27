package controllers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	controllers "go-api/adapters/controllers"
	mock "go-api/adapters/uc.mock"

	"gopkg.in/h2non/baloo.v3"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
)

func TestGetBlog_Success(t *testing.T) {
	url := "/api/blogs/1"
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	ucHandler := mock.NewMockHandler(mockCtrl)

	// TODO(okubo): useCaseを引数に入れたいが、address参照になっているので、一致せずに落ちる
	// とはいえ、この辺りテストしないと意味ないので、ここは厳重に扱いたい
	ucHandler.EXPECT().BlogGet(gomock.Any()).Times(1)
	gE := gin.Default()

	controllers.NewRouter(ucHandler).SetRoutes(gE)

	ts := httptest.NewServer(gE)
	defer ts.Close()

	if err := baloo.New(ts.URL).
		Get(url).
		Expect(t).
		Status(http.StatusOK).
		Done(); err != nil {
		t.Error(err)
	}
}
