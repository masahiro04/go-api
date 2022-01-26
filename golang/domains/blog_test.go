package domains_test

import (
	"clean_architecture/golang/domains"
	"clean_architecture/golang/domains/blog"
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
