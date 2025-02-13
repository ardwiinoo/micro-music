package usecase

import (
	"context"
	"errors"
	"github.com/ardwiinoo/micro-music/musics/internal/applications/service"
	"mime/multipart"
	"net/http"

	"github.com/google/uuid"

	"github.com/ardwiinoo/micro-music/musics/internal/commons/constants"
	"github.com/ardwiinoo/micro-music/musics/internal/domains/songs"
	"github.com/ardwiinoo/micro-music/musics/internal/domains/songs/entities"
	"github.com/ardwiinoo/micro-music/musics/internal/domains/users"
)

type AddSongUseCase interface {
	Execute(ctx context.Context, payload entities.AddSong, songFile *multipart.FileHeader) (id string, err error)
}

type addSongUseCase struct {
	songRepository  songs.SongRepository
	userRepository  users.UserRepository
	firebaseStorage service.FirebaseService
}

func NewAddSongUseCase(songRepository songs.SongRepository, userRepository users.UserRepository, firebaseStorage service.FirebaseService) AddSongUseCase {
	return &addSongUseCase{
		songRepository:  songRepository,
		userRepository:  userRepository,
		firebaseStorage: firebaseStorage,
	}
}

func (a *addSongUseCase) Execute(ctx context.Context, payload entities.AddSong, songFile *multipart.FileHeader) (id string, err error) {

	publicID, ok := ctx.Value(constants.PublicIDContextKey).(uuid.UUID)
	if !ok {
		return "", errors.New("ADD_SONG_USE_CASE.PUBLIC_ID_NOT_FOUND")
	}

	user, err := a.userRepository.GetDetailUserByPublicId(ctx, publicID.String())
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

	file, err := songFile.Open()
	if err != nil {
		return "", err
	}
	defer file.Close()

	buffer := make([]byte, 512)
	_, err = file.Read(buffer)
	if err != nil {
		return "", err
	}

	contentType := http.DetectContentType(buffer)
	if contentType != "audio/mpeg" && contentType != "audio/mp3" && contentType != "audio/aac" {
		return "", errors.New("ADD_SONG_USE_CASE.INVALID_FILE_TYPE")
	}

	file.Seek(0, 0)

	publicUrl, err := a.firebaseStorage.Upload(ctx, songFile.Filename, file)
	if err != nil {
		return "", err
	}

	payload.Url = publicUrl

	id, err = a.songRepository.AddSong(ctx, payload)
	if err != nil {
		return "", err
	}

	return id, nil
}
