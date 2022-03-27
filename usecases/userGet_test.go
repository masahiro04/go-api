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

func TestUserGetSuccess(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	user := factories.User()

	i := mock.NewMockedInteractor(mockCtrl)
	i.UserDao.EXPECT().GetById(user.ID.Value).Return(&user, nil).Times(1)

	ginContext, _ := gin.CreateTestContext(httptest.NewRecorder())
	pre := presenters.New(ginContext)
	form := json.NewPresenter(pre)
	useCase := uc.GetUserUseCase{
		OutputPort: form,
		InputPort:  uc.GetUserParams{ID: user.ID.Value},
	}

	i.GetUCHandler().UserGet(useCase)

	assert.NoError(t, nil)
	assert.NoError(t, form.Present())
}

func TestUserGetFails(t *testing.T) {
	user := factories.User()

	mutations := map[string]mock.Tester{
		"shouldPass": {
			Calls: func(i *mock.Interactor) { // change nothing
			},
			ShouldPass: true},
		"error return on UserDao.GetById": {
			Calls: func(i *mock.Interactor) {
				i.UserDao.EXPECT().GetById(user.ID.Value).Return(nil, errors.New(""))
			}},
		"nil nil return on UserDao.GetById": {
			Calls: func(i *mock.Interactor) {
				i.UserDao.EXPECT().GetById(user.ID.Value).Return(nil, nil)
			}},
	}

	validCalls := func(i *mock.Interactor) {
		i.UserDao.EXPECT().GetById(user.ID.Value).Return(&user, nil).AnyTimes()
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
			useCase := uc.GetUserUseCase{
				OutputPort: form,
				InputPort:  uc.GetUserParams{ID: user.ID.Value},
			}

			i.GetUCHandler().UserGet(useCase)

			if mutation.ShouldPass {
				assert.NoError(t, nil)
				return
			}

			assert.NoError(t, nil)
			assert.Error(t, form.Present())
		})
	}
}
