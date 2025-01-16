package usecase

import (
	"context"
	"errors"

	"github.com/ardwiinoo/micro-music/musics/internal/commons/constants"
	"github.com/ardwiinoo/micro-music/musics/internal/domains/songs"
	"github.com/ardwiinoo/micro-music/musics/internal/domains/songs/entities"
	"github.com/ardwiinoo/micro-music/musics/internal/domains/users"

)

type AddSongUseCase interface {
	Execute(ctx context.Context, payload entities.AddSong) (id string, err error)
}

type addSongUseCase struct {
	songRepository songs.SongRepository
	userRepository users.UserRepository
}

func NewAddSongUseCase(songRepository songs.SongRepository, userRepository users.UserRepository) AddSongUseCase {
	return &addSongUseCase{
		songRepository: songRepository,
		userRepository: userRepository,
	}
}

func (a *addSongUseCase) Execute(ctx context.Context, payload entities.AddSong) (id string, err error) {

	publicID, ok := ctx.Value(constants.PublicIDContextKey).(string)
	if !ok {
		return "", errors.New("ADD_SONG_USE_CASE.PUBLIC_ID_NOT_FOUND")
	}

	user, err := a.userRepository.GetDetailUserByPublicId(ctx, publicID)
	if err != nil {
		return "", errors.New("ADD_SONG_USE_CASE.USER_NOT_FOUND")
	}

	if user.Role != constants.AdminRoleID {
		return "", errors.New("ADD_SONG_USE_CASE.USER_NOT_AUTHORIZED")
	}

	err = payload.Validate()
	if err != nil {
		return "", err
	}
	
	id, err = a.songRepository.AddSong(ctx, payload)
	if err != nil {
		return "", err
	}

	return id, nil
}