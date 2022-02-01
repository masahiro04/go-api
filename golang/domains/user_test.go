package domains_test

import (
	"clean_architecture/golang/domains"
	"clean_architecture/golang/domains/user"
	"clean_architecture/golang/testData"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewUserSuccess(t *testing.T) {
	// NOTE(okubo): Body.valueとの比較したいけど、小文字はexportされないので、Value経由で比較
	nameInput := "大久保"
	emailInput := "test@example.con"
	name, _ := user.NewName(nameInput)
	email, _ := user.NewEmail(emailInput)
	newUser := domains.NewUser(name, email)
	t.Run("", func(t *testing.T) {
		assert.Equal(t, newUser.Name, name)
		assert.Equal(t, newUser.Email, email)
	})
}

func TestBuildUserSuccess(t *testing.T) {
	idInput := 1
	nameInput := "大久保"
	emailInput := "test@example.con"

	id, _ := user.NewId(idInput)

	name, _ := user.NewName(nameInput)
	email, _ := user.NewEmail(emailInput)

	newUser := domains.BuildUser(id, name, email, time.Time{}, time.Time{})

	t.Run("", func(t *testing.T) {
		assert.Equal(t, newUser.Name, name)
		assert.Equal(t, newUser.Email, email)
	})
}

// func TestUpdatedAtSuccess(t *testing.T) {
// 	_user := testData.User()
//
// 	t.Run("", func(t *testing.T) {
// 		assert.Equal(t, _user.UpdatedAt, _user.UpdatedAt)
// 	})
// }

func TestUpdateNameSuccess(t *testing.T) {
	_user := testData.User()
	input, _ := user.NewName("修正後名前")
	updatedName := _user.UpdateName(input)

	t.Run("", func(t *testing.T) {
		assert.Equal(t, updatedName.Name, input)
	})
}

func TestUpdateEmailSuccess(t *testing.T) {
	_user := testData.User()
	input, _ := user.NewEmail("updated@example.com")
	updatedEmail := _user.UpdateEmail(input)

	t.Run("", func(t *testing.T) {
		assert.Equal(t, updatedEmail.Email, input)
	})
}
