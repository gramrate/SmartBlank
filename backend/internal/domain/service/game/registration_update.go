package game

import (
	"backend/internal/domain/dto"
	"backend/internal/domain/entity"
	"context"
	"fmt"
)

func (s *Service) UpdateRegistrationPlayer(ctx context.Context, req *dto.UpdateRegistrationPlayerRequest) (*entity.GameState, error) {
	return s.updateGame(ctx, req.LobbyID, func(game *entity.GameState) error {
		idx := findRegistrationIndex(game.Registration.Players, req.Position)
		if idx == -1 {
			return fmt.Errorf("registration player not found: %d", req.Position)
		}
		if req.Name != "" {
			game.Registration.Players[idx].Name = req.Name
		}
		if req.NewPosition > 0 && req.NewPosition != req.Position {
			if hasRegistrationPosition(game.Registration.Players, req.NewPosition) {
				return fmt.Errorf("position %d already taken", req.NewPosition)
			}
			game.Registration.Players[idx].Position = req.NewPosition
		}
		return nil
	})
}
