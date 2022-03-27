package user

import (
	"go-api/domains/validator"
)

type Name struct {
	// [Nameを表現する値オブジェクト]
	// バリデーションルールは以下
	// - 空ではないこと
	Value string `validate:"required" ja:"名前"`
}

func NewName(value string) (Name, error) {
	name := Name{Value: value}
	err := validator.Validate(name)
	if err != nil {
		return name, err
	}
	return name, nil
}

func UpdateName(input *string) (*Name, error) {
	name := Name{Value: *input}
	err := validator.Validate(&name)

	if err != nil {
		return &name, err
	}
	return &name, nil
}
