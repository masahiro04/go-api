package domain

import (
	"time"
)

type Blog struct {
	ID        int
	Title     string `validate:"required,max=255" ja:"タイトル"`
	Body      string `validate:"required,max=255" ja:"内容"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type BlogCollection []*Blog
type BlogUpdatableProperty int

const (
	BlogTitle BlogUpdatableProperty = iota
	BlogBody
)

func UpdateBlog(initial *Blog, opts ...func(fields *Blog)) {
	for _, v := range opts {
		v(initial)
	}
}

func SetBlogTitle(input *string) func(fields *Blog) {
	return func(initial *Blog) {
		if input != nil {
			initial.Title = *input
		}
	}
}

func SetBlogBody(input *string) func(fields *Blog) {
	return func(initial *Blog) {
		if input != nil {
			initial.Body = *input
		}
	}
}

func (blogs BlogCollection) ApplyLimitAndOffset(limit, offset int) BlogCollection {
	if limit <= 0 {
		return []*Blog{}
	}

	blogsSize := len(blogs)
	min := offset
	if min < 0 {
		min = 0
	}

	if min > blogsSize {
		return []*Blog{}
	}

	max := min + limit
	if max > blogsSize {
		max = blogsSize
	}

	return blogs[min:max]
}
