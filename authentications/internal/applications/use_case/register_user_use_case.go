package usecase

import (
	"context"

	"github.com/ardwiinoo/micro-music/authentications/config"
	"github.com/ardwiinoo/micro-music/authentications/internal/applications/security"
	"github.com/ardwiinoo/micro-music/authentications/internal/commons/constants"
	"github.com/ardwiinoo/micro-music/authentications/internal/domains/users"
	"github.com/ardwiinoo/micro-music/authentications/internal/domains/users/entities"
	"github.com/ardwiinoo/micro-music/authentications/internal/infrastructures/broker/rabbitmq"
	"github.com/google/uuid"
)

type RegisterUserUseCase interface {
	Execute(ctx context.Context, request entities.RegisterUser) (publicId uuid.UUID, err error)
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

func (r *registerUserUseCase) Execute(ctx context.Context, payload entities.RegisterUser) (publicId uuid.UUID, err error) {

	err = payload.Validate()
	if err != nil {
		return uuid.Nil, err
	}

	err = r.userRepository.VerifyAvailableEmail(ctx, payload.Email)
	if err != nil {
		return uuid.Nil, err
	}

	hashedPassword, err := r.passwordHash.Hash(payload.Password)
	if err != nil {
		return uuid.Nil, err
	}

	payload.Password = hashedPassword
	publicId, err = r.userRepository.AddUser(ctx, payload)
	if err != nil {
		return uuid.Nil, err
	}

	eventPayload := map[string]string{
		"public_id": publicId.String(),
		"email":   payload.Email,
		"name":    payload.FullName,
		"activation_link": config.Cfg.App.VerificationUrl + "/auth/activate/" + publicId.String(),
	}
	
	err = r.rabbitMQ.PublishEvent(constants.QueueTypes.EmailQueue, constants.EventType.UserRegistered, eventPayload)
	if err != nil {
		return uuid.Nil, err
	}

	return publicId, nil
}