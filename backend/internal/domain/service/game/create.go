package game

import (
	"backend/internal/domain/dto"
	"backend/internal/domain/entity"
	"context"

	"github.com/google/uuid"
)

func (s *Service) Create(ctx context.Context, req *dto.CreateGameRequest) (*entity.GameState, error) {
	lobbyID := req.LobbyID
	if lobbyID == "" {
		lobbyID = uuid.NewString()
	}
	game := &entity.GameState{
		LobbyID:    lobbyID,
		LobbyName:  req.LobbyName,
		IsActive:   true,
		StageType:  entity.StageStart,
		Players:    entity.PlayersState{AlivePlayers: 0, Players: []entity.PlayerState{}},
		Registration: entity.RegistrationState{Players: []entity.RegistrationPlayerState{}},
		Days:       []entity.DayState{},
		Nights:     []entity.NightState{},
		DayNumber:  0,
		NightNumber: 0,
		Deal:       entity.DealState{},
		Music:      entity.NightMusicState{Paused: false, Volume: 1.0},
	}

	if _, err := s.repo.CreateGame(ctx, game); err != nil {
		return nil, err
	}

	return game, nil
}
