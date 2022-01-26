package testData

import (
	"clean_architecture/golang/domains"
	blog "clean_architecture/golang/domains/blog"
	"strconv"
)

func Blog() domains.Blog {
	id, _ := blog.NewId(1)
	title, _ := blog.NewTitle("タイトル")
	body, _ := blog.NewBody("内容")
	return domains.BuildBlog(id, title, body)
}

func PointeredBlog() *domains.Blog {
	id, _ := blog.NewId(1)
	title, _ := blog.NewTitle("タイトル")
	body, _ := blog.NewBody("内容")
	newBlog := domains.BuildBlog(id, title, body)
	return &newBlog
}

func Blogs(length int) []domains.Blog {
	var blogs []domains.Blog
	n := 1
	for n < length {
		id, _ := blog.NewId(n)
		title, _ := blog.NewTitle("タイトル" + strconv.Itoa(n))
		body, _ := blog.NewBody("内容" + strconv.Itoa(n))
		blogs = append(blogs, domains.BuildBlog(id, title, body))
		n++
	}
	return blogs
}
func PointeredBlogs(length int) []*domains.Blog {
	var blogs []*domains.Blog
	n := 1
	for n < length {
		id, _ := blog.NewId(n)
		title, _ := blog.NewTitle("タイトル" + strconv.Itoa(n))
		body, _ := blog.NewBody("内容" + strconv.Itoa(n))
		newBlog := domains.BuildBlog(id, title, body)
		blogs = append(blogs, &newBlog)
		n++
	}
	return blogs
}
