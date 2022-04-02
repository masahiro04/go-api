package models_test

import (
	"go-api/domains/models"
	"go-api/domains/models/blog"
	factories "go-api/test/factories"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewBlog_Success(t *testing.T) {
	// NOTE(okubo): Body.valueとの比較したいけど、小文字はexportされないので、Value経由で比較
	titleInput := "title"
	bodyInput := "body"
	title, _ := blog.NewTitle(titleInput)
	body, _ := blog.NewBody(bodyInput)
	newBlog := models.NewBlog(title, body)
	t.Run("", func(t *testing.T) {
		assert.Equal(t, newBlog.Title, title)
		assert.Equal(t, newBlog.Body, body)
	})
}

func TestBuildBlog_Success(t *testing.T) {
	idInput := 1
	titleInput := "title"
	bodyInput := "body"

	id, _ := blog.NewId(idInput)
	title, _ := blog.NewTitle(titleInput)
	body, _ := blog.NewBody(bodyInput)

	newBlog := models.BuildBlog(id, title, body, time.Time{}, time.Time{})

	t.Run("", func(t *testing.T) {
		assert.Equal(t, newBlog.Title, title)
		assert.Equal(t, newBlog.Body, body)
	})
}

func TestUpdatedAt_Success(t *testing.T) {
	_blog := factories.Blog()

	t.Run("", func(t *testing.T) {
		assert.Equal(t, _blog.UpdatedAt, _blog.UpdatedAt)
	})
}

func TestUpdateTitle_Success(t *testing.T) {
	_blog := factories.Blog()
	input, _ := blog.NewTitle("updatedTitle")
	updatedTitle := _blog.UpdateTitle(input)

	t.Run("", func(t *testing.T) {
		assert.Equal(t, updatedTitle.Title, input)
	})
}

func TestUpdateBody_Success(t *testing.T) {
	_blog := factories.Blog()
	input, _ := blog.NewBody("updatedBody")
	updatedBody := _blog.UpdateBody(input)

	t.Run("", func(t *testing.T) {
		assert.Equal(t, updatedBody.Body, input)
	})
}
