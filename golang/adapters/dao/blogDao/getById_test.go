package blogDao_test

import (
	"clean_architecture/golang/adapters/dao/blogDao"
	"clean_architecture/golang/domains"
	"clean_architecture/golang/testData"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-txdb"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewTest() (*gorm.DB, error) {
	conn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		"test-db",
		5432,
		"postgresql",
		"postgresql",
		"test-api",
	)

	txdb.Register("txdb", "postgres", conn)
	dialector := postgres.New(
		postgres.Config{
			DriverName: "txdb",
			DSN:        conn,
		},
	)
	db, _ := gorm.Open(dialector, &gorm.Config{})

	return db, nil
}

func prepare(seeds []interface{}) (*gorm.DB, error) {
	db, err := NewTest()
	if err != nil {
		return nil, err
	}

	for _, s := range seeds {
		if err := db.Create(s).Error; err != nil {
			return nil, err
		}
	}

	return db, nil
}

func TestGetById(t *testing.T) {
	t.Parallel()

	blog := testData.Blog()
	dummyBlog := testData.BlogWithID(100) // 100でnot foud起こす
	seeds := []interface{}{
		&blogDao.BlogDto{
			ID:    blog.ID.Value,
			Title: blog.Title.Value,
			Body:  blog.Body.Value,
		},
	}

	db, err := prepare(seeds)

	if err != nil {
		t.Fatal(err)
	}

	rw := blogDao.New(db)

	tests := map[string]struct {
		blog  *domains.Blog
		noErr bool
	}{
		"Get a blog": {
			blog:  &blog,
			noErr: true,
		},
		"Not found": {
			blog:  &dummyBlog,
			noErr: false,
		},
	}

	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			b, err := rw.GetById(tt.blog.ID.Value)

			if tt.noErr {
				assert.NoError(t, err)
				assert.Equal(t, b.Title.Value, blog.Title.Value)
				assert.Equal(t, b.Body.Value, blog.Body.Value)
			} else {
				assert.Error(t, err)
			}

		})
	}
}
