package dto

import "backend/internal/domain/entity"

type LobbyIDRequest struct {
	LobbyID string `json:"lobby_id"`
}

type CreateGameRequest struct {
	LobbyID   string `json:"lobby_id"`
	LobbyName string `json:"lobby_name"`
}

type AddRegistrationPlayerRequest struct {
	LobbyID  string `json:"lobby_id"`
	Name     string `json:"name"`
	Position int    `json:"position"`
}

type UpdateRegistrationPlayerRequest struct {
	LobbyID    string `json:"lobby_id"`
	Position   int    `json:"position"`
	Name       string `json:"name"`
	NewPosition int   `json:"new_position"`
}

type RemoveRegistrationPlayerRequest struct {
	LobbyID  string `json:"lobby_id"`
	Position int    `json:"position"`
}

type SwapRegistrationPositionsRequest struct {
	LobbyID   string `json:"lobby_id"`
	PositionA int    `json:"position_a"`
	PositionB int    `json:"position_b"`
}

type GenerateSeatingRequest struct {
	LobbyID string `json:"lobby_id"`
}

type SetStageRequest struct {
	LobbyID  string           `json:"lobby_id"`
	StageType entity.StageType `json:"stage_type"`
}

type AssignRoleRequest struct {
	LobbyID  string       `json:"lobby_id"`
	Position int          `json:"position"`
	Role     entity.Role  `json:"role"`
}

type ForbiddenRoleRequest struct {
	LobbyID  string       `json:"lobby_id"`
	Position int          `json:"position"`
	Roles    []entity.Role `json:"roles"`
}

type AutoDealRequest struct {
	LobbyID        string `json:"lobby_id"`
	MafiaCount     int    `json:"mafia_count"`
	IncludeSheriff bool   `json:"include_sheriff"`
	IncludeDon     bool   `json:"include_don"`
}

type AddFoulRequest struct {
	LobbyID  string `json:"lobby_id"`
	Position int    `json:"position"`
	Delta    int    `json:"delta"`
}

type SetCardRequest struct {
	LobbyID  string           `json:"lobby_id"`
	Position int              `json:"position"`
	Card     entity.CardColor `json:"card"`
}

type RemovePlayerRequest struct {
	LobbyID  string `json:"lobby_id"`
	Position int    `json:"position"`
}

type StartVoteRequest struct {
	LobbyID           string `json:"lobby_id"`
	SuggestedPlayers  []int  `json:"suggested_players"`
	IsReVote          bool   `json:"is_revote"`
}

type SetVoteRequest struct {
	LobbyID    string `json:"lobby_id"`
	Position   int    `json:"position"`
	VotesCount int    `json:"votes_count"`
	IsReVote   bool   `json:"is_revote"`
}

type ResolveVoteRequest struct {
	LobbyID  string `json:"lobby_id"`
	IsReVote bool   `json:"is_revote"`
}

type KickAllVoteRequest struct {
	LobbyID string `json:"lobby_id"`
	Votes   int    `json:"votes"`
}

type StartDayRequest struct {
	LobbyID string `json:"lobby_id"`
}

type StartNightRequest struct {
	LobbyID string `json:"lobby_id"`
}

type MafiaTargetRequest struct {
	LobbyID  string `json:"lobby_id"`
	Position int    `json:"position"`
}

type MafiaMissRequest struct {
	LobbyID string `json:"lobby_id"`
}

type SheriffCheckRequest struct {
	LobbyID  string `json:"lobby_id"`
	Position int    `json:"position"`
}

type DonCheckRequest struct {
	LobbyID  string `json:"lobby_id"`
	Position int    `json:"position"`
}

type ApplyNightResultsRequest struct {
	LobbyID string `json:"lobby_id"`
}

type SetBestTurnRequest struct {
	LobbyID  string `json:"lobby_id"`
	Position int    `json:"position"`
	BestTurn []int  `json:"best_turn"`
}

type UpdateMusicRequest struct {
	LobbyID string  `json:"lobby_id"`
	Paused  bool    `json:"paused"`
	Volume  float64 `json:"volume"`
}

type EndGameRequest struct {
	LobbyID string                `json:"lobby_id"`
	Winner  entity.Team           `json:"winner"`
	Players []entity.PlayerEndState `json:"players"`
}
