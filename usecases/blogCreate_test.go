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

func TestCreateBlog_Success(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	blog := factories.Blog()

	// Mock
	i := mock.NewMockedInteractor(mockCtrl)
	i.BlogDao.EXPECT().Create(gomock.Any()).Return(&blog, nil).Times(1)

	// UseCase
	ginContext, _ := gin.CreateTestContext(httptest.NewRecorder())
	pre := presenters.New(ginContext)
	form := json.NewPresenter(pre)
	useCase := uc.CreateBlogUseCase{
		OutputPort: form,
		InputPort: uc.CreateBlogParams{
			Title: blog.Title.Value,
			Body:  blog.Body.Value,
		},
	}

	i.GetUCHandler().BlogCreate(useCase)

	assert.NoError(t, form.Present())
}

func TestCreateBlog_Fail(t *testing.T) {
	blog := factories.Blog()

	mutations := map[string]mock.Tester{
		"shouldPass": {
			Calls: func(i *mock.Interactor) { // change nothing
			},
			ShouldPass: true},
		// "company not validated": {
		// 	Calls: func(i *mock.Interactor) {
		// 		i.Validator.EXPECT().Validate(gomock.Any()).Return(errors.New(""))
		// 	}},
		"error return on uRW.GetByUuId": {
			Calls: func(i *mock.Interactor) {
				i.BlogDao.EXPECT().Create(gomock.Any()).Return(nil, errors.New("")).Times(1)
			}},
	}

	validCalls := func(i *mock.Interactor) {
		i.BlogDao.EXPECT().Create(gomock.Any()).Return(&blog, nil).AnyTimes()
		// i.Validator.EXPECT().Validate(gomock.Any()).Return(nil).AnyTimes()
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
			useCase := uc.CreateBlogUseCase{
				OutputPort: form,
				InputPort: uc.CreateBlogParams{
					Title: blog.Title.Value,
					Body:  blog.Body.Value,
				},
			}

			i.GetUCHandler().BlogCreate(useCase)

			if mutation.ShouldPass {
				assert.NoError(t, nil)
				return
			}

			assert.Error(t, form.Present())
		})
	}
}
