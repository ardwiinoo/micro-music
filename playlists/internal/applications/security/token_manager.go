package security

type TokenManager interface {
	VerifyToken(token string) (map[string]interface{}, error)
}