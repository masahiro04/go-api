package user

import (
	"clean_architecture/golang/domains/validator"
)

type Name struct {
	// [Blogの説明を表現する値オブジェクト]
	// バリデーションルールは以下
	// - 空ではないこと
	// - 100文字以下であること
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
