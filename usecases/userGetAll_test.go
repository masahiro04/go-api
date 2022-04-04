package usecases_test

// func TestGetAllUser_Success(t *testing.T) {
// 	var _users = factories.Users(5)
// 	t.Run("most obvious", func(t *testing.T) {
// 		mockCtrl := gomock.NewController(t)
// 		defer mockCtrl.Finish()
//
// 		i := mock.NewMockedInteractor(mockCtrl)
// 		i.UserDao.EXPECT().GetAll().Return(&_users, nil).Times(1)
//
// 		// UseCase
// 		ginContext, _ := gin.CreateTestContext(httptest.NewRecorder())
// 		pre := presenters.New(ginContext)
// 		form := json.NewPresenter(pre)
// 		useCase := uc.GetUsersUseCase{
// 			OutputPort: form,
// 			InputPort:  uc.GetUsersParams{Limit: 1, Offset: 1},
// 		}
//
// 		expectedUsers := _users.ApplyLimitAndOffset(useCase.InputPort.Limit, useCase.InputPort.Offset)
//
// 		assert.Equal(t, 1, len(expectedUsers))
// 		assert.Equal(t, "名前2", expectedUsers[0].Name.Value)
//
// 		i.GetUCHandler().UserGetAll(useCase)
//
// 		assert.NoError(t, nil)
// 		assert.NoError(t, form.Present())
// 	})
// }
//
// func TestGetAllUser_Fail(t *testing.T) {
// 	var _users = factories.Users(5)
// 	mutations := map[string]mock.Tester{
// 		"shouldPass": {
// 			Calls: func(i *mock.Interactor) { // change nothing
// 			},
// 			ShouldPass: true},
// 		"error return on uRW.GetFiltered": {
// 			Calls: func(i *mock.Interactor) {
// 				i.UserDao.EXPECT().GetAll().Return(nil, errors.New(""))
// 			}},
// 	}
//
// 	validCalls := func(i *mock.Interactor) {
// 		i.UserDao.EXPECT().GetAll().Return(&_users, nil).AnyTimes()
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
// 			pre := presenters.New(ginContext)
// 			form := json.NewPresenter(pre)
// 			useCase := uc.GetUsersUseCase{
// 				OutputPort: form,
// 				InputPort:  uc.GetUsersParams{Limit: 1, Offset: 4},
// 			}
//
// 			expectedUsers := _users.ApplyLimitAndOffset(useCase.InputPort.Limit, useCase.InputPort.Offset)
// 			count := len(expectedUsers)
// 			assert.Equal(t, 0, count)
//
// 			i.GetUCHandler().UserGetAll(useCase)
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
