package usecase

import dto "github.com/ardwiinoo/micro-music/authentications/internal/applications/dto/users"

type RegisterUserUseCase interface {
}

func NewRegisterUserUseCase() RegisterUserUseCase {
	return &registerUserUseCase{}
}

func (uc *RegisterUserUseCase) Excecute(input dto.UserRegister) (err error) {
	return
}