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

func TestEditBlog_Success(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	blog := factories.Blog()

	i := mock.NewMockedInteractor(mockCtrl)
	i.BlogDao.EXPECT().GetById(blog.ID.Value).Return(&blog, nil).Times(1)
	i.BlogDao.EXPECT().Update(blog.ID.Value, blog).Return(&blog, nil).Times(1)

	ginContext, _ := gin.CreateTestContext(httptest.NewRecorder())
	pre := presenters.New(ginContext)
	form := json.NewPresenter(pre)

	useCase := uc.EditBlogUseCase{
		OutputPort: form,
		InputPort: uc.EditBlogParams{
			Id:    blog.ID.Value,
			Title: blog.Title.Value,
			Body:  blog.Body.Value,
		},
	}

	i.GetUCHandler().BlogEdit(useCase)

	assert.NoError(t, nil)
	assert.NoError(t, form.Present())
}

func TestEditBlog_Fail(t *testing.T) {
	blog := factories.Blog()

	ginContext, _ := gin.CreateTestContext(httptest.NewRecorder())
	pre := presenters.New(ginContext)
	form := json.NewPresenter(pre)

	useCase := uc.EditBlogUseCase{
		OutputPort: form,
		InputPort: uc.EditBlogParams{
			Id:    blog.ID.Value,
			Title: blog.Title.Value,
			Body:  blog.Body.Value,
		},
	}
	mutations := map[string]mock.Tester{
		"error return on BlogDao.GetById": {
			Calls: func(i *mock.Interactor) {
				i.BlogDao.EXPECT().GetById(blog.ID.Value).Return(nil, errors.New(""))
			}},
		"nil, nil return on BlogDao.GetById": {
			Calls: func(i *mock.Interactor) {
				i.BlogDao.EXPECT().GetById(blog.ID.Value).Return(nil, nil)
			}},

		"error returns when blog.title is blank": {
			Calls: func(i *mock.Interactor) {
				useCase.InputPort.Title = ""
				i.BlogDao.EXPECT().GetById(blog.ID.Value).Return(&blog, nil).AnyTimes()
			}},
		"error returns when blog.body is blank": {
			Calls: func(i *mock.Interactor) {
				useCase.InputPort.Id = blog.ID.Value
				useCase.InputPort.Title = blog.Title.Value
				useCase.InputPort.Body = ""
				i.BlogDao.EXPECT().GetById(blog.ID.Value).Return(&blog, nil).AnyTimes()
			}},
		"error returns when blog cannot update": {
			Calls: func(i *mock.Interactor) {
				useCase.InputPort.Id = blog.ID.Value
				useCase.InputPort.Title = blog.Title.Value
				useCase.InputPort.Body = blog.Body.Value

				i.BlogDao.EXPECT().GetById(blog.ID.Value).Return(&blog, nil).AnyTimes()
				i.BlogDao.EXPECT().Update(gomock.Any(), gomock.Any()).Return(nil, errors.New("")).AnyTimes()
			}},
	}

	// validCalls := func(i *mock.Interactor) {
	// 	i.Logger.EXPECT().Log(gomock.Any()).AnyTimes()
	// 	i.BlogDao.EXPECT().GetById(blog.ID.Value).Return(&blog, nil).AnyTimes()
	// 	i.BlogDao.EXPECT().Update(blog.ID.Value, blog).Return(&blog, nil).AnyTimes()
	// 	i.Validator.EXPECT().Validate(gomock.Any()).Return(nil).AnyTimes()
	// }

	for testName, mutation := range mutations {
		t.Run(testName, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			i := mock.NewMockedInteractor(mockCtrl)
			mutation.Calls(&i)
			// validCalls(&i)

			i.GetUCHandler().BlogEdit(useCase)

			if mutation.ShouldPass {
				assert.NoError(t, nil)
				return
			}

			assert.NoError(t, nil)
			assert.Error(t, form.Present())
		})
	}
}
