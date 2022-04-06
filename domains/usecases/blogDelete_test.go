package usecases_test

import (
	"errors"
	"net/http/httptest"
	"testing"

	"go-api/adapters/presenters"
	"go-api/adapters/presenters/json"

	"github.com/gin-gonic/gin"

	mock "go-api/adapters/uc.mock"

	factories "go-api/test/factories"
	uc "go-api/usecases"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestDeleteBlog_Success(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	blog := factories.Blog()

	i := mock.NewMockedInteractor(mockCtrl)
	i.BlogDao.EXPECT().Delete(blog.ID.Value).Return(nil).Times(1)

	ginContext, _ := gin.CreateTestContext(httptest.NewRecorder())
	pre := presenters.New(ginContext)
	form := json.NewPresenter(pre)
	useCase := uc.DeleteBlogUseCase{
		OutputPort: form,
		InputPort: uc.DeleteBlogParams{
			Id: blog.ID.Value,
		},
	}

	i.GetUCHandler().BlogDelete(useCase)

	assert.NoError(t, nil)
	assert.NoError(t, form.Present())
}

func TestDeleteBlog_Fail(t *testing.T) {
	blog := factories.Blog()

	mutations := map[string]mock.Tester{
		"shouldPass": {
			Calls: func(i *mock.Interactor) { // change nothing
			},
			ShouldPass: true},
		"failed to save the company": {
			Calls: func(i *mock.Interactor) {
				i.BlogDao.EXPECT().Delete(blog.ID.Value).Return(errors.New(""))
			}},
	}

	validCalls := func(i *mock.Interactor) {
		i.BlogDao.EXPECT().Delete(blog.ID.Value).Return(nil).AnyTimes()
	}

	for testName, mutation := range mutations {
		t.Run(testName, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			i := mock.NewMockedInteractor(mockCtrl)
			mutation.Calls(&i)
			validCalls(&i)

			ginContext, _ := gin.CreateTestContext(httptest.NewRecorder())
			pre := presenters.New(ginContext)
			form := json.NewPresenter(pre)
			useCase := uc.DeleteBlogUseCase{
				OutputPort: form,
				InputPort: uc.DeleteBlogParams{
					Id: blog.ID.Value,
				},
			}

			i.GetUCHandler().BlogDelete(useCase)

			if mutation.ShouldPass {
				assert.NoError(t, form.Present())
				return
			}

			assert.Error(t, form.Present())
		})
	}
}
