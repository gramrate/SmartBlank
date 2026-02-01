package game

import (
	"backend/internal/domain/dto"
	"backend/internal/domain/entity"
	"context"
	"errors"
)

func (s *Service) ApplyNightResults(ctx context.Context, req *dto.ApplyNightResultsRequest) (*entity.GameState, error) {
	return s.updateGame(ctx, req.LobbyID, func(game *entity.GameState) error {
		nightIdx := lastNightIndex(game)
		if nightIdx == -1 {
			return errors.New("no night started")
		}
		night := &game.Nights[nightIdx]
		if night.MafiaMiss || night.MafiaTarget == nil {
			return nil
		}
		kickPlayers(game, []int{*night.MafiaTarget})
		if night.Number == 1 && game.FirstNightKillPos == nil {
			pos := *night.MafiaTarget
			game.FirstNightKillPos = &pos
		}
		return nil
	})
}
