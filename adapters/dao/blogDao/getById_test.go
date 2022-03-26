package blogDao_test

// import (
// 	"go-api/adapters/dao/blogDao"
// 	"go-api/domains"
// 	"go-api/testData"
// 	"testing"
//
// 	_ "github.com/lib/pq"
// 	"github.com/stretchr/testify/assert"
// 	"gorm.io/gorm"
// )
//
// func TestGetById(t *testing.T) {
//
// 	blog := testData.Blog()
// 	dummyBlog := testData.BlogWithID(100) // 100でnot foud起こす
// 	seeds := []interface{}{
// 		&blogDao.BlogDto{
// 			ID:    blog.ID.Value,
// 			Title: blog.Title.Value,
// 			Body:  blog.Body.Value,
// 		},
// 	}
//
// 	db, err := Prepare("user_dao", seeds)
//
// 	sqlDB, _ := db.DB()
// 	defer sqlDB.Close()
//
// 	if err != nil {
// 		t.Fatal(err)
// 	}
//
// 	rw := blogDao.New(db)
//
// 	tests := []struct {
// 		name      string
// 		blog      *domains.Blog
// 		wantError error
// 	}{
// 		{
// 			name:      "Get a blog",
// 			blog:      &blog,
// 			wantError: nil,
// 		},
// 		{
// 			name:      "Return not found error",
// 			blog:      &dummyBlog,
// 			wantError: gorm.ErrRecordNotFound,
// 		},
// 	}
//
// 	for _, tt := range tests {
// 		tt := tt
// 		t.Run(tt.name, func(t *testing.T) {
// 			_, err := rw.GetById(tt.blog.ID.Value)
// 			if tt.wantError == nil {
// 				assert.NoError(t, err)
// 			} else {
// 				assert.Error(t, err)
// 				assert.Equal(t, err, tt.wantError)
// 			}
// 		})
// 	}
// }
