package infrastructure

import (
	"context"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
)

func NewFirebaseAuthClient() *auth.Client {
	app, err := firebase.NewApp(context.Background(), nil)
	if err != nil {
		panic(err)
	}
	client, err := app.Auth(context.Background())
	if err != nil {
		panic(err)
	}
	return client
}
