package game

import (
	"backend/internal/domain/dto"
	"backend/internal/domain/entity"
	"context"
	"errors"
	"fmt"
)

func (s *Service) DonCheck(ctx context.Context, req *dto.DonCheckRequest) (*entity.GameState, error) {
	return s.updateGame(ctx, req.LobbyID, func(game *entity.GameState) error {
		nightIdx := lastNightIndex(game)
		if nightIdx == -1 {
			return errors.New("no night started")
		}
		idx := findPlayerIndex(game.Players.Players, req.Position)
		if idx == -1 {
			return fmt.Errorf("player not found: %d", req.Position)
		}
		isSheriff := game.Players.Players[idx].Role == entity.RoleSheriff
		game.Nights[nightIdx].DonCheck = &req.Position
		game.Nights[nightIdx].DonCheckResult = &isSheriff
		return nil
	})
}
