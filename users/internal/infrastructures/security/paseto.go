package security

import (
	"crypto/ed25519"
	"encoding/json"
	"errors"
	"github.com/ardwiinoo/micro-music/users/internal/applications/security"
	"github.com/o1egl/paseto"
	"time"
)

type pasetoTokenManager struct {
	publicKey ed25519.PublicKey
}

func NewPasetoTokenManager(publicKey ed25519.PublicKey) security.TokenManager {
	return &pasetoTokenManager{
		publicKey: publicKey,
	}
}

// VerifyToken implements security.TokenManager.
func (p *pasetoTokenManager) VerifyToken(token string) (map[string]interface{}, error) {
	var claims paseto.JSONToken

	err := paseto.NewV2().Verify(token, p.publicKey, &claims, nil)
	if err != nil {
		return nil, err
	}

	if claims.Expiration.Before(time.Now()) {
		return nil, errors.New("token has expired")
	}

	claimsJSON, err := claims.MarshalJSON()
	if err != nil {
		return nil, err
	}

	var payload map[string]interface{}
	err = json.Unmarshal(claimsJSON, &payload)
	if err != nil {
		return nil, err
	}

	return payload, nil
}
