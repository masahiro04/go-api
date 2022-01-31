package user_test

import (
	"clean_architecture/golang/domains/user"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPasswordSuccess(t *testing.T) {
	input := "12345678"
	newPassword, err := user.NewPassword(input)
	t.Run("hoge", func(t *testing.T) {
		assert.Equal(t, newPassword.Value, input)
		assert.Equal(t, err, nil)
	})
}

func TestNewPasswordFailWithBlankString(t *testing.T) {
	input := ""
	newPassword, err := user.NewPassword(input)
	t.Run("hoge", func(t *testing.T) {
		assert.Equal(t, newPassword.Value, input)
		assert.NotNil(t, err)
	})
}

func TestNewPasswordFailWithLessThan6Chars(t *testing.T) {
	input := "12345"
	newPassword, err := user.NewPassword(input)
	t.Run("hoge", func(t *testing.T) {
		assert.Equal(t, newPassword.Value, input)
		assert.NotNil(t, err)
	})
}

func TestUpdatePasswordSuccess(t *testing.T) {
	input := "12345678"
	updatedPassword, err := user.UpdatePassword(&input)
	t.Run("hoge", func(t *testing.T) {
		assert.Equal(t, updatedPassword.Value, input)
		assert.Equal(t, err, nil)
	})
}

func TestUpdatePasswordFail(t *testing.T) {
	input := ""
	updatedPassword, err := user.UpdatePassword(&input)
	t.Run("hoge", func(t *testing.T) {
		assert.Equal(t, updatedPassword.Value, input)
		assert.NotNil(t, err)
	})
}

func TestPasswordValueSuccess(t *testing.T) {
	input := "12345678"
	Password, _ := user.NewPassword(input)
	t.Run("hoge", func(t *testing.T) {
		assert.Equal(t, Password.Value, input)
	})
}
