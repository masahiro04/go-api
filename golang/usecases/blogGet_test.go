package uc_test

import (
	"errors"
	"net/http/httptest"
	"testing"

	"clean_architecture/golang/adapters/presenters"
	"clean_architecture/golang/adapters/presenters/json"

	"github.com/gin-gonic/gin"

	mock "clean_architecture/golang/adapters/uc.mock"
	"clean_architecture/golang/testData"

	uc "clean_architecture/golang/usecases"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestBlogGetSuccess(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	blog := testData.Blog()

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

func TestBlogGetFails(t *testing.T) {
	blog := testData.Blog()

	mutations := map[string]mock.Tester{
		"shouldPass": {
			Calls: func(i *mock.Interactor) { // change nothing
			},
			ShouldPass: true},
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
