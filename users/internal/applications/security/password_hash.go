package security

type PasswordHash interface {
	Hash(password string) (hashedPassword string, err error)
	Compare(plainPassword string, encryptedPassword string) (err error)
}
