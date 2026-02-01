package game

import (
	"backend/internal/domain/dto"
	"backend/internal/domain/entity"
	"context"
)

func (s *Service) EndGame(ctx context.Context, req *dto.EndGameRequest) (*entity.GameState, error) {
	return s.updateGame(ctx, req.LobbyID, func(game *entity.GameState) error {
		game.End = &entity.EndState{
			Winner:  req.Winner,
			Players: req.Players,
		}
		game.IsActive = false
		game.StageType = entity.StageEnd
		return nil
	})
}
