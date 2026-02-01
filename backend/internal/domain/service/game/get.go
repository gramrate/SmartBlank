package game

import (
	"backend/internal/domain/dto"
	"backend/internal/domain/entity"
	"context"
)

func (s *Service) Get(ctx context.Context, req *dto.LobbyIDRequest) (*entity.GameState, error) {
	return s.repo.GetGameByLobbyID(ctx, req.LobbyID)
}
