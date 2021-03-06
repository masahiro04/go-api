package user

import (
	"go-api/domains/models/validator"
)

type Email struct {
	// [Emailを表現する値オブジェクト]
	// バリデーションルールは以下
	// - 空ではないこと
	// - Emailであること
	Value string `validate:"required,email" ja:"メールアドレス"`
}

func NewEmail(value string) (Email, error) {
	email := Email{Value: value}
	err := validator.Validate(email)
	if err != nil {
		return email, err
	}
	return email, nil
}

func UpdateEmail(input *string) (*Email, error) {
	email := Email{Value: *input}
	err := validator.Validate(&email)
	if err != nil {
		return &email, err
	}
	return &email, nil
}
