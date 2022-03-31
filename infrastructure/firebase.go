package infrastructure

import (
	"context"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

func NewFirebaseAuthClient() *auth.Client {

	// TODO(okubo): 指定を変更する
	options := option.WithCredentialsFile("./serviceAccountKey.json")
	app, err := firebase.NewApp(context.Background(), nil, options)
	if err != nil {
		panic(err)
	}
	client, err := app.Auth(context.Background())
	if err != nil {
		panic(err)
	}
	return client
}
