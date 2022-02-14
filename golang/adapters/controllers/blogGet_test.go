package controllers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	controllers "clean_architecture/golang/adapters/controllers"
	mock "clean_architecture/golang/adapters/uc.mock"

	"gopkg.in/h2non/baloo.v3"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
)

func TestBlogsGetSuccess(t *testing.T) {
	url := "/api/blogs/1"
	// c, _ := gin.CreateTestContext(httptest.NewRecorder())
	//
	// blog := testData.Blog()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	ucHandler := mock.NewMockHandler(mockCtrl)

	// useCase := uc.GetBlogUseCase{
	// 	OutputPort: json.NewPresenter(presenters.New(c)),
	// 	InputPort:  uc.GetBlogParams{Id: blog.ID.Value},
	// }

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
		//JSONSchema(testData.CompanySingleRespDefinition).
		Status(http.StatusOK).
		Done(); err != nil {
		t.Error(err)
	}
}
