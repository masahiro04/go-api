package userDao_test

import (
	"go-api/adapters/dao/userDao"
	"go-api/domains"
	factories "go-api/test/factories"
	testhelpers "go-api/test/testHelpers"
	"testing"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestDelete(t *testing.T) {
	user := factories.User()
	seeds := []interface{}{
		&userDao.UserDto{
			ID:    user.ID.Value,
			Name:  user.Name.Value,
			Email: user.Email.Value,
		},
	}

	db, err := testhelpers.Prepare("user_delete_dao", seeds)

	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	if err != nil {
		t.Fatal(err)
	}

	rw := userDao.New(db)

	tests := []struct {
		name      string
		user      *domains.User
		wantError error
	}{
		{
			name:      "Delete a user",
			user:      &user,
			wantError: nil,
		},
		// {
		// 	name:      "Return not found error",
		// 	user:      &dummyUser,
		// 	wantError: gorm.ErrRecordNotFound,
		// },
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			err := rw.Delete(tt.user.ID.Value)

			if tt.wantError == nil {
				assert.NoError(t, err)
				// assert.Equal(t, *tt.user, *b)
			} else {
				assert.Error(t, err)
				assert.Equal(t, err, tt.wantError)
			}
		})
	}
}
