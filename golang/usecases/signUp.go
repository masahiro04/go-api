package uc

type SignUpParams struct {
	User SignUpUserParams
}

type SignUpUserParams struct {
	Email    string
	Password string
}

type SignUpCompanyParams struct {
	Name string
}

// func (i interactor) SignUp(params SignUpParams, presenter Presenter) {
// 	var err error
// 	user := &domains.User{
// 		Email:    params.User.Email,
// 		Password: &params.User.Password,
// 	}
//
// 	err = i.dbTransaction.WithTx(func(tx *gorm.DB) error {
// 		uuId, err := i.firebaseHandler.CreateUser(*user)
// 		if err != nil {
// 			if rollbackErr := tx.Rollback(); rollbackErr != nil {
// 				i.logger.Log(err)
// 			}
// 			return err
// 		}
//
// 		company, err = i.companyRW.CreateTx(*company, tx)
// 		if err != nil {
// 			if rollbackErr := tx.Rollback(); rollbackErr != nil {
// 				i.logger.Log(err)
// 				if err = i.firebaseHandler.DeleteUser(user.UuId); err != nil {
// 					i.logger.Log(err)
// 				}
// 			}
//
// 			return err
// 		}
//
// 		user.UuId = *uuId
// 		user.CompanyId = company.ID
//
// 		_, err = i.userRW.CreateTx(*user, tx)
// 		if err != nil {
// 			if rollbackErr := tx.Rollback(); rollbackErr != nil {
// 				i.logger.Log(err)
// 				if err = i.firebaseHandler.DeleteUser(user.UuId); err != nil {
// 					i.logger.Log(err)
// 				}
// 			}
//
// 			return err
// 		}
//
// 		return nil
// 	})
//
// 	if err != nil {
// 		presenter.Raise(domain.UnprocessableEntity, err)
// 		return
// 	}
//
// 	link, err := i.firebaseHandler.EmailSignInLink(user.Email)
// 	if err != nil {
// 		presenter.Raise(domain.UnprocessableEntity, err)
// 		return
// 	}
//
// 	message, err := i.mailPresenter.SendMessage(
// 		user.Email,
// 		domain.MessageTitleForSignUp,
// 		domain.MessageBodyForSignUp,
// 		&domain.MessageTemplateData{Link: link})
// 	if err != nil {
// 		presenter.Raise(domain.UnprocessableEntity, err)
// 		return
// 	}
//
// 	err = i.awsSqs.SendMessage(*message)
// 	if err != nil {
// 		presenter.Raise(domain.UnprocessableEntity, err)
// 		return
// 	}
//
// 	presenter.CreateSignUp(user)
// }
