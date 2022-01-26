package blog

import (
	"fmt"
)

type Title struct {
	// [Blogの説明を表現する値オブジェクト]
	// バリデーションルールは以下
	// - 空ではないこと
	// - 100文字以下であること
	value string `validate:"required,max=100" ja:"タイトル"`
}

func NewTitle(value string) (Title, error) {
	fmt.Println("newTitle")
	fmt.Println(value)

	// err := validator.Validate(value)
	// if err != nil {
	// 	return Title{value: value}, err
	// }
	return Title{value: value}, nil
}

func UpdateTitle(input *string) (*Title, error) {
	// fmt.Println("updateTitle")
	// err := validator.Validate(input)
	// if err != nil {
	// 	return &Title{value: *input}, err
	// }
	return &Title{value: *input}, nil
}

func (title Title) Value() string {
	return title.value
}
