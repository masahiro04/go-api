package blogDao_test

import (
	"go-api/adapters/dao/blogDao"
	"go-api/domains"
	factories "go-api/test/factories"
	testhelpers "go-api/test/testHelpers"
	"testing"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	blog := factories.Blog()
	seeds := []interface{}{}

	db, err := testhelpers.Prepare("user_create_dao", seeds)

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
			name:      "Create a blog",
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
			_, err := rw.Create(blog)
			if tt.wantError == nil {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
				assert.Equal(t, err, tt.wantError)
			}
		})
	}
}
