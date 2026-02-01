package game

import (
	"backend/internal/domain/dto"
	"backend/internal/domain/entity"
	"context"
	"errors"
)

func (s *Service) MafiaMiss(ctx context.Context, req *dto.MafiaMissRequest) (*entity.GameState, error) {
	return s.updateGame(ctx, req.LobbyID, func(game *entity.GameState) error {
		nightIdx := lastNightIndex(game)
		if nightIdx == -1 {
			return errors.New("no night started")
		}
		game.Nights[nightIdx].MafiaTarget = nil
		game.Nights[nightIdx].MafiaMiss = true
		return nil
	})
}
