package repository

import (
	"context"
	"database/sql"
	"github.com/ardwiinoo/micro-music/users/internal/commons/exceptions"
	"github.com/ardwiinoo/micro-music/users/internal/domains/users"
	"github.com/ardwiinoo/micro-music/users/internal/domains/users/entities"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"time"
)

type userRepositoryPostgres struct {
	db *sqlx.DB
}

func (u userRepositoryPostgres) DeleteUserById(ctx context.Context, userId int) (err error) {
	query := `DELETE FROM users WHERE id = $1`

	_, err = u.db.ExecContext(ctx, query, userId)

	return
}

func (u userRepositoryPostgres) GetDetailUserByPublicId(ctx context.Context, publicId string) (userDetail entities.DetailUser, err error) {
	query := `SELECT * FROM users WHERE public_id = $1`

	err = u.db.GetContext(ctx, &userDetail, query, publicId)

	return
}

// VerifyAvailableEmail implements users.UserRepository.
func (u userRepositoryPostgres) VerifyAvailableEmail(ctx context.Context, email string) (err error) {
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

// AddUser implements users.UserRepository.
func (u userRepositoryPostgres) AddUser(ctx context.Context, payload entities.AddUser) (publicId uuid.UUID, err error) {
	query := `
		INSERT INTO users
			(email, password, role, public_id, full_name, created_at, updated_at)
		VALUES 
		    (:email, :password, :role, :public_id, :full_name, :created_at, :updated_at)
		RETURNING public_id
	`

	params := map[string]interface{}{
		"email":      payload.Email,
		"password":   payload.Password,
		"role":       payload.Role,
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

	err = stmt.QueryRowxContext(ctx, params).Scan(&publicId)

	if err != nil {
		return
	}

	return
}

// GetListUser implements users.UserRepository.
func (u userRepositoryPostgres) GetListUser(ctx context.Context) (users []entities.DetailUser, err error) {
	query := `SELECT * FROM users`

	err = u.db.SelectContext(ctx, &users, query)

	return
}

func NewUserRepositoryPostgres(db *sqlx.DB) users.UserRepository {
	return &userRepositoryPostgres{
		db: db,
	}
}
