package game

import (
	"backend/internal/domain/dto"
	"backend/internal/domain/entity"
	"context"
	"errors"
)

func (s *Service) MafiaTarget(ctx context.Context, req *dto.MafiaTargetRequest) (*entity.GameState, error) {
	return s.updateGame(ctx, req.LobbyID, func(game *entity.GameState) error {
		nightIdx := lastNightIndex(game)
		if nightIdx == -1 {
			return errors.New("no night started")
		}
		game.Nights[nightIdx].MafiaTarget = &req.Position
		game.Nights[nightIdx].MafiaMiss = false
		return nil
	})
}
