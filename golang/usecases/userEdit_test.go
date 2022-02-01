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

func TestUserEditSuccess(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	user := testData.User()

	i := mock.NewMockedInteractor(mockCtrl)
	i.UserDao.EXPECT().GetById(user.ID.Value).Return(&user, nil).Times(1)
	i.UserDao.EXPECT().Update(user.ID.Value, user).Return(&user, nil).Times(1)

	ginContext, _ := gin.CreateTestContext(httptest.NewRecorder())
	pre := presenters.New(ginContext)
	form := json.NewPresenter(pre)

	useCase := uc.EditUserUseCase{
		OutputPort: form,
		InputPort: uc.EditUserParams{
			ID:    user.ID.Value,
			Name:  user.Name.Value,
			Email: user.Email.Value,
		},
	}

	i.GetUCHandler().UserEdit(useCase)

	assert.NoError(t, nil)
	assert.NoError(t, form.Present())
}

func TestUserEditFails(t *testing.T) {
	user := testData.User()

	ginContext, _ := gin.CreateTestContext(httptest.NewRecorder())
	pre := presenters.New(ginContext)
	form := json.NewPresenter(pre)

	useCase := uc.EditUserUseCase{
		OutputPort: form,
		InputPort: uc.EditUserParams{
			ID:    user.ID.Value,
			Name:  user.Name.Value,
			Email: user.Email.Value,
		},
	}
	mutations := map[string]mock.Tester{
		"error return on UserDao.GetById": {
			Calls: func(i *mock.Interactor) {
				i.UserDao.EXPECT().GetById(user.ID.Value).Return(nil, errors.New(""))
			}},
		"nil, nil return on UserDao.GetById": {
			Calls: func(i *mock.Interactor) {
				i.UserDao.EXPECT().GetById(user.ID.Value).Return(nil, nil)
			}},

		"error returns when User.title is blank": {
			Calls: func(i *mock.Interactor) {
				useCase.InputPort.Name = ""
				i.UserDao.EXPECT().GetById(user.ID.Value).Return(&user, nil).AnyTimes()
			}},
		"error returns when User.body is blank": {
			Calls: func(i *mock.Interactor) {
				useCase.InputPort.ID = user.ID.Value
				useCase.InputPort.Name = user.Name.Value
				useCase.InputPort.Email = ""
				i.UserDao.EXPECT().GetById(user.ID.Value).Return(&user, nil).AnyTimes()
			}},
		"error returns when User cannot update": {
			Calls: func(i *mock.Interactor) {
				useCase.InputPort.ID = user.ID.Value
				useCase.InputPort.Name = user.Name.Value
				useCase.InputPort.Email = user.Email.Value

				i.UserDao.EXPECT().GetById(user.ID.Value).Return(&user, nil).AnyTimes()
				i.UserDao.EXPECT().Update(gomock.Any(), gomock.Any()).Return(nil, errors.New("")).AnyTimes()
			}},
	}

	// validCalls := func(i *mock.Interactor) {
	// 	i.Logger.EXPECT().Log(gomock.Any()).AnyTimes()
	// 	i.UserDao.EXPECT().GetById(user.ID.Value).Return(&user, nil).AnyTimes()
	// 	i.UserDao.EXPECT().Update(User.ID.Value, User).Return(&User, nil).AnyTimes()
	// 	i.Validator.EXPECT().Validate(gomock.Any()).Return(nil).AnyTimes()
	// }

	for testName, mutation := range mutations {
		t.Run(testName, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			i := mock.NewMockedInteractor(mockCtrl)
			mutation.Calls(&i)
			// validCalls(&i)

			i.GetUCHandler().UserEdit(useCase)

			if mutation.ShouldPass {
				assert.NoError(t, nil)
				return
			}

			assert.NoError(t, nil)
			assert.Error(t, form.Present())
		})
	}
}
