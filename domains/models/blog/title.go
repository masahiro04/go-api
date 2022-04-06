package blog

import "go-api/domains/models/validator"

type Title struct {
	// [Titleを表現する値オブジェクト]
	// バリデーションルールは以下
	// - 空ではないこと
	// - 100文字以下であること
	Value string `validate:"required,max=100" ja:"タイトル"`
}

func NewTitle(value string) (Title, error) {
	title := Title{Value: value}
	err := validator.Validate(title)
	if err != nil {
		return title, err
	}
	return title, nil
}

func UpdateTitle(input *string) (*Title, error) {
	title := Title{Value: *input}
	err := validator.Validate(title)
	if err != nil {
		return &title, err
	}
	return &title, nil
}
