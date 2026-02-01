package game

import (
	"backend/internal/domain/dto"
	"backend/internal/domain/entity"
	"context"
	"fmt"
)

func (s *Service) AssignRole(ctx context.Context, req *dto.AssignRoleRequest) (*entity.GameState, error) {
	return s.updateGame(ctx, req.LobbyID, func(game *entity.GameState) error {
		idx := findPlayerIndex(game.Players.Players, req.Position)
		if idx == -1 {
			return fmt.Errorf("player not found: %d", req.Position)
		}
		game.Players.Players[idx].Role = req.Role
		upsertForcedRole(&game.Deal, req.Position, req.Role)
		return nil
	})
}
