package blog_test

import (
	"go-api/domains/models/blog"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTitle_Success(t *testing.T) {
	input := "title"
	newTitle, err := blog.NewTitle(input)
	t.Run("hoge", func(t *testing.T) {
		assert.Equal(t, newTitle.Value, input)
		assert.Equal(t, err, nil)
	})
}

func TestNewTitle_Fail(t *testing.T) {
	input := ""
	newTitle, err := blog.NewTitle(input)
	t.Run("hoge", func(t *testing.T) {
		assert.Equal(t, newTitle.Value, input)
		assert.NotNil(t, err)
	})
}

func TestUpdateTitle_Success(t *testing.T) {
	input := "title"
	updatedTitle, err := blog.UpdateTitle(&input)
	t.Run("hoge", func(t *testing.T) {
		assert.Equal(t, updatedTitle.Value, input)
		assert.Equal(t, err, nil)
	})
}

func TestUpdateTitle_Fail(t *testing.T) {
	input := ""
	updatedTitle, err := blog.UpdateTitle(&input)
	t.Run("hoge", func(t *testing.T) {
		assert.Equal(t, updatedTitle.Value, input)
		assert.NotNil(t, err)
	})
}

func TestTitleValue_Success(t *testing.T) {
	input := "title"
	title, _ := blog.NewTitle(input)
	t.Run("hoge", func(t *testing.T) {
		assert.Equal(t, title.Value, input)
	})
}
