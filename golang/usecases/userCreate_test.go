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

func TestUserCreateSuccess(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	user := testData.User()

	// Mock
	i := mock.NewMockedInteractor(mockCtrl)
	i.UserDao.EXPECT().Create(gomock.Any()).Return(&user, nil).Times(1)

	// UseCase
	ginContext, _ := gin.CreateTestContext(httptest.NewRecorder())
	pre := presenters.New(ginContext)
	form := json.NewPresenter(pre)
	useCase := uc.CreateUserUseCase{
		OutputPort: form,
		InputPort: uc.CreateUserParams{
			Name:     user.Name.Value,
			Email:    user.Email.Value,
			Password: "hogehgoe",
		},
	}

	i.GetUCHandler().UserCreate(useCase)

	assert.NoError(t, form.Present())
}

func TestUserCreateFails(t *testing.T) {
	user := testData.User()

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
				i.UserDao.EXPECT().Create(gomock.Any()).Return(nil, errors.New("")).Times(1)
			}},
	}

	validCalls := func(i *mock.Interactor) {
		i.UserDao.EXPECT().Create(gomock.Any()).Return(&user, nil).AnyTimes()
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
			useCase := uc.CreateUserUseCase{
				OutputPort: form,
				InputPort: uc.CreateUserParams{
					Name:     user.Name.Value,
					Email:    user.Email.Value,
					Password: "hogehgeo",
				},
			}

			i.GetUCHandler().UserCreate(useCase)

			if mutation.ShouldPass {
				assert.NoError(t, nil)
				return
			}

			assert.Error(t, form.Present())
		})
	}
}
