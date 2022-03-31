package firebase

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/spf13/viper"

	"go-api/domains"
	uc "go-api/usecases"

	"firebase.google.com/go/auth"
)

type firebaseClient interface {
	VerifyIDToken(context.Context, string) (*auth.Token, error)
	// GetUser(ctx context.Context, uid string) (*auth.UserRecord, error)
	CreateUser(ctx context.Context, params *auth.UserToCreate) (*auth.UserRecord, error)
	// UpdateUser(ctx context.Context, uid string, user *auth.UserToUpdate) (ur *auth.UserRecord, err error)
	// DeleteUser(ctx context.Context, uid string) error
	// SetCustomUserClaims(ctx context.Context, uid string, customClaims map[string]interface{}) error
	EmailVerificationLinkWithSettings(ctx context.Context, email string, settings *auth.ActionCodeSettings) (string, error)
	EmailSignInLink(ctx context.Context, email string, settings *auth.ActionCodeSettings) (string, error)
	SessionCookie(ctx context.Context, idToken string, expiresIn time.Duration) (string, error)
	VerifySessionCookieAndCheckRevoked(ctx context.Context, sessionCookie string) (*auth.Token, error)
	RevokeRefreshTokens(ctx context.Context, uid string) error
}

type tokenHandler struct {
	client firebaseClient
}

func New(client firebaseClient) uc.FirebaseHandler {
	return tokenHandler{client}
}

func (tH tokenHandler) VerifyIDToken(idToken string) (token *auth.Token, err error) {
	token, err = tH.client.VerifyIDToken(context.Background(), idToken)
	if err != nil {
		return nil, err
	}
	return token, nil
}

// func (tH tokenHandler) GetUser(uuId string) (user *domains.User, err error) {
// 	firebaseUser, err := tH.client.GetUser(context.Background(), uuId)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	user = &domain.User{
// 		EmailVerified: &firebaseUser.EmailVerified,
// 	}
//
// 	return user, nil
// }

func (tH tokenHandler) CreateUser(user domains.User) (uuId *string, err error) {
	params := (&auth.UserToCreate{}).
		Email(user.Email.Value).
		// EmailSignInLinkを実現するにはこちらが必須
		EmailVerified(true).
		Disabled(false)

	if &user.Password.Value != nil {
		params.Password(*&user.Password.Value)
	}

	firebaseUser, err := tH.client.CreateUser(context.Background(), params)
	if err != nil {
		return nil, err
	}

	return &firebaseUser.UID, nil
}

// func (tH tokenHandler) UpdateUser(uuId string, updateParams *domain.UserUpdatableProperty) error {
// 	params := &auth.UserToUpdate{}
// 	params.Disabled(false) //NOTO:すべてnilだと落ちるのでここで対応している。
//
// 	if updateParams.Email != nil {
// 		params.Email(*updateParams.Email)
// 		params.EmailVerified(false)
// 	}
//
// 	if updateParams.Password != nil {
// 		params.Password(*updateParams.Password)
// 	}
//
// 	if updateParams.LastName != nil || updateParams.FirstName != nil {
// 		params.DisplayName(*updateParams.LastName + " " + *updateParams.FirstName)
// 	}
//
// 	_, err := tH.client.UpdateUser(context.Background(), uuId, params)
// 	if err != nil {
// 		return err
// 	}
//
// 	return nil
// }

// 有効にする場合
// func (tH tokenHandler) ActivateUser(uuId string) error {
// 	params := &auth.UserToUpdate{}
// 	params.Disabled(false)
//
// 	_, err := tH.client.UpdateUser(context.Background(), uuId, params)
// 	if err != nil {
// 		return err
// 	}
//
// 	return nil
// }

// 無効にする場合
// func (tH tokenHandler) DisableUser(uuId string) error {
// 	params := &auth.UserToUpdate{}
// 	params.Disabled(true)
//
// 	_, err := tH.client.UpdateUser(context.Background(), uuId, params)
// 	if err != nil {
// 		return err
// 	}
//
// 	return nil
// }

// 完全に削除する場合
// func (tH tokenHandler) DeleteUser(uuId string) error {
// 	err := tH.client.DeleteUser(context.Background(), uuId)
// 	if err != nil {
// 		return err
// 	}
//
// 	return nil
// }

func (tH tokenHandler) EmailVerificationLinkWithSettings(email string) (*string, error) {
	actionCodeSettings := &auth.ActionCodeSettings{URL: fmt.Sprintf("%s/login", viper.GetString("server.fronthost"))}
	link, err := tH.client.EmailVerificationLinkWithSettings(
		context.Background(), email, actionCodeSettings)
	if err != nil {
		return nil, err
	}

	return &link, err
}

func (tH tokenHandler) EmailSignInLink(email string) (*string, error) {
	actionCodeSettings := &auth.ActionCodeSettings{URL: fmt.Sprintf("%s/login", viper.GetString("server.fronthost"))}
	link, err := tH.client.EmailVerificationLinkWithSettings(context.Background(), email, actionCodeSettings)
	if err != nil {
		log.Fatalf("error generating email link: %v\n", err)
	}

	return &link, err
}

func (tH tokenHandler) SessionCookie(idToken string, expiresIn time.Duration) (*string, error) {
	cookie, err := tH.client.SessionCookie(context.Background(), idToken, expiresIn)
	if err != nil {
		return nil, err
	}

	return &cookie, nil
}

func (tH tokenHandler) VerifySessionCookieAndCheckRevoked(sessionCookie string) (uuid *string, err error) {
	decoded, err := tH.client.VerifySessionCookieAndCheckRevoked(context.Background(), sessionCookie)
	if err != nil {
		return nil, err
	}

	return &decoded.UID, nil
}

func (tH tokenHandler) RevokeRefreshTokens(uuId string) error {
	err := tH.client.RevokeRefreshTokens(context.Background(), uuId)
	if err != nil {
		return err
	}

	return nil
}
