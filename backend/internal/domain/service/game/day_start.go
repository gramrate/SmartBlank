package game

import (
	"backend/internal/domain/dto"
	"backend/internal/domain/entity"
	"context"
)

func (s *Service) StartDay(ctx context.Context, req *dto.StartDayRequest) (*entity.GameState, error) {
	return s.updateGame(ctx, req.LobbyID, func(game *entity.GameState) error {
		game.DayNumber++
		game.StageType = entity.StageDay
		game.Days = append(game.Days, entity.DayState{Number: game.DayNumber})
		return nil
	})
}
