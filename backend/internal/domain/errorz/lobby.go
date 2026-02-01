package errorz

import "errors"

var (
	LobbyNotFound      = errors.New("lobby not found")
	LobbyAlreadyClosed = errors.New("lobby already closed")
)
