package userDao_test

import (
	"clean_architecture/golang/adapters/dao/userDao"
	"clean_architecture/golang/testData"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAll(t *testing.T) {

	users := testData.Users(5)

	seeds := []interface{}{
		&userDao.UserDto{
			ID:    users.Value[0].ID.Value,
			Name:  users.Value[0].Name.Value,
			Email: users.Value[0].Email.Value,
		},
		&userDao.UserDto{
			ID:    users.Value[1].ID.Value,
			Name:  users.Value[1].Name.Value,
			Email: users.Value[1].Email.Value,
		},
		&userDao.UserDto{
			ID:    users.Value[2].ID.Value,
			Name:  users.Value[2].Name.Value,
			Email: users.Value[2].Email.Value,
		},
		&userDao.UserDto{
			ID:    users.Value[3].ID.Value,
			Name:  users.Value[3].Name.Value,
			Email: users.Value[3].Email.Value,
		},
	}

	db, err := Prepare("users_dao", seeds)

	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	if err != nil {
		t.Fatal(err)
	}

	rw := userDao.New(db)

	tests := []struct {
		name      string
		length    int
		wantError error
	}{
		{
			name:      "Get users",
			length:    users.Size(),
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
