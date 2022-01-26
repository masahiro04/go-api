package domains

import (
	"clean_architecture/golang/domains/blog"
	"time"
)

// TODO: privateに扱うために、小文字に変更する
type Blog struct {
	id        blog.ID
	title     blog.Title
	body      blog.Body
	createdAt time.Time
	updatedAt time.Time
	deletedAt time.Time
}

func NewBlog(title blog.Title, body blog.Body) Blog {
	return Blog{
		title: title,
		body:  body,
	}
}

// repositoryやfactory経由の生成において使用する関数
// 生成時のバリデーションをしないことに注意
func BuildBlog(id blog.ID, title blog.Title, body blog.Body) Blog {
	return Blog{
		id:    id,
		title: title,
		body:  body,
	}
}

func (b *Blog) ID() blog.ID {
	return b.id
}

func (b *Blog) Title() blog.Title {
	return b.title
}

func (b *Blog) Body() blog.Body {
	return b.body
}

func (b *Blog) UpdateTitle(title blog.Title) *Blog {
	b.title = title
	return b
}

func (b *Blog) UpdateBody(body blog.Body) *Blog {
	b.body = body
	return b
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
