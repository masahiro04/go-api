package userDao_test

import (
	"clean_architecture/golang/adapters/dao/userDao"
	"clean_architecture/golang/domains"
	userModel "clean_architecture/golang/domains/user"
	"clean_architecture/golang/testData"
	"testing"
	"time"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestUpdate(t *testing.T) {
	user := testData.User()
	seeds := []interface{}{
		&userDao.UserDto{
			ID:    user.ID.Value,
			Name:  user.Name.Value,
			Email: user.Email.Value,
		},
	}

	db, err := Prepare("user_update_dao", seeds)

	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	if err != nil {
		t.Fatal(err)
	}

	rw := userDao.New(db)

	id, _ := userModel.NewId(user.ID.Value)
	name, _ := userModel.NewName("大久保22")
	email, _ := userModel.NewEmail("text222@example.com")
	updatedUser := domains.BuildUser(id, name, email, time.Time{}, time.Time{})

	tests := []struct {
		name      string
		user      *domains.User
		wantError error
	}{
		{
			name:      "Update a user",
			user:      &updatedUser,
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
			b, err := rw.Update(tt.user.ID.Value, *tt.user)

			if tt.wantError == nil {
				assert.NoError(t, err)
				assert.Equal(t, *tt.user, *b)
			} else {
				assert.Error(t, err)
				assert.Equal(t, err, tt.wantError)
			}
		})
	}
}
