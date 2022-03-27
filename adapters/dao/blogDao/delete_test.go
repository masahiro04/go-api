package blogDao_test

import (
	"go-api/adapters/dao/blogDao"
	"go-api/domains"
	factories "go-api/test/factories"
	"testing"

	testhelpers "go-api/test/testHelpers"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestDelete(t *testing.T) {
	blog := factories.Blog()
	seeds := []interface{}{
		&blogDao.BlogDto{
			ID:    blog.ID.Value,
			Title: blog.Title.Value,
			Body:  blog.Body.Value,
		},
	}

	db, err := testhelpers.Prepare("user_delete_dao", seeds)

	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	if err != nil {
		t.Fatal(err)
	}

	rw := blogDao.New(db)

	tests := []struct {
		name      string
		blog      *domains.Blog
		wantError error
	}{
		{
			name:      "Delete a blog",
			blog:      &blog,
			wantError: nil,
		},
		// {
		// 	name:      "Return not found error",
		// 	blog:      &dummyBlog,
		// 	wantError: gorm.ErrRecordNotFound,
		// },
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			err := rw.Delete(tt.blog.ID.Value)

			if tt.wantError == nil {
				assert.NoError(t, err)
				// assert.Equal(t, *tt.blog, *b)
			} else {
				assert.Error(t, err)
				assert.Equal(t, err, tt.wantError)
			}
		})
	}
}
