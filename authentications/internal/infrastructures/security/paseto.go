package security

import (
	"crypto/ed25519"
	"encoding/json"
	"errors"
	"time"

	"github.com/o1egl/paseto"

	"github.com/ardwiinoo/micro-music/authentications/internal/applications/security"
)

type pasetoTokenManager struct {
	privateKey ed25519.PrivateKey
	publicKey  ed25519.PublicKey
}

func NewPasetoTokenManager(privateKey ed25519.PrivateKey, publicKey ed25519.PublicKey) security.TokenManager {
	return &pasetoTokenManager{
		privateKey: privateKey,
		publicKey:  publicKey,
	}
}

// GenerateToken implements security.TokenManager.
func (p *pasetoTokenManager) GenerateToken(payload map[string]interface{}, expiration time.Duration) (string, error) {
	now := time.Now()
	exp := now.Add(expiration)

	claims := paseto.JSONToken{
		IssuedAt:   now,
		Expiration: exp,
	}

	for key, value := range payload {
		valStr, ok := value.(string)
		if !ok {
			jsonValue, err := json.Marshal(value)
			if err != nil {
				return "", err
			}
			valStr = string(jsonValue)
		}
		claims.Set(key, valStr)
	}

	token, err := paseto.NewV2().Sign(p.privateKey, claims, nil)
	if err != nil {
		return "", err
	}

	return token, nil
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