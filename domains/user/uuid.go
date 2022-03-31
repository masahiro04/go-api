package user

import (
	"go-api/domains/validator"
)

type UUID struct {
	// [UUIDを表現する値オブジェクト]
	// バリデーションルールは以下
	// - 空ではないこと
	// - Emailであること
	Value string `validate:"required" ja:"UUID"`
}

func NewUUID(value string) (UUID, error) {
	uuid := UUID{Value: value}
	err := validator.Validate(uuid)
	if err != nil {
		return uuid, err
	}
	return uuid, nil
}

func UpdateUUID(input *string) (*UUID, error) {
	uuid := UUID{Value: *input}
	err := validator.Validate(&uuid)
	if err != nil {
		return &uuid, err
	}
	return &uuid, nil
}
