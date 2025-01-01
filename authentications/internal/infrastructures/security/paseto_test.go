package security_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/ardwiinoo/micro-music/authentications/config"
	"github.com/ardwiinoo/micro-music/authentications/internal/infrastructures/security"
)

func init() {
	if err := config.LoadConfig("../../../.env"); err != nil {
		panic(err)
	}
}

func TestPasetoTokenManager(t *testing.T) {
	tokenManager := security.NewPasetoTokenManager(config.Cfg.App.AppSecret.AppPrivateKey, config.Cfg.App.AppSecret.AppPublicKey)

	t.Run("should generate and verify token successfully", func(t *testing.T) {
		payload := map[string]interface{}{
			"user_id": "12345",
			"role":    "admin",
		}
		expiration := time.Minute * 5

		token, err := tokenManager.GenerateToken(payload, expiration)
		assert.NoError(t, err)
		assert.NotEmpty(t, token)

		verifiedPayload, err := tokenManager.VerifyToken(token)
		assert.NoError(t, err)
		assert.Equal(t, "12345", verifiedPayload["user_id"])
		assert.Equal(t, "admin", verifiedPayload["role"])
	})

	t.Run("should fail to verify expired token", func(t *testing.T) {
		payload := map[string]interface{}{
			"user_id": "12345",
			"role":    "admin",
		}
		expiration := -time.Minute * 1 // Expired token

		token, err := tokenManager.GenerateToken(payload, expiration)
		assert.NoError(t, err)

		_, err = tokenManager.VerifyToken(token)
		assert.Error(t, err)
		assert.EqualError(t, err, "token has expired")
	})
}