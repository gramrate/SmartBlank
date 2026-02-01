package game

import (
	"backend/internal/domain/dto"
	"backend/internal/domain/entity"
	"context"
	"fmt"
)

func (s *Service) AddFoul(ctx context.Context, req *dto.AddFoulRequest) (*entity.GameState, error) {
	return s.updateGame(ctx, req.LobbyID, func(game *entity.GameState) error {
		idx := findPlayerIndex(game.Players.Players, req.Position)
		if idx == -1 {
			return fmt.Errorf("player not found: %d", req.Position)
		}
		delta := req.Delta
		if delta == 0 {
			delta = 1
		}
		game.Players.Players[idx].Fouls += delta
		if game.Players.Players[idx].Fouls < 0 {
			game.Players.Players[idx].Fouls = 0
		}
		return nil
	})
}
