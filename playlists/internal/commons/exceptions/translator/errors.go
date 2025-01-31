package translator

import "github.com/ardwiinoo/micro-music/playlists/internal/commons/exceptions"

var DomainErrorMapping = map[string]exceptions.AppError{
	"ADD_PLAYLIST.NOT_CONTAIN_NEEDED_PROPERTY":      exceptions.InvariantError("Add playlist not contain needed property"),
	"ADD_PLAYLIST_USE_CASE.PUBLIC_ID_NOT_FOUND":     exceptions.InvariantError("Public id not found"),
	"ADD_PLAYLIST_USE_CASE.USER_NOT_FOUND":          exceptions.NotFoundError("User not found"),
	"ADD_PLAYLIST_USE_CASE.USER_NOT_AUTHORIZED":     exceptions.ForbiddenError("User not authorized"),
	"DETAIL_PLAYLIST.NOT_CONTAIN_NEEDED_PROPERTY":   exceptions.InvariantError("Detail playlist not contain needed property"),
	"DELETE_PLAYLIST_USE_CASE.PUBLIC_ID_NOT_FOUND":  exceptions.InvariantError("Public id not found"),
	"DELETE_PLAYLIST_USE_CASE.USER_NOT_FOUND":       exceptions.NotFoundError("User not found"),
	"DELETE_PLAYLIST_USE_CASE.USER_NOT_AUTHORIZED":  exceptions.ForbiddenError("User not authorized"),
	"DELETE_PLAYLIST_USE_CASE.PLAYLIST_NOT_FOUND":   exceptions.NotFoundError("Playlist not found"),
	"ADD_PLAYLIST_SONG.NOT_CONTAIN_NEEDED_PROPERTY": exceptions.InvariantError("Add playlist song not contain needed property"),
}
