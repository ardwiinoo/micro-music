package usecase

import (
	"context"
	"strconv"

	"github.com/ardwiinoo/micro-music/authentications/internal/applications/security"
	"github.com/ardwiinoo/micro-music/authentications/internal/domains/users"
	"github.com/ardwiinoo/micro-music/authentications/internal/domains/users/entities"
	"github.com/ardwiinoo/micro-music/authentications/internal/infrastructures/broker/rabbitmq"
)

type RegisterUserUseCase interface {
	Execute(ctx context.Context, request entities.RegisterUser) (userId int, err error)
}

type registerUserUseCase struct {
	userRepository users.UserRepository
	passwordHash security.PasswordHash
	rabbitMQ       *rabbitmq.RabbitMQ
	
}

func NewRegisterUserUseCase(userRepository users.UserRepository, passwordHash security.PasswordHash, rabbitMQ *rabbitmq.RabbitMQ) RegisterUserUseCase {
	return &registerUserUseCase{
		userRepository: userRepository,
		passwordHash: passwordHash,
		rabbitMQ: rabbitMQ,
	}
}

func (r *registerUserUseCase) Execute(ctx context.Context, payload entities.RegisterUser) (userId int, err error) {

	err = payload.Validate()
	if err != nil {
		return 0, err
	}

	err = r.userRepository.VerifyAvailableEmail(ctx, payload.Email)
	if err != nil {
		return 0, err
	}

	hashedPassword, err := r.passwordHash.Hash(payload.Password)
	if err != nil {
		return 0, err
	}

	payload.Password = hashedPassword
	id, err := r.userRepository.AddUser(ctx, payload)
	if err != nil {
		return 0, err
	}

	eventPayload := map[string]string{
		"user_id": strconv.Itoa(id),
		"email":   payload.Email,
		"name":    payload.FullName,
	}
	
	err = r.rabbitMQ.PublishEvent("email_events", "user_registered", eventPayload)
	if err != nil {
		return 0, err
	}

	return id, nil
}