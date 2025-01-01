package entities_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ardwiinoo/micro-music/authentications/internal/domains/users/entities"
)

func TestRegisterUser_Validate(t *testing.T) {
	t.Run("Should return error when missing properties", func(t *testing.T) {
		user := &entities.RegisterUser{
			FullName: "",
			Email: "",
			Password: "",
		}

		err := user.Validate()

		assert.NotNil(t, err)
		assert.EqualError(t, err, "REGISTER_USER.NOT_CONTAIN_NEEDED_PROPERTY")
	})

	t.Run("Should return error when email is invalid", func(t *testing.T) {
		user := &entities.RegisterUser{
			FullName: "test_user",
			Email: "ardwiinoo@",
			Password: "supersecretpassword",
		}

		err := user.Validate()

		assert.NotNil(t, err)
		assert.EqualError(t, err, "REGISTER_USER.EMAIL_INVALID")
	})

	t.Run("Should return error when password is invalid", func(t *testing.T) {
		user := &entities.RegisterUser{
			FullName: "test_user",
			Email: "ardwiinoo@gmail.com",
			Password: "pass",
		}

		err := user.Validate()

		assert.NotNil(t, err)
		assert.EqualError(t, err, "REGISTER_USER.PASSWORD_INVALID_LENGTH")
	})

	t.Run("Shuld return nil when all properties are valid", func(t *testing.T) {
		user := &entities.RegisterUser{
			FullName: "test_user",
			Email: "ardwiinoo@gmail.com",
			Password: "supersecretpassword",
		}

		err := user.Validate()

		assert.Nil(t, err)
	})
}