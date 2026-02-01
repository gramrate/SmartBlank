package game

import (
	"backend/internal/domain/dto"
	"backend/internal/domain/entity"
	"context"
)

func (s *Service) StartNight(ctx context.Context, req *dto.StartNightRequest) (*entity.GameState, error) {
	return s.updateGame(ctx, req.LobbyID, func(game *entity.GameState) error {
		game.NightNumber++
		game.StageType = entity.StageNight
		game.Nights = append(game.Nights, entity.NightState{Number: game.NightNumber})
		return nil
	})
}
