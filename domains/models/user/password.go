package user

import (
	"go-api/domains/models/validator"
)

type Password struct {
	// [Passwordを表現する値オブジェクト]
	// バリデーションルールは以下
	// - 空ではないこと
	// - 6文字以上であること
	Value string `validate:"required,gte=6" ja:"パスワード"`
}

func NewPassword(value string) (Password, error) {
	password := Password{Value: value}
	err := validator.Validate(password)
	if err != nil {
		return password, err
	}
	return password, nil
}

func UpdatePassword(input *string) (*Password, error) {
	password := Password{Value: *input}
	err := validator.Validate(&password)
	if err != nil {
		return &password, err
	}
	return &password, nil
}
