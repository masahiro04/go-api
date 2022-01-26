package uc_test

import (
	"testing"
)

// var blog1 = &domains.Blog{ID: 1, Title: "Blog1"}
// var blog2 = &domains.Blog{ID: 2, Title: "Blog2"}
// var blog3 = &domains.Blog{ID: 3, Title: "Blog3"}
// var blog4 = &domains.Blog{ID: 4, Title: "Blog4"}
// var expectedBlogs = domains.BlogCollection{testData.PointeredBlogs()}

func TestBlogGetAll_happyCase(t *testing.T) {
	// t.Run("most obvious", func(t *testing.T) {
	// 	mockCtrl := gomock.NewController(t)
	// 	defer mockCtrl.Finish()
	//
	// 	i := mock.NewMockedInteractor(mockCtrl)
	// 	i.BlogRW.EXPECT().GetAll().Return(expectedBlogs, nil).Times(1)
	//
	// 	// UseCase
	// 	ginContext, _ := gin.CreateTestContext(httptest.NewRecorder())
	// 	pre := presenter.New(ginContext)
	// 	form := formatter.NewPresenter(pre)
	// 	useCase := uc.GetBlogsUseCase{
	// 		OutputPort: form,
	// 		InputPort:  uc.GetBlogsParams{Limit: 1, Offset: 1},
	// 	}
	//
	// 	expectedBlogs = domains.BlogCollection(expectedBlogs).ApplyLimitAndOffset(useCase.InputPort.Limit, useCase.InputPort.Offset)
	// 	count := len(expectedBlogs)
	// 	assert.Equal(t, 1, count)
	// 	assert.Equal(t, blog1.Title, "Blog1")
	//
	// 	i.GetUCHandler().BlogGetAll(useCase)
	//
	// 	assert.NoError(t, nil)
	// 	assert.NoError(t, form.Present())
	// })
}

// func TestBlogGetAll_fails(t *testing.T) {
// 	mutations := map[string]mock.Tester{
// 		"shouldPass": {
// 			Calls: func(i *mock.Interactor) { // change nothing
// 			},
// 			ShouldPass: true},
// 		"error return on uRW.GetFiltered": {
// 			Calls: func(i *mock.Interactor) {
// 				i.BlogRW.EXPECT().GetAll().Return(nil, errors.New(""))
// 			}},
// 	}
//
// 	validCalls := func(i *mock.Interactor) {
// 		i.BlogRW.EXPECT().GetAll().Return(expectedBlogs, nil).AnyTimes()
// 	}
//
// 	for testName, mutation := range mutations {
// 		t.Run(testName, func(t *testing.T) {
// 			mockCtrl := gomock.NewController(t)
// 			defer mockCtrl.Finish()
//
// 			i := mock.NewMockedInteractor(mockCtrl)
// 			mutation.Calls(&i)
// 			validCalls(&i)
//
// 			// UseCase
// 			ginContext, _ := gin.CreateTestContext(httptest.NewRecorder())
// 			pre := presenter.New(ginContext)
// 			form := formatter.NewPresenter(pre)
// 			useCase := uc.GetBlogsUseCase{
// 				OutputPort: form,
// 				InputPort:  uc.GetBlogsParams{Limit: 1, Offset: 4},
// 			}
//
// 			expectedBlogs = domains.BlogCollection(expectedBlogs).ApplyLimitAndOffset(useCase.InputPort.Limit, useCase.InputPort.Offset)
// 			count := len(expectedBlogs)
// 			assert.Equal(t, 0, count)
//
// 			i.GetUCHandler().BlogGetAll(useCase)
//
// 			if mutation.ShouldPass {
// 				assert.NoError(t, nil)
// 				return
// 			}
//
// 			assert.NoError(t, nil)
// 			assert.Error(t, form.Present())
// 		})
// 	}
// }
