package user_test

import (
	"go-api/domains/models/user"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPassword_Success(t *testing.T) {
	input := "12345678"
	newPassword, err := user.NewPassword(input)
	t.Run("hoge", func(t *testing.T) {
		assert.Equal(t, newPassword.Value, input)
		assert.Equal(t, err, nil)
	})
}

func TestNewPasswordWithBlankString_Fail(t *testing.T) {
	input := ""
	newPassword, err := user.NewPassword(input)
	t.Run("hoge", func(t *testing.T) {
		assert.Equal(t, newPassword.Value, input)
		assert.NotNil(t, err)
	})
}

func TestNewPasswordWithLessThan6Chars_Fail(t *testing.T) {
	input := "12345"
	newPassword, err := user.NewPassword(input)
	t.Run("hoge", func(t *testing.T) {
		assert.Equal(t, newPassword.Value, input)
		assert.NotNil(t, err)
	})
}

func TestUpdatePassword_Success(t *testing.T) {
	input := "12345678"
	updatedPassword, err := user.UpdatePassword(&input)
	t.Run("hoge", func(t *testing.T) {
		assert.Equal(t, updatedPassword.Value, input)
		assert.Equal(t, err, nil)
	})
}

func TestUpdatePassword_Fail(t *testing.T) {
	input := ""
	updatedPassword, err := user.UpdatePassword(&input)
	t.Run("hoge", func(t *testing.T) {
		assert.Equal(t, updatedPassword.Value, input)
		assert.NotNil(t, err)
	})
}

func TestPasswordValue_Success(t *testing.T) {
	input := "12345678"
	Password, _ := user.NewPassword(input)
	t.Run("hoge", func(t *testing.T) {
		assert.Equal(t, Password.Value, input)
	})
}
