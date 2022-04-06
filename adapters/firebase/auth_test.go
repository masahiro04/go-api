package firebase_test

import (
	"context"
	"testing"
	"time"

	firebase "go-api/adapters/firebase"
	factories "go-api/test/factories"

	"firebase.google.com/go/auth"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

var mockedUID = "MockedUID"

const mockedToken = "MockedToken"
const mockedEmail = "test@example.com"
const mockedLink = "http://localhost:3000"

type fakeAuthClient struct{}

func (f *fakeAuthClient) VerifyIDToken(context.Context, string) (*auth.Token, error) {
	return &auth.Token{
		UID: mockedUID,
	}, nil
}

func (f *fakeAuthClient) GetUser(context.Context, string) (*auth.UserRecord, error) {
	return &auth.UserRecord{
		UserInfo: &auth.UserInfo{
			Email: mockedEmail},
	}, nil
}

func (f *fakeAuthClient) CreateUser(context.Context, *auth.UserToCreate) (*auth.UserRecord, error) {
	return &auth.UserRecord{
		UserInfo: &auth.UserInfo{
			UID:   mockedUID,
			Email: mockedEmail},
	}, nil
}

// func (f *fakeAuthClient) UpdateUser(context.Context, string, *auth.UserToUpdate) (*auth.UserRecord, error) {
// 	return &auth.UserRecord{
// 		UserInfo: &auth.UserInfo{
// 			UID:   mockedUID,
// 			Email: mockedEmail},
// 	}, nil
// }

func (f *fakeAuthClient) DeleteUser(ctx context.Context, uid string) error {
	return nil
}

// func (f *fakeAuthClient) SetCustomUserClaims(ctx context.Context, uid string, customClaims map[string]interface{}) error {
// 	return nil
// }

func (f *fakeAuthClient) EmailVerificationLinkWithSettings(
	ctx context.Context, email string, settings *auth.ActionCodeSettings) (string, error) {
	return mockedLink, nil
}

func (f *fakeAuthClient) EmailSignInLink(
	ctx context.Context, email string, settings *auth.ActionCodeSettings) (string, error) {
	return mockedLink, nil
}

func (f *fakeAuthClient) SessionCookie(ctx context.Context, idToken string, expiresIn time.Duration) (string, error) {
	return mockedLink, nil
}

func (f *fakeAuthClient) VerifySessionCookieAndCheckRevoked(ctx context.Context, sessionCookie string) (*auth.Token, error) {
	return &auth.Token{UID: mockedUID}, nil
}

func (f *fakeAuthClient) RevokeRefreshTokens(ctx context.Context, uid string) error {
	return nil
}

func TestNew_happyCase(t *testing.T) {
	Client := firebase.New(&fakeAuthClient{})

	assert.NotEmpty(t, Client)
}

// func TestGetUserUuid_happyCase(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()
//
// 	Client := firebase.New(&fakeAuthClient{})
// 	_, err := Client.GetUser(mockedToken)
// 	assert.Empty(t, err)
//
// 	// assert.Equal(t, &mockedUID, firebaseUuId)
// }

func TestCreateUser_happyCase(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	user := factories.User()

	Client := firebase.New(&fakeAuthClient{})
	firebaseUuId, err := Client.CreateUser(user)

	assert.Empty(t, err)
	assert.Equal(t, &mockedUID, firebaseUuId)
}

// func TestUpdateUser_happyCase(t *testing.T) {
// 	var mockedEmail = mockedEmail
// 	var mockedPassword = "password"
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()
//
// 	user := factories.User()
//
// 	role := user.Role.String()
// 	Client := firebase.New(&fakeAuthClient{})
// 	err := Client.UpdateUser(mockedUID, &domain.UserUpdatableProperty{
// 		Email:     &mockedEmail,
// 		Role:      &role,
// 		Password:  &mockedPassword,
// 		LastName:  &mockedEmail,
// 		FirstName: &mockedEmail,
// 	})
// 	assert.Empty(t, err)
// }

// func TestActivateUser_happyCase(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()
//
// 	Client := firebase.New(&fakeAuthClient{})
// 	err := Client.ActivateUser(mockedUID)
// 	assert.Empty(t, err)
// }
// func TestDisableUser_happyCase(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()
//
// 	Client := firebase.New(&fakeAuthClient{})
// 	err := Client.DisableUser(mockedUID)
// 	assert.Empty(t, err)
// }

// func TestDeleteUser_happyCase(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()
//
// 	Client := firebase.New(&fakeAuthClient{})
// 	err := Client.DeleteUser(mockedUID)
// 	assert.Empty(t, err)
// }

func TestEmailVerificationLinkWithSettings_happyCase(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	var mockedLink = mockedLink

	Client := firebase.New(&fakeAuthClient{})
	link, err := Client.EmailVerificationLinkWithSettings(mockedEmail)
	assert.Empty(t, err)
	assert.Equal(t, link, &mockedLink)
}

func TestEmailSignInLink_happyCase(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	var mockedLink = mockedLink

	Client := firebase.New(&fakeAuthClient{})
	link, err := Client.EmailSignInLink(mockedEmail)
	assert.Empty(t, err)
	assert.Equal(t, link, &mockedLink)
}

func TestSessionCookie_happyCase(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	var mockedLink = mockedLink

	Client := firebase.New(&fakeAuthClient{})
	expiresIn := time.Hour * 24 * 5
	link, err := Client.SessionCookie(mockedLink, expiresIn)
	assert.Empty(t, err)
	assert.Equal(t, link, &mockedLink)
}

func TestVerifySessionCookieAndCheckRevoked_happyCase(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	var mockedLink = mockedLink

	Client := firebase.New(&fakeAuthClient{})
	uuid, err := Client.VerifySessionCookieAndCheckRevoked(mockedLink)
	assert.Empty(t, err)
	assert.Equal(t, uuid, &mockedUID)
}

func TestRevokeRefreshTokens_happyCase(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	Client := firebase.New(&fakeAuthClient{})
	err := Client.RevokeRefreshTokens(mockedUID)
	assert.Empty(t, err)
}
