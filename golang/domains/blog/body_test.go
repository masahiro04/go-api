package blog_test

import (
	"clean_architecture/golang/domains/blog"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBodySuccess(t *testing.T) {
	// NOTE(okubo): Body.valueとの比較したいけど、小文字はexportされないので、Value経由で比較
	input := "body"
	newBody, err := blog.NewBody(input)
	t.Run("hoge", func(t *testing.T) {
		assert.Equal(t, newBody.Value(), input)
		assert.Equal(t, err, nil)
	})
}

func TestNewBodyFail(t *testing.T) {
	// NOTE(okubo): 成功参考にerrをテスト
}

func TestUpdateBodySuccess(t *testing.T) {
	// NOTE(okubo): Body.valueとの比較したいけど、小文字はexportされないので、Value経由で比較
	input := "body"
	updatedBody, err := blog.UpdateBody(&input)
	t.Run("hoge", func(t *testing.T) {
		assert.Equal(t, updatedBody.Value(), input)
		assert.Equal(t, err, nil)
	})
}

func TestUpdateBodyFail(t *testing.T) {
	// NOTE(okubo): 成功参考にerrをテスト
}

func TestValueSuccess(t *testing.T) {
	input := "body"
	title, _ := blog.NewBody(input)
	t.Run("hoge", func(t *testing.T) {
		assert.Equal(t, title.Value(), input)
	})
}
