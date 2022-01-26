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
		assert.Equal(t, newBlog.Title(), title)
		assert.Equal(t, newBlog.Body(), body)
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
		assert.Equal(t, newBlog.Title(), title)
		assert.Equal(t, newBlog.Body(), body)
	})
}

func TestIDSuccess(t *testing.T) {
	_blog := testData.Blog()
	id, _ := blog.NewId(_blog.ID().Value())

	t.Run("", func(t *testing.T) {
		assert.Equal(t, _blog.ID(), id)
	})
}
func TestTitleSuccess(t *testing.T) {
	_blog := testData.Blog()
	title, _ := blog.NewTitle(_blog.Title().Value())

	t.Run("", func(t *testing.T) {
		assert.Equal(t, _blog.Title(), title)
	})
}
func TestBodySuccess(t *testing.T) {
	_blog := testData.Blog()
	body, _ := blog.NewBody(_blog.Body().Value())

	t.Run("", func(t *testing.T) {
		assert.Equal(t, _blog.Body(), body)
	})
}

func TestCreatedAtSuccess(t *testing.T) {
	_blog := testData.Blog()

	t.Run("", func(t *testing.T) {
		assert.Equal(t, _blog.CreatedAt(), _blog.CreatedAt())
	})
}

func TestUpdatedAtSuccess(t *testing.T) {
	_blog := testData.Blog()

	t.Run("", func(t *testing.T) {
		assert.Equal(t, _blog.UpdatedAt(), _blog.UpdatedAt())
	})
}

func TestUpdateTitleSuccess(t *testing.T) {
	_blog := testData.Blog()
	input, _ := blog.NewTitle("updatedTitle")
	updatedTitle := _blog.UpdateTitle(input)

	t.Run("", func(t *testing.T) {
		assert.Equal(t, updatedTitle.Title(), input)
	})
}

func TestUpdateBodySuccess(t *testing.T) {
	_blog := testData.Blog()
	input, _ := blog.NewBody("updatedBody")
	updatedBody := _blog.UpdateBody(input)

	t.Run("", func(t *testing.T) {
		assert.Equal(t, updatedBody.Body(), input)
	})
}

// func TestBlogCollection_ApplyLimitAndOffset(t *testing.T) {
// 	var testBlogs = testData.Blogs(5)
// 	var blog1 = &testBlogs[0]
// 	var blog2 = &testBlogs[1]
// 	var blog3 = &testBlogs[2]
// 	var blog4 = &testBlogs[3]
// 	var blogs = domains.BlogCollection{blog1, blog2, blog3, blog4}
//
// 	t.Run("complete", func(t *testing.T) {
// 		assert.Equal(t, blogs, blogs.ApplyLimitAndOffset(100, 0))
// 		assert.Equal(t, blogs, blogs.ApplyLimitAndOffset(4, 0))
// 		assert.Equal(t, blogs, blogs.ApplyLimitAndOffset(4, -1))
// 	})
// 	t.Run("empty", func(t *testing.T) {
// 		assert.Equal(t, domains.BlogCollection{}, blogs.ApplyLimitAndOffset(100, 10))
// 		assert.Equal(t, domains.BlogCollection{}, blogs.ApplyLimitAndOffset(3, 4))
// 		assert.Equal(t, domains.BlogCollection{}, blogs.ApplyLimitAndOffset(-1, 0))
// 	})
// }
