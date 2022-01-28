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
		assert.Equal(t, newBlog.Title, title)
		assert.Equal(t, newBlog.Body, body)
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
		assert.Equal(t, newBlog.Title, title)
		assert.Equal(t, newBlog.Body, body)
	})
}

func TestUpdatedAtSuccess(t *testing.T) {
	_blog := testData.Blog()

	t.Run("", func(t *testing.T) {
		assert.Equal(t, _blog.UpdatedAt, _blog.UpdatedAt)
	})
}

func TestUpdateTitleSuccess(t *testing.T) {
	_blog := testData.Blog()
	input, _ := blog.NewTitle("updatedTitle")
	updatedTitle := _blog.UpdateTitle(input)

	t.Run("", func(t *testing.T) {
		assert.Equal(t, updatedTitle.Title, input)
	})
}

func TestUpdateBodySuccess(t *testing.T) {
	_blog := testData.Blog()
	input, _ := blog.NewBody("updatedBody")
	updatedBody := _blog.UpdateBody(input)

	t.Run("", func(t *testing.T) {
		assert.Equal(t, updatedBody.Body, input)
	})
}
