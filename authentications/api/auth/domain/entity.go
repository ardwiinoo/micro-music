package domain

import (
	"strings"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"github.com/ardwiinoo/micro-music/authentications/api/auth/app"
	"github.com/ardwiinoo/micro-music/authentications/internal/exception"

)

type UserEntity struct {
	Id        int       `db:"id"`
	Email     string    `db:"email"`
	FullName  string    `db:"full_name"`
	Password  string    `db:"password"`
	PublicId  uuid.UUID `db:"public_id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func NewFromRegisterRequest(req app.RegisterRequestPayload) UserEntity {
	return UserEntity{
		Email:    req.Email,
		FullName: req.FullName,
		Password: req.Password,
		PublicId: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func NewFromLoginRequest(req app.LoginRequestPayload) UserEntity {
	return UserEntity{
		Email:    req.Email,
		Password: req.Password,
	}
}

func (u UserEntity) Validate() (err error) {
	if err = u.ValidateEmail(); err != nil {
		return
	}

	if err = u.ValidatePassword(); err != nil {
		return
	}

	return
}

func (u UserEntity) ValidateEmail() (err error) {
	if u.Email == "" {
		return exception.ErrEmailRequired
	}

	emails := strings.Split(u.Email, "@")
	
	if (len(emails) < 2) {
		return exception.ErrEmailInvalid
	}

	return
}

func (u UserEntity) ValidatePassword() (err error) {
	if u.Password == "" {
		return exception.ErrPasswordRequired
	}

	if len(u.Password) < 8 {
		return exception.ErrPasswordInvalidLength
	}

	return
}

func (u UserEntity) IsExists() bool {
	return u.Id != 0
}

func (u *UserEntity) EncryptPassword(salt int) (err error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)

	if err != nil {
		return
	}

	u.Password = string(hashedPassword)
	return nil
}

func (u UserEntity) VerifyPasswordFromEncrypted(plain string) (err error) {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(plain))
}

func (u UserEntity) VerifyPasswordFromPlain(encrypted string) (err error) {
	return bcrypt.CompareHashAndPassword([]byte(encrypted), []byte(u.Password))
}