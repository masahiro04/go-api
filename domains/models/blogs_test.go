package models_test

import (
	"go-api/domains/models"
	factories "go-api/test/factories"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBlogs_Success(t *testing.T) {
	_blogs := factories.Blogs(5)

	t.Run("", func(t *testing.T) {
		assert.Equal(t, len(_blogs.Value), 4)
	})
}

func TestBlogsSize_Success(t *testing.T) {
	_blogs := factories.Blogs(5)

	t.Run("", func(t *testing.T) {
		assert.Equal(t, _blogs.Size(), 4)
	})
}

func TestBlogsEmptyBlogs_Success(t *testing.T) {
	newBlogs := models.EmptyBlogs()

	t.Run("", func(t *testing.T) {
		assert.Equal(t, newBlogs.Size(), 0)
	})
}

func TestBlogsApplyLimitAndOffset_Success(t *testing.T) {
	var _blogs = factories.Blogs(5)

	t.Run("complete", func(t *testing.T) {
		assert.Equal(t, _blogs.Value, _blogs.ApplyLimitAndOffset(100, 0))
		assert.Equal(t, _blogs.Value, _blogs.ApplyLimitAndOffset(4, 0))
		assert.Equal(t, _blogs.Value, _blogs.ApplyLimitAndOffset(4, -1))
	})
	t.Run("empty", func(t *testing.T) {
		assert.Equal(t, models.EmptyBlogs().Value, _blogs.ApplyLimitAndOffset(100, 10))
		assert.Equal(t, models.EmptyBlogs().Value, _blogs.ApplyLimitAndOffset(3, 4))
		assert.Equal(t, models.EmptyBlogs().Value, _blogs.ApplyLimitAndOffset(-1, 0))
	})
}
