package blog_test

import (
	"go-api/domains/blog"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBodySuccess(t *testing.T) {
	input := "body"
	newBody, err := blog.NewBody(input)
	t.Run("hoge", func(t *testing.T) {
		assert.Equal(t, newBody.Value, input)
		assert.Equal(t, err, nil)
	})
}

func TestNewBodyFail(t *testing.T) {
	input := ""
	newBody, err := blog.NewBody(input)
	t.Run("hoge", func(t *testing.T) {
		assert.Equal(t, newBody.Value, input)
		assert.NotNil(t, err)
	})
}

func TestUpdateBodySuccess(t *testing.T) {
	input := "body"
	updatedBody, err := blog.UpdateBody(&input)
	t.Run("hoge", func(t *testing.T) {
		assert.Equal(t, updatedBody.Value, input)
		assert.Equal(t, err, nil)
	})
}

func TestUpdateBodyFail(t *testing.T) {
	input := ""
	updatedBody, err := blog.UpdateBody(&input)
	t.Run("hoge", func(t *testing.T) {
		assert.Equal(t, updatedBody.Value, input)
		assert.NotNil(t, err)
	})
}

func TestValueSuccess(t *testing.T) {
	input := "body"
	body, _ := blog.NewBody(input)
	t.Run("hoge", func(t *testing.T) {
		assert.Equal(t, body.Value, input)
	})
}
