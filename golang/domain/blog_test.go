package domain_test

import (
	"testing"

	"clean_architecture/golang/domain"
	"clean_architecture/golang/testData"
	"github.com/stretchr/testify/assert"
)

func TestBlogUpdate(t *testing.T) {
	blog := testData.Blog()
	updatedBlog := testData.Blog()

	t.Run("nil pointers should leave struct intact", func(t *testing.T) {
		domain.UpdateBlog(&blog,
			domain.SetBlogTitle(nil),
			domain.SetBlogBody(nil),
		)
		assert.Equal(t, updatedBlog, blog)
	})

	t.Run("update all structs", func(t *testing.T) {
		domain.UpdateBlog(&blog,
			domain.SetBlogTitle(&updatedBlog.Title),
			domain.SetBlogBody(&updatedBlog.Body),
		)

		assert.Equal(t, blog.Title, updatedBlog.Title)
		assert.Equal(t, blog.Body, updatedBlog.Body)
	})
}

var blog1 = &domain.Blog{}
var blog2 = &domain.Blog{}
var blog3 = &domain.Blog{}
var blog4 = &domain.Blog{}

func TestBlogCollection_ApplyLimitAndOffset(t *testing.T) {
	blogs := domain.BlogCollection{blog1, blog2, blog3, blog4}

	t.Run("complete", func(t *testing.T) {
		assert.Equal(t, blogs, blogs.ApplyLimitAndOffset(100, 0))
		assert.Equal(t, blogs, blogs.ApplyLimitAndOffset(4, 0))
		assert.Equal(t, blogs, blogs.ApplyLimitAndOffset(4, -1))
	})
	t.Run("empty", func(t *testing.T) {
		assert.Equal(t, domain.BlogCollection{}, blogs.ApplyLimitAndOffset(100, 10))
		assert.Equal(t, domain.BlogCollection{}, blogs.ApplyLimitAndOffset(3, 4))
		assert.Equal(t, domain.BlogCollection{}, blogs.ApplyLimitAndOffset(-1, 0))
	})
}
