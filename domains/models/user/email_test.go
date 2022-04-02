package user_test

import (
	"go-api/domains/models/user"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewEmail_Success(t *testing.T) {
	input := "hogehoge@example.com"
	newEmail, err := user.NewEmail(input)
	t.Run("hoge", func(t *testing.T) {
		assert.Equal(t, newEmail.Value, input)
		assert.Equal(t, err, nil)
	})
}

func TestNewEmailWithBlankString_Fail(t *testing.T) {
	input := ""
	newEmail, err := user.NewEmail(input)
	t.Run("hoge", func(t *testing.T) {
		assert.Equal(t, newEmail.Value, input)
		assert.NotNil(t, err)
	})
}

func TestNewEmailWithNotCorrectFormat_Fail(t *testing.T) {
	input := "hogehoge"
	newEmail, err := user.NewEmail(input)
	t.Run("hoge", func(t *testing.T) {
		assert.Equal(t, newEmail.Value, input)
		assert.NotNil(t, err)
	})
}

func TestUpdateEmail_Success(t *testing.T) {
	input := "hogehoge@example.com"
	updatedEmail, err := user.UpdateEmail(&input)
	t.Run("hoge", func(t *testing.T) {
		assert.Equal(t, updatedEmail.Value, input)
		assert.Equal(t, err, nil)
	})
}

func TestUpdateEmail_Fail(t *testing.T) {
	input := ""
	updatedEmail, err := user.UpdateEmail(&input)
	t.Run("hoge", func(t *testing.T) {
		assert.Equal(t, updatedEmail.Value, input)
		assert.NotNil(t, err)
	})
}

func TestEmailValue_Success(t *testing.T) {
	input := "hogehoge@example.com"
	Email, _ := user.NewEmail(input)
	t.Run("hoge", func(t *testing.T) {
		assert.Equal(t, Email.Value, input)
	})
}
