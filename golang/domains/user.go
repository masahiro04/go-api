package domains

import (
	"clean_architecture/golang/domains/blog"
	"time"
)

// TODO: privateに扱うために、小文字に変更する
type User struct {
	ID        blog.ID
	Name      blog.Title
	Email     blog.Body
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

// func NewUser(title blog.Title, body blog.Body) Blog {
// 	return Blog{
// 		Title: title,
// 		Body:  body,
// 	}
// }
//
// // repositoryやfactory経由の生成において使用する関数
// // 生成時のバリデーションをしないことに注意
// func BuildBlog(id blog.ID, title blog.Title, body blog.Body, createdAt time.Time, updatedAt time.Time) Blog {
// 	return Blog{
// 		ID:        id,
// 		Title:     title,
// 		Body:      body,
// 		CreatedAt: createdAt,
// 	}
// }
//
// func (b *Blog) UpdateTitle(title blog.Title) *Blog {
// 	b.Title = title
// 	return b
// }
//
// func (b *Blog) UpdateBody(body blog.Body) *Blog {
// 	b.Body = body
// 	return b
// }
