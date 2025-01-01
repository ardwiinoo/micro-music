package security_test

import (
	"testing"

	"github.com/ardwiinoo/micro-music/authentications/internal/infrastructures/security"
	"github.com/stretchr/testify/assert"
)

func TestPasswordHash(t *testing.T) {
	passwordHash := security.NewPasswordHash()

	t.Run("should hash and compare passwords successfully", func(t *testing.T) {
		password := "securepassword"

		hashedPassword, err := passwordHash.Hash(password)
		assert.NoError(t, err)
		assert.NotEmpty(t, hashedPassword)

		err = passwordHash.Compare(password, hashedPassword)
		assert.NoError(t, err)
	})

	t.Run("should fail to compare incorrect password", func(t *testing.T) {
		password := "securepassword"
		wrongPassword := "wrongpassword"

		hashedPassword, err := passwordHash.Hash(password)
		assert.NoError(t, err)

		err = passwordHash.Compare(wrongPassword, hashedPassword)
		assert.Error(t, err)
	})
}