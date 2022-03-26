package domains_test

import (
	"go-api/domains"
	"go-api/testData"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBlogsSuccess(t *testing.T) {
	_blogs := testData.Blogs(5)

	t.Run("", func(t *testing.T) {
		assert.Equal(t, len(_blogs.Value), 4)
	})
}

func TestBlogsSizeSuccess(t *testing.T) {
	_blogs := testData.Blogs(5)

	t.Run("", func(t *testing.T) {
		assert.Equal(t, _blogs.Size(), 4)
	})
}

func TestBlogsEmptyBlogsSuccess(t *testing.T) {
	newBlogs := domains.EmptyBlogs()

	t.Run("", func(t *testing.T) {
		assert.Equal(t, newBlogs.Size(), 0)
	})
}

func TestBlogsApplyLimitAndOffset(t *testing.T) {
	var _blogs = testData.Blogs(5)

	t.Run("complete", func(t *testing.T) {
		assert.Equal(t, _blogs.Value, _blogs.ApplyLimitAndOffset(100, 0))
		assert.Equal(t, _blogs.Value, _blogs.ApplyLimitAndOffset(4, 0))
		assert.Equal(t, _blogs.Value, _blogs.ApplyLimitAndOffset(4, -1))
	})
	t.Run("empty", func(t *testing.T) {
		assert.Equal(t, domains.EmptyBlogs().Value, _blogs.ApplyLimitAndOffset(100, 10))
		assert.Equal(t, domains.EmptyBlogs().Value, _blogs.ApplyLimitAndOffset(3, 4))
		assert.Equal(t, domains.EmptyBlogs().Value, _blogs.ApplyLimitAndOffset(-1, 0))
	})
}
