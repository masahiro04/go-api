package uc_test

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

func TestGetBlog_Success(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	blog := factories.Blog()

	i := mock.NewMockedInteractor(mockCtrl)
	i.BlogDao.EXPECT().GetById(blog.ID.Value).Return(&blog, nil).Times(1)

	ginContext, _ := gin.CreateTestContext(httptest.NewRecorder())
	pre := presenters.New(ginContext)
	form := json.NewPresenter(pre)
	useCase := uc.GetBlogUseCase{
		OutputPort: form,
		InputPort:  uc.GetBlogParams{Id: blog.ID.Value},
	}

	i.GetUCHandler().BlogGet(useCase)

	assert.NoError(t, nil)
	assert.NoError(t, form.Present())
}

func TestGetBlog_Fail(t *testing.T) {
	blog := factories.Blog()

	mutations := map[string]mock.Tester{
		"error return on BlogDao.GetById": {
			Calls: func(i *mock.Interactor) {
				i.BlogDao.EXPECT().GetById(blog.ID.Value).Return(nil, errors.New(""))
			}},
		"nil nil return on BlogDao.GetById": {
			Calls: func(i *mock.Interactor) {
				i.BlogDao.EXPECT().GetById(blog.ID.Value).Return(nil, nil)
			}},
	}

	validCalls := func(i *mock.Interactor) {
		i.BlogDao.EXPECT().GetById(blog.ID.Value).Return(&blog, nil).AnyTimes()
	}

	for testName, mutation := range mutations {
		t.Run(testName, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			i := mock.NewMockedInteractor(mockCtrl)
			mutation.Calls(&i)
			validCalls(&i)

			// UseCase
			ginContext, _ := gin.CreateTestContext(httptest.NewRecorder())
			pre := presenters.New(ginContext)
			form := json.NewPresenter(pre)
			useCase := uc.GetBlogUseCase{
				OutputPort: form,
				InputPort:  uc.GetBlogParams{Id: blog.ID.Value},
			}

			i.GetUCHandler().BlogGet(useCase)

			if mutation.ShouldPass {
				assert.NoError(t, nil)
				return
			}

			assert.NoError(t, nil)
			assert.Error(t, form.Present())
		})
	}
}
