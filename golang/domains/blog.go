package domains

import (
	"clean_architecture/golang/domains/blog"
	"time"
)

type Blog struct {
	ID        blog.ID
	Title     blog.Title
	Body      blog.Body
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

func NewBlog(title blog.Title, body blog.Body) Blog {
	return Blog{
		Title: title,
		Body:  body,
	}
}

// repositoryやfactory経由の生成において使用する関数
// 生成時のバリデーションをしないことに注意
func BuildBlog(id blog.ID, title blog.Title, body blog.Body) Blog {
	return Blog{
		ID:    id,
		Title: title,
		Body:  body,
	}
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
