package blog_test

import (
	"go-api/domains/models/blog"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIdNewID_Success(t *testing.T) {
	// NOTE(okubo): Body.valueとの比較したいけど、小文字はexportされないので、Value経由で比較
	input := 1
	newID, err := blog.NewId(input)
	t.Run("hoge", func(t *testing.T) {
		assert.Equal(t, newID.Value, input)
		assert.Equal(t, err, nil)
	})
}

func TestIdString_Success(t *testing.T) {
	input := 1
	newID, err := blog.NewId(input)

	t.Run("hoge", func(t *testing.T) {
		assert.Equal(t, newID.String(), strconv.Itoa(input))
		assert.Equal(t, err, nil)
	})
}

func TestIDValue_Success(t *testing.T) {
	input := 1
	newID, _ := blog.NewId(input)
	t.Run("hoge", func(t *testing.T) {
		assert.Equal(t, newID.Value, input)
	})
}
