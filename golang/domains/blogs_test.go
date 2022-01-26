package domains_test

import (
	"clean_architecture/golang/domains"
	"clean_architecture/golang/testData"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBlogsSuccess(t *testing.T) {
	_blogs := testData.Blogs(5)

	newBlogs := domains.NewBlogs(_blogs)

	t.Run("", func(t *testing.T) {
		assert.Equal(t, len(newBlogs.Value()), 4)
	})
}

func TestBlogsValueSuccess(t *testing.T) {
	_blogs := testData.Blogs(5)

	newBlogs := domains.NewBlogs(_blogs)

	t.Run("", func(t *testing.T) {
		assert.Equal(t, newBlogs.Value(), _blogs)
	})
}

func TestBlogsSizeSuccess(t *testing.T) {
	_blogs := testData.Blogs(5)

	newBlogs := domains.NewBlogs(_blogs)

	t.Run("", func(t *testing.T) {
		assert.Equal(t, newBlogs.Size(), 4)
	})
}

func TestBlogsEmptyBlogsSuccess(t *testing.T) {
	newBlogs := domains.EmptyBlogs()

	t.Run("", func(t *testing.T) {
		assert.Equal(t, newBlogs.Size(), 0)
	})
}

func TestBlogsApplyLimitAndOffset(t *testing.T) {
	var testBlogs = testData.Blogs(5)
	var blogs = domains.NewBlogs(testBlogs)

	t.Run("complete", func(t *testing.T) {
		assert.Equal(t, blogs.Value(), blogs.ApplyLimitAndOffset(100, 0))
		assert.Equal(t, blogs.Value(), blogs.ApplyLimitAndOffset(4, 0))
		assert.Equal(t, blogs.Value(), blogs.ApplyLimitAndOffset(4, -1))
	})
	t.Run("empty", func(t *testing.T) {
		assert.Equal(t, domains.EmptyBlogs().Value(), blogs.ApplyLimitAndOffset(100, 10))
		assert.Equal(t, domains.EmptyBlogs().Value(), blogs.ApplyLimitAndOffset(3, 4))
		assert.Equal(t, domains.EmptyBlogs().Value(), blogs.ApplyLimitAndOffset(-1, 0))
	})
}
