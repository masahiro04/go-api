package blog

type Body struct {
	// [Blogの説明を表現する値オブジェクト]
	// バリデーションルールは以下
	// - 空ではないこと
	// - 100文字以下であること
	value string `validate:"required,max=100" ja:"内容"`
}

func NewBody(value string) (Body, error) {
	// TODO(okubo): validationがうまく起動しないので、修正する
	// fmt.Println("newBody")
	// err := validator.Validate(value)
	// if err != nil {
	// 	return Body{value: value}, err
	// }
	return Body{value: value}, nil
}

func UpdateBody(input *string) (*Body, error) {
	// fmt.Println("updateBody")
	// err := validator.Validate(&input)
	// if err != nil {
	// 	return &Body{value: *input}, err
	// }
	return &Body{value: *input}, nil
}

func (body Body) Value() string {
	return body.value
}
