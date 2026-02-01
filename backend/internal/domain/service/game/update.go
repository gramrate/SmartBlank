package game

import (
	"backend/internal/domain/entity"
	"context"
	"errors"
)

func (s *Service) updateGame(ctx context.Context, lobbyID string, mutate func(*entity.GameState) error) (*entity.GameState, error) {
	if lobbyID == "" {
		return nil, errors.New("lobby_id is required")
	}
	game, err := s.repo.GetGameByLobbyID(ctx, lobbyID)
	if err != nil {
		return nil, err
	}
	if err := mutate(game); err != nil {
		return nil, err
	}
	if err := s.repo.UpdateGame(ctx, lobbyID, game); err != nil {
		return nil, err
	}
	return game, nil
}
