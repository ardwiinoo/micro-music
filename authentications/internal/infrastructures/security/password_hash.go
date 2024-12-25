package security

import (
	"golang.org/x/crypto/bcrypt"

	"github.com/ardwiinoo/micro-music/authentications/internal/applications/security"
)

type passwordHash struct {
}

func NewPasswordHash() security.PasswordHash {
	return &passwordHash{}
}

// Compare implements security.PasswordHash.
func (p *passwordHash) Compare(plainPassword string, encryptedPassword string) (err error) {
	return bcrypt.CompareHashAndPassword([]byte(encryptedPassword), []byte(plainPassword))
}

// Hash implements security.PasswordHash.
func (p *passwordHash) Hash(password string) (hashedPassword string, err error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}
	
	return string(hashed), nil
}