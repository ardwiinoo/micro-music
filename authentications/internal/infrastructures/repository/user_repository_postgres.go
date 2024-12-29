package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"

	"github.com/ardwiinoo/micro-music/authentications/internal/commons/exceptions"
	"github.com/ardwiinoo/micro-music/authentications/internal/domains/users"
	"github.com/ardwiinoo/micro-music/authentications/internal/domains/users/entities"
)

type userRepositoryPostgres struct {
	db *sqlx.DB
}

func NewUserRepositoryPostgres(db *sqlx.DB) users.UserRepository {
	return &userRepositoryPostgres{
		db: db,
	}
}

// AddUser implements users.UserRepository.
func (u *userRepositoryPostgres) AddUser(ctx context.Context, payload entities.RegisterUser) (id int, err error) {
	query := `
		INSERT INTO users 
			(email, password, role, public_id, full_name, created_at, updated_at)
		VALUES 
			(:email, :password, :role, :public_id, :full_name, :created_at, :updated_at)
		RETURNING id
	`

	params := map[string]interface{}{
		"email":      payload.Email,
		"password":   payload.Password,
		"role":       1,
		"public_id":  uuid.Must(uuid.NewRandom()).String(),
		"full_name":  payload.FullName,
		"created_at": time.Now(),
		"updated_at": time.Now(),
	}

	stmt, err := u.db.PrepareNamedContext(ctx, query)

	if err != nil {
		return
	}

	defer stmt.Close()

	err = stmt.QueryRowxContext(ctx, params).Scan(&id)

	if err != nil {
		return
	}

	return
}

// VerifyAvailableEmail implements users.UserRepository.
func (u *userRepositoryPostgres) VerifyAvailableEmail(ctx context.Context, email string) (err error) {
	query := `SELECT id FROM users WHERE email = $1`

	var id int
	err = u.db.GetContext(ctx, &id, query, email)

	if err == nil {
		return exceptions.InvariantError("email already in used")
	}

	if err == sql.ErrNoRows {
		return nil
	}

	return
}

// GetUserByEmail implements users.UserRepository.
func (u *userRepositoryPostgres) GetUserByEmail(ctx context.Context, email string) (user entities.DetailUser, err error) {
	query := `
		SELECT 
			id, email, full_name, password, public_id, role, created_at, updated_at 
		FROM users 
		WHERE email = $1`

	err = u.db.GetContext(ctx, &user, query, email)

	if err == sql.ErrNoRows {
		return user, exceptions.NotFoundError("user not found")
	}

	if err != nil {
		return
	}

	return
}
