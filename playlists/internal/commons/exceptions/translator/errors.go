package translator

import "github.com/ardwiinoo/micro-music/playlists/internal/commons/exceptions"

var DomainErrorMapping = map[string]exceptions.AppError{
	"ADD_PLAYLIST.NOT_CONTAIN_NEEDED_PROPERTY": exceptions.InvariantError("Add playlist not contain needed property"),
	"DETAIL_PLAYLIST.NOT_CONTAIN_NEEDED_PROPERTY": exceptions.InvariantError("Detail playlist not contain needed property"),
}