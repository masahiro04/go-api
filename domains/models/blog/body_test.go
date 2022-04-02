package blog_test

import (
	"go-api/domains/models/blog"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBody_Success(t *testing.T) {
	input := "body"
	newBody, err := blog.NewBody(input)
	t.Run("hoge", func(t *testing.T) {
		assert.Equal(t, newBody.Value, input)
		assert.Equal(t, err, nil)
	})
}

func TestNewBody_Fail(t *testing.T) {
	input := ""
	newBody, err := blog.NewBody(input)
	t.Run("hoge", func(t *testing.T) {
		assert.Equal(t, newBody.Value, input)
		assert.NotNil(t, err)
	})
}

func TestUpdateBody_Success(t *testing.T) {
	input := "body"
	updatedBody, err := blog.UpdateBody(&input)
	t.Run("hoge", func(t *testing.T) {
		assert.Equal(t, updatedBody.Value, input)
		assert.Equal(t, err, nil)
	})
}

func TestUpdateBody_Fail(t *testing.T) {
	input := ""
	updatedBody, err := blog.UpdateBody(&input)
	t.Run("hoge", func(t *testing.T) {
		assert.Equal(t, updatedBody.Value, input)
		assert.NotNil(t, err)
	})
}

func TestValue_Success(t *testing.T) {
	input := "body"
	body, _ := blog.NewBody(input)
	t.Run("hoge", func(t *testing.T) {
		assert.Equal(t, body.Value, input)
	})
}
