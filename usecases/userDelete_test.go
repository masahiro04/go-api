package uc_test

import (
	"errors"
	"net/http/httptest"
	"testing"

	"go-api/adapters/presenters"
	"go-api/adapters/presenters/json"

	"github.com/gin-gonic/gin"

	mock "go-api/adapters/uc.mock"
	"go-api/testData"

	uc "go-api/usecases"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUserDeleteSuccess(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	user := testData.User()

	i := mock.NewMockedInteractor(mockCtrl)
	i.UserDao.EXPECT().Delete(user.ID.Value).Return(nil).Times(1)

	ginContext, _ := gin.CreateTestContext(httptest.NewRecorder())
	pre := presenters.New(ginContext)
	form := json.NewPresenter(pre)
	useCase := uc.DeleteUserUseCase{
		OutputPort: form,
		InputPort: uc.DeleteUserParams{
			ID: user.ID.Value,
		},
	}

	i.GetUCHandler().UserDelete(useCase)

	assert.NoError(t, nil)
	assert.NoError(t, form.Present())
}

func TestUserDeleteFails(t *testing.T) {
	user := testData.User()

	mutations := map[string]mock.Tester{
		"shouldPass": {
			Calls: func(i *mock.Interactor) { // change nothing
			},
			ShouldPass: true},
		"failed to save the company": {
			Calls: func(i *mock.Interactor) {
				i.UserDao.EXPECT().Delete(user.ID.Value).Return(errors.New(""))
			}},
	}

	validCalls := func(i *mock.Interactor) {
		i.UserDao.EXPECT().Delete(user.ID.Value).Return(nil).AnyTimes()
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
			useCase := uc.DeleteUserUseCase{
				OutputPort: form,
				InputPort: uc.DeleteUserParams{
					ID: user.ID.Value,
				},
			}

			i.GetUCHandler().UserDelete(useCase)

			if mutation.ShouldPass {
				assert.NoError(t, form.Present())
				return
			}

			assert.Error(t, form.Present())
		})
	}
}
