package models_test

import (
	"go-api/domains/models"
	"go-api/domains/models/user"

	"go-api/test/factories"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewUser_Success(t *testing.T) {
	// NOTE(okubo): Body.valueとの比較したいけど、小文字はexportされないので、Value経由で比較
	uuidInput := "uuidinput"
	nameInput := "大久保"
	emailInput := "test@example.con"

	uuid, _ := user.NewUUID(uuidInput)
	name, _ := user.NewName(nameInput)
	email, _ := user.NewEmail(emailInput)
	password, _ := user.NewPassword("hogehgoe")
	newUser := models.NewUser(uuid, name, email, password)
	t.Run("", func(t *testing.T) {
		assert.Equal(t, newUser.Name, name)
		assert.Equal(t, newUser.Email, email)
	})
}

func TestBuildUser_Success(t *testing.T) {
	idInput := 1

	uuidInput := "uuidinput"
	nameInput := "大久保"
	emailInput := "test@example.con"

	id, _ := user.NewId(idInput)
	uuid, _ := user.NewUUID(uuidInput)
	name, _ := user.NewName(nameInput)
	email, _ := user.NewEmail(emailInput)

	newUser := models.BuildUser(id, uuid, name, email, time.Time{}, time.Time{})

	t.Run("", func(t *testing.T) {
		assert.Equal(t, newUser.Name, name)
		assert.Equal(t, newUser.Email, email)
	})
}

// func TestUpdatedAt_Success(t *testing.T) {
// 	_user := factories.User()
//
// 	t.Run("", func(t *testing.T) {
// 		assert.Equal(t, _user.UpdatedAt, _user.UpdatedAt)
// 	})
// }

func TestUpdateName_Success(t *testing.T) {
	_user := factories.User()
	input, _ := user.NewName("修正後名前")
	updatedName := _user.UpdateName(input)

	t.Run("", func(t *testing.T) {
		assert.Equal(t, updatedName.Name, input)
	})
}

func TestUpdateEmail_Success(t *testing.T) {
	_user := factories.User()
	input, _ := user.NewEmail("updated@example.com")
	updatedEmail := _user.UpdateEmail(input)

	t.Run("", func(t *testing.T) {
		assert.Equal(t, updatedEmail.Email, input)
	})
}
