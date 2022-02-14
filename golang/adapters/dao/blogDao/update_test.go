package blogDao_test

import (
	"clean_architecture/golang/adapters/dao/blogDao"
	"clean_architecture/golang/domains"
	blogModel "clean_architecture/golang/domains/blog"
	"clean_architecture/golang/testData"
	"testing"
	"time"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestUpdate(t *testing.T) {
	blog := testData.Blog()
	seeds := []interface{}{
		&blogDao.BlogDto{
			ID:    blog.ID.Value,
			Title: blog.Title.Value,
			Body:  blog.Body.Value,
		},
	}

	db, err := Prepare("user_update_dao", seeds)

	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	if err != nil {
		t.Fatal(err)
	}

	rw := blogDao.New(db)

	id, _ := blogModel.NewId(blog.ID.Value)
	title, _ := blogModel.NewTitle("タイトル22")
	body, _ := blogModel.NewBody("内容22")
	updatedBlog := domains.BuildBlog(id, title, body, time.Time{}, time.Time{})

	tests := []struct {
		name      string
		blog      *domains.Blog
		wantError error
	}{
		{
			name:      "Update a blog",
			blog:      &updatedBlog,
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
			b, err := rw.Update(tt.blog.ID.Value, *tt.blog)

			if tt.wantError == nil {
				assert.NoError(t, err)
				assert.Equal(t, *tt.blog, *b)
			} else {
				assert.Error(t, err)
				assert.Equal(t, err, tt.wantError)
			}
		})
	}
}
