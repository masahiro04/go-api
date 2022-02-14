package blogDao_test

import (
	"clean_architecture/golang/adapters/dao/blogDao"
	"clean_architecture/golang/testData"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAll(t *testing.T) {

	blogs := testData.Blogs(5)

	seeds := []interface{}{
		&blogDao.BlogDto{
			ID:    blogs.Value[0].ID.Value,
			Title: blogs.Value[0].Title.Value,
			Body:  blogs.Value[0].Body.Value,
		},
		&blogDao.BlogDto{
			ID:    blogs.Value[1].ID.Value,
			Title: blogs.Value[1].Title.Value,
			Body:  blogs.Value[1].Body.Value,
		},
		&blogDao.BlogDto{
			ID:    blogs.Value[2].ID.Value,
			Title: blogs.Value[2].Title.Value,
			Body:  blogs.Value[2].Body.Value,
		},
		&blogDao.BlogDto{
			ID:    blogs.Value[3].ID.Value,
			Title: blogs.Value[3].Title.Value,
			Body:  blogs.Value[3].Body.Value,
		},
	}

	db, err := Prepare("users_dao", seeds)

	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	if err != nil {
		t.Fatal(err)
	}

	rw := blogDao.New(db)

	tests := []struct {
		name      string
		length    int
		wantError error
	}{
		{
			name:      "Get blogs",
			length:    blogs.Size(),
			wantError: nil,
		},
	}

	for _, tt := range tests {
		// NOTE(okubo): ttにtt入れるとparallelのscopeの問題を回避できるので、一旦そのままで実装してます
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			// TODO(okubo): parallelの方が圧倒的に早いけど、goroutineの影響？で
			// db.Close()のタイミング合わないので、一旦は並行処理は断念
			// t.Parallel()

			b, err := rw.GetAll()
			if tt.wantError == nil {
				assert.NoError(t, err)
				assert.Equal(t, b.Size(), tt.length)
			} else {
				assert.Error(t, err)
			}

		})
	}
}
