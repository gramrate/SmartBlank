package dto

import "github.com/google/uuid"

type AddPlayerRequest struct {
	Name    string `json:"name"`
	LobbyID string `json:"lobby_id"`
}

type CreateLobbyRequest struct{}
type CreateLobbyResponse struct {
	LobbyID uuid.UUID `json:"lobby_id"`
}

type CloseLobbyRequest struct {
	LobbyID uuid.UUID `json:"lobby_id"`
}
type CloseLobbyResponse struct{}
