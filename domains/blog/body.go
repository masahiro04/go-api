package blog

import (
	"go-api/domains/validator"
)

type Body struct {
	// [Bodyを表現する値オブジェクト]
	// バリデーションルールは以下
	// - 空ではないこと
	Value string `validate:"required" ja:"内容"`
}

func NewBody(value string) (Body, error) {
	body := Body{Value: value}
	err := validator.Validate(body)
	if err != nil {
		return body, err
	}
	return body, nil
}

func UpdateBody(input *string) (*Body, error) {
	body := Body{Value: *input}
	err := validator.Validate(&body)
	if err != nil {
		return &body, err
	}
	return &body, nil
}
