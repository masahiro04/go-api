package blog_test

import (
	"clean_architecture/golang/domains/blog"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTitleSuccess(t *testing.T) {
	input := "title"
	newTitle, err := blog.NewTitle(input)
	t.Run("hoge", func(t *testing.T) {
		assert.Equal(t, newTitle.Value, input)
		assert.Equal(t, err, nil)
	})
}

func TestNewTitleFail(t *testing.T) {
	input := ""
	newTitle, err := blog.NewTitle(input)
	t.Run("hoge", func(t *testing.T) {
		assert.Equal(t, newTitle.Value, input)
		assert.NotNil(t, err)
	})
}

func TestUpdateTitleSuccess(t *testing.T) {
	input := "title"
	updatedTitle, err := blog.UpdateTitle(&input)
	t.Run("hoge", func(t *testing.T) {
		assert.Equal(t, updatedTitle.Value, input)
		assert.Equal(t, err, nil)
	})
}

func TestUpdateTitleFail(t *testing.T) {
	input := ""
	updatedTitle, err := blog.UpdateTitle(&input)
	t.Run("hoge", func(t *testing.T) {
		assert.Equal(t, updatedTitle.Value, input)
		assert.NotNil(t, err)
	})
}

func TestTitleValueSuccess(t *testing.T) {
	input := "title"
	title, _ := blog.NewTitle(input)
	t.Run("hoge", func(t *testing.T) {
		assert.Equal(t, title.Value, input)
	})
}
