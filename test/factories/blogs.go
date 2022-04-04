package factories

import (
	"go-api/domains/models"
	"go-api/domains/models/blog"

	// "go-api/presenter/json"
	"strconv"
	"time"
)

func Blog() models.Blog {
	id, _ := blog.NewId(1)
	title, _ := blog.NewTitle("タイトル")
	body, _ := blog.NewBody("内容")
	return models.BuildBlog(id, title, body, time.Time{}, time.Time{})
}

func BlogWithID(newID int) models.Blog {
	id, _ := blog.NewId(newID)
	title, _ := blog.NewTitle("タイトル")
	body, _ := blog.NewBody("内容")
	return models.BuildBlog(id, title, body, time.Time{}, time.Time{})
}

func NewBlog() models.Blog {
	title, _ := blog.NewTitle("タイトル")
	body, _ := blog.NewBody("内容")
	return models.NewBlog(title, body)
}

func Blogs(length int) models.Blogs {
	var blogs []models.Blog
	n := 1
	for n < length {
		id, _ := blog.NewId(n)
		title, _ := blog.NewTitle("タイトル" + strconv.Itoa(n))
		body, _ := blog.NewBody("内容" + strconv.Itoa(n))
		blogs = append(blogs, models.BuildBlog(id, title, body, time.Time{}, time.Time{}))
		n++
	}
	return models.NewBlogs(blogs)
}
