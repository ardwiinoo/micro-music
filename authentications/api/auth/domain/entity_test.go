package domain

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ardwiinoo/micro-music/authentications/internal/exception"
)

func TestValidateUserEntity(t *testing.T) {
	t.Run("it should validate user entity successfully", func(t *testing.T) {
		user := UserEntity{
			Email: "ardwiinoo@gmail.com",
			Password: "secretpassword",
		}

		err := user.Validate()
		require.Nil(t, err)
	})

	t.Run("it should throw ErrEmailRequired when email is not present", func(t *testing.T) {
		user := UserEntity{
			Email: "",
			Password: "secretpassword",
		}

		err := user.Validate()
		require.NotNil(t, err)
		require.Equal(t, exception.ErrEmailRequired, err)
	})

	t.Run("it should throw ErrEmailInvalid when email is not valid email", func(t *testing.T) {
		user := UserEntity{
			Email: "invalid",
			Password: "secretpassword",
		}

		err := user.Validate()
		require.NotNil(t, err)
		require.Equal(t, exception.ErrEmailInvalid, err)
	})

	t.Run("it should throw ErrPasswordRequired when password is not present", func(t *testing.T) {
		user := UserEntity{
			Email: "ardwiinoo@gmail.com",
			Password: "",
		}

		err := user.Validate()
		require.NotNil(t, err)
		require.Equal(t, exception.ErrPasswordRequired, err)
	})

	t.Run("it should throw ErrPasswordInvalidLength when password less than 8 characters", func(t *testing.T) {
		user := UserEntity{
			Email: "ardwiinoo@gmail.com",
			Password: "passwor",
		}

		err := user.Validate()
		require.NotNil(t, err)
		require.Equal(t, exception.ErrPasswordInvalidLength, err)
	})
}

func TestIsExistsUser(t *testing.T) {
	t.Run("it should return true when user exists", func(t *testing.T) {
		user := UserEntity{
			Id:    1,
			Email: "ardwiinoo@gmail.com",
		}

		isExists := user.IsExists()

		require.NotNil(t, isExists)
		require.Equal(t, true, isExists)
	})

	t.Run("it should return false when user does not exist", func(t *testing.T) { 
		user := UserEntity{
			Id:    0,
			Email: "notexists@gmail.com",
		}

		isExists := user.IsExists()

		require.NotNil(t, isExists)
		require.Equal(t, false, isExists)
	})
}

func TestEncryptPassword(t *testing.T) {
	t.Run("it should encrypt password correctly", func(t *testing.T) {
		user := UserEntity{
			Email: "ardwiinoo@gmail.com",
			Password: "secretpassword",
		}

		err := user.EncryptPassword(10)
		require.NoError(t, err)
		require.NotEqual(t, "secretpassword", user.Password)
		require.Greater(t, len(user.Password), 0)
	})
}

func TestVerifyPasswordFromEncrypted(t *testing.T) {
	t.Run("it should return no error for correct password", func(t *testing.T) {
		user := UserEntity{
			Password: "$2y$10$pH/ln0UgI1OulyUL3UI7EuE.XVMCZ6FAvsEu5L73.H7GNrdDVMpxW", // "secretpassword"
		}

		err := user.VerifyPasswordFromEncrypted("secretpassword")
		require.NoError(t, err)
	})

	t.Run("it should return error for incorrect password", func(t *testing.T) {
		user := UserEntity{
			Password: "$2y$10$pH/ln0UgI1OulyUL3UI7EuE.XVMCZ6FAvsEu5L73.H7GNrdDVMpxW", // "secretpassword"
		}

		err := user.VerifyPasswordFromEncrypted("wrongpassword")
		require.Error(t, err)
	})
}

func TestVerifyPasswordFromPlain(t *testing.T) {
	t.Run("it should return no error for correct password", func(t *testing.T) {
		encryptedPassword := "$2y$10$pH/ln0UgI1OulyUL3UI7EuE.XVMCZ6FAvsEu5L73.H7GNrdDVMpxW" // "secretpassword"
		user := UserEntity{
			Password: "secretpassword",
		}

		err := user.VerifyPasswordFromPlain(encryptedPassword)
		require.NoError(t, err)
	})

	t.Run("it should return error for incorrect password", func(t *testing.T) {
		encryptedPassword := "$2y$10$pH/ln0UgI1OulyUL3UI7EuE.XVMCZ6FAvsEu5L73.H7GNrdDVMpxW" // "secretpassword"
		user := UserEntity{
			Password: "wrongpassword",
		}

		err := user.VerifyPasswordFromPlain(encryptedPassword)
		require.Error(t, err)
	})
}