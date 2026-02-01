package game

import (
	"backend/internal/domain/dto"
	"backend/internal/domain/entity"
	"context"
	"fmt"
)

func (s *Service) RemoveRegistrationPlayer(ctx context.Context, req *dto.RemoveRegistrationPlayerRequest) (*entity.GameState, error) {
	return s.updateGame(ctx, req.LobbyID, func(game *entity.GameState) error {
		idx := findRegistrationIndex(game.Registration.Players, req.Position)
		if idx == -1 {
			return fmt.Errorf("registration player not found: %d", req.Position)
		}
		game.Registration.Players = append(game.Registration.Players[:idx], game.Registration.Players[idx+1:]...)
		return nil
	})
}
