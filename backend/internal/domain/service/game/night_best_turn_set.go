package game

import (
	"backend/internal/domain/dto"
	"backend/internal/domain/entity"
	"context"
	"errors"
	"fmt"
)

func (s *Service) SetBestTurn(ctx context.Context, req *dto.SetBestTurnRequest) (*entity.GameState, error) {
	return s.updateGame(ctx, req.LobbyID, func(game *entity.GameState) error {
		if game.FirstNightKillPos == nil || *game.FirstNightKillPos != req.Position {
			return errors.New("best turn allowed only for first night killed player")
		}
		idx := findPlayerIndex(game.Players.Players, req.Position)
		if idx == -1 {
			return fmt.Errorf("player not found: %d", req.Position)
		}
		game.Players.Players[idx].BestTurn = req.BestTurn
		return nil
	})
}
