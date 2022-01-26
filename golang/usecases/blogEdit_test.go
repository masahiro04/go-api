package uc_test

import (
	"errors"
	"net/http/httptest"
	"testing"

	formatter "clean_architecture/golang/adapters/json.formatter"
	presenter "clean_architecture/golang/adapters/json.presenter"

	"github.com/gin-gonic/gin"

	mock "clean_architecture/golang/adapters/uc.mock"

	"clean_architecture/golang/testData"
	uc "clean_architecture/golang/usecases"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCompanyEdit_happyCase(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	blog := testData.Blog()

	i := mock.NewMockedInteractor(mockCtrl)
	i.BlogRW.EXPECT().GetById(blog.ID).Return(&blog, nil).Times(1)
	i.BlogRW.EXPECT().Update(blog.ID, blog).Return(&blog, nil).Times(1)
	i.Validator.EXPECT().Validate(gomock.Any()).Return(nil).Times(1)

	ginContext, _ := gin.CreateTestContext(httptest.NewRecorder())
	pre := presenter.New(ginContext)
	form := formatter.NewPresenter(pre)

	useCase := uc.EditBlogUseCase{
		OutputPort: form,
		InputPort: uc.EditBlogParams{
			Id:    blog.ID.Value(),
			Title: blog.Title.Value(),
			Body:  blog.Body.Value(),
		},
	}

	i.GetUCHandler().BlogEdit(useCase)

	assert.NoError(t, nil)
	assert.NoError(t, form.Present())
}

func TestCompanyEdit_fails(t *testing.T) {
	blog := testData.Blog()

	mutations := map[string]mock.Tester{
		"shouldPass": {Calls: func(i *mock.Interactor) {
			// change nothing
		}, ShouldPass: true},
		"error return on CompanyRW.GetById": {
			Calls: func(i *mock.Interactor) {
				i.BlogRW.EXPECT().GetById(blog.ID).Return(nil, errors.New(""))
			}},
		"nil, nil return on CompanyRW.GetById": {
			Calls: func(i *mock.Interactor) {
				i.BlogRW.EXPECT().GetById(blog.ID).Return(nil, nil)
			}},
		// TODO エラーハンドリングしっかりしたあとに有効にする
		// "uRW.GetByID returns wrong ID": {
		// 	Calls: func(i *mock.Interactor) {
		// 		i.BlogRW.EXPECT().GetById(blog.ID).Return(&domains.Blog{ID: 12}, nil)
		// 	}},
		"company not validated": {
			Calls: func(i *mock.Interactor) {
				i.Validator.EXPECT().Validate(gomock.Any()).Return(errors.New(""))
			}},
		"failed to save the user": {
			Calls: func(i *mock.Interactor) {
				i.BlogRW.EXPECT().Update(blog.ID, blog).Return(nil, errors.New(""))
			}},
	}

	validCalls := func(i *mock.Interactor) {
		i.Logger.EXPECT().Log(gomock.Any()).AnyTimes()
		i.BlogRW.EXPECT().GetById(blog.ID).Return(&blog, nil).AnyTimes()
		i.BlogRW.EXPECT().Update(blog.ID, blog).Return(&blog, nil).AnyTimes()
		i.Validator.EXPECT().Validate(gomock.Any()).Return(nil).AnyTimes()
	}

	for testName, mutation := range mutations {
		t.Run(testName, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			i := mock.NewMockedInteractor(mockCtrl)
			mutation.Calls(&i)
			validCalls(&i)

			ginContext, _ := gin.CreateTestContext(httptest.NewRecorder())
			pre := presenter.New(ginContext)
			form := formatter.NewPresenter(pre)

			useCase := uc.EditBlogUseCase{
				OutputPort: form,
				InputPort: uc.EditBlogParams{
					Id:    blog.ID.Value(),
					Title: blog.Title.Value(),
					Body:  blog.Body.Value(),
				},
			}

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
