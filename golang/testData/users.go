package testData

import (
	"clean_architecture/golang/domains"
	"clean_architecture/golang/domains/user"
	"time"
)

func User() domains.User {
	id, _ := user.NewId(1)
	name, _ := user.NewName("大久保")
	email, _ := user.NewEmail("test@example.com")
	return domains.BuildUser(id, name, email, time.Time{}, time.Time{})
}

// func Blogs(length int) domains.Blogs {
// 	var blogs []domains.Blog
// 	n := 1
// 	for n < length {
// 		id, _ := blog.NewId(n)
// 		title, _ := blog.NewTitle("タイトル" + strconv.Itoa(n))
// 		body, _ := blog.NewBody("内容" + strconv.Itoa(n))
// 		blogs = append(blogs, domains.BuildBlog(id, title, body, time.Time{}, time.Time{}))
// 		n++
// 	}
// 	return domains.NewBlogs(blogs)
// }
