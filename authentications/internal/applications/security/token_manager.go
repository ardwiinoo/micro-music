package security

import "time"

type TokenManager interface {
	GenerateToken(payload map[string]interface{}, expiration time.Duration) (string, error)
	VerifyToken(token string) (map[string]interface{}, error)
}