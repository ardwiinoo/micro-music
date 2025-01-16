package translator

import "github.com/ardwiinoo/micro-music/musics/internal/commons/exceptions"

var DomainErrorMapping = map[string]exceptions.AppError{
	"DETAIL_SONG.NOT_CONTAIN_NEEDED_PROPERTY": exceptions.InvariantError("Detail song not contain needed property"),
	"ADD_SONG.NOT_CONTAIN_NEEDED_PROPERTY":    exceptions.InvariantError("Add song not contain needed property"),
	"DETAIL_USER.NOT_CONTAIN_NEEDED_PROPERTY": exceptions.InvariantError("Detail user not contain needed property"),
	"ADD_SONG_USE_CASE.PUBLIC_ID_NOT_FOUND":   exceptions.UnauthorizedError("Public ID not found"),
	"ADD_SONG_USE_CASE.USER_NOT_FOUND":        exceptions.NotFoundError("User not found"),
	"ADD_SONG_USE_CASE.USER_NOT_AUTHORIZED":   exceptions.UnauthorizedError("User not authorized"),
}