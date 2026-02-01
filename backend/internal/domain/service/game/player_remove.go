package game

import (
	"backend/internal/domain/dto"
	"backend/internal/domain/entity"
	"context"
	"fmt"
)

func (s *Service) RemovePlayer(ctx context.Context, req *dto.RemovePlayerRequest) (*entity.GameState, error) {
	return s.updateGame(ctx, req.LobbyID, func(game *entity.GameState) error {
		idx := findPlayerIndex(game.Players.Players, req.Position)
		if idx == -1 {
			return fmt.Errorf("player not found: %d", req.Position)
		}
		game.Players.Players[idx].IsDisqualified = true
		game.Players.Players[idx].IsAlive = false
		game.Players.Players[idx].PendingDeath = false
		recountAlive(game)
		return nil
	})
}
