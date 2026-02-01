package game

import (
	"backend/internal/domain/dto"
	"backend/internal/domain/entity"
	"context"
)

func (s *Service) UpdateMusic(ctx context.Context, req *dto.UpdateMusicRequest) (*entity.GameState, error) {
	return s.updateGame(ctx, req.LobbyID, func(game *entity.GameState) error {
		game.Music.Paused = req.Paused
		if req.Volume >= 0 {
			game.Music.Volume = req.Volume
		}
		return nil
	})
}
