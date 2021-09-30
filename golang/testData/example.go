package testData

import (
	"clean_architecture/golang/domain"
)

func Blog() domain.Blog {
	return domain.Blog{
		ID:    1,
		Title: "タイトル",
		Body:  "Body",
	}
}
