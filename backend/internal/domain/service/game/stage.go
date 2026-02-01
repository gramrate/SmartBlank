package game

import (
	"backend/internal/domain/dto"
	"backend/internal/domain/entity"
	"context"
)

func (s *Service) SetStage(ctx context.Context, req *dto.SetStageRequest) (*entity.GameState, error) {
	return s.updateGame(ctx, req.LobbyID, func(game *entity.GameState) error {
		game.StageType = req.StageType
		return nil
	})
}
