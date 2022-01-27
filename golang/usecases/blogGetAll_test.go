package uc_test

import (
	"clean_architecture/golang/adapters/presenters"
	"clean_architecture/golang/adapters/presenters/json"
	mock "clean_architecture/golang/adapters/uc.mock"
	"clean_architecture/golang/testData"
	uc "clean_architecture/golang/usecases"
	"errors"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestBlogGetAllSuccess(t *testing.T) {
	var _blogs = testData.Blogs(5)
	t.Run("most obvious", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		i := mock.NewMockedInteractor(mockCtrl)
		i.BlogRW.EXPECT().GetAll().Return(&_blogs, nil).Times(1)

		// UseCase
		ginContext, _ := gin.CreateTestContext(httptest.NewRecorder())
		pre := presenters.New(ginContext)
		form := json.NewPresenter(pre)
		useCase := uc.GetBlogsUseCase{
			OutputPort: form,
			InputPort:  uc.GetBlogsParams{Limit: 1, Offset: 1},
		}

		expectedBlogs := _blogs.ApplyLimitAndOffset(useCase.InputPort.Limit, useCase.InputPort.Offset)

		assert.Equal(t, 1, len(expectedBlogs))
		assert.Equal(t, "タイトル2", expectedBlogs[0].Title.Value)

		i.GetUCHandler().BlogGetAll(useCase)

		assert.NoError(t, nil)
		assert.NoError(t, form.Present())
	})
}

func TestBlogGetAllFails(t *testing.T) {
	var _blogs = testData.Blogs(5)
	mutations := map[string]mock.Tester{
		"shouldPass": {
			Calls: func(i *mock.Interactor) { // change nothing
			},
			ShouldPass: true},
		"error return on uRW.GetFiltered": {
			Calls: func(i *mock.Interactor) {
				i.BlogRW.EXPECT().GetAll().Return(nil, errors.New(""))
			}},
	}

	validCalls := func(i *mock.Interactor) {
		i.BlogRW.EXPECT().GetAll().Return(&_blogs, nil).AnyTimes()
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
			useCase := uc.GetBlogsUseCase{
				OutputPort: form,
				InputPort:  uc.GetBlogsParams{Limit: 1, Offset: 4},
			}

			expectedBlogs := _blogs.ApplyLimitAndOffset(useCase.InputPort.Limit, useCase.InputPort.Offset)
			count := len(expectedBlogs)
			assert.Equal(t, 0, count)

			i.GetUCHandler().BlogGetAll(useCase)

			if mutation.ShouldPass {
				assert.NoError(t, nil)
				return
			}

			assert.NoError(t, nil)
			assert.Error(t, form.Present())
		})
	}
}
