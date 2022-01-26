package domains_test

import (
	"clean_architecture/golang/domains"
	"clean_architecture/golang/domains/blog"
	"clean_architecture/golang/testData"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBlogSuccess(t *testing.T) {
	// NOTE(okubo): Body.valueとの比較したいけど、小文字はexportされないので、Value経由で比較
	titleInput := "title"
	bodyInput := "body"
	title, _ := blog.NewTitle(titleInput)
	body, _ := blog.NewBody(bodyInput)
	newBlog := domains.NewBlog(title, body)
	t.Run("", func(t *testing.T) {
		assert.Equal(t, newBlog.Title.Value(), titleInput)
		assert.Equal(t, newBlog.Body.Value(), bodyInput)
	})
}
func TestBuildBlogSuccess(t *testing.T) {
	idInput := 1
	titleInput := "title"
	bodyInput := "body"

	id, _ := blog.NewId(idInput)
	title, _ := blog.NewTitle(titleInput)
	body, _ := blog.NewBody(bodyInput)

	newBlog := domains.BuildBlog(id, title, body)

	t.Run("", func(t *testing.T) {
		assert.Equal(t, newBlog.Title.Value(), titleInput)
		assert.Equal(t, newBlog.Body.Value(), bodyInput)
	})
}

// func TestUpdateBlogSuccess(t *testing.T) {
// 	initial := testData.Blog()
// 	inputBlog := initial
// 	inputTitle := "updatedTitle"
// 	inputBody := "updatedBody"
// 	title, _ := blog.UpdateTitle(&inputTitle)
// 	body, _ := blog.UpdateBody(&inputBody)
//
// 	updatedBlog := domains.UpdateBlog(initial, title, body)
//
// 	t.Run("", func(t *testing.T) {
// 		assert.Equal(t, newBlog.Title.Value(), titleInput)
// 		assert.Equal(t, newBlog.Body.Value(), bodyInput)
// 	})
// }

func TestBlogCollection_ApplyLimitAndOffset(t *testing.T) {
	var testBlogs = testData.Blogs(5)
	var blog1 = &testBlogs[0]
	var blog2 = &testBlogs[1]
	var blog3 = &testBlogs[2]
	var blog4 = &testBlogs[3]
	var blogs = domains.BlogCollection{blog1, blog2, blog3, blog4}

	t.Run("complete", func(t *testing.T) {
		assert.Equal(t, blogs, blogs.ApplyLimitAndOffset(100, 0))
		assert.Equal(t, blogs, blogs.ApplyLimitAndOffset(4, 0))
		assert.Equal(t, blogs, blogs.ApplyLimitAndOffset(4, -1))
	})
	t.Run("empty", func(t *testing.T) {
		assert.Equal(t, domains.BlogCollection{}, blogs.ApplyLimitAndOffset(100, 10))
		assert.Equal(t, domains.BlogCollection{}, blogs.ApplyLimitAndOffset(3, 4))
		assert.Equal(t, domains.BlogCollection{}, blogs.ApplyLimitAndOffset(-1, 0))
	})
}
