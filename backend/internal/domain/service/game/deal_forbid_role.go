package game

import (
	"backend/internal/domain/dto"
	"backend/internal/domain/entity"
	"context"
)

func (s *Service) SetForbiddenRole(ctx context.Context, req *dto.ForbiddenRoleRequest) (*entity.GameState, error) {
	return s.updateGame(ctx, req.LobbyID, func(game *entity.GameState) error {
		upsertForbiddenRole(&game.Deal, req.Position, req.Roles)
		return nil
	})
}
