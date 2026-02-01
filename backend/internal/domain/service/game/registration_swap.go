package game

import (
	"backend/internal/domain/dto"
	"backend/internal/domain/entity"
	"context"
	"errors"
)

func (s *Service) SwapRegistrationPositions(ctx context.Context, req *dto.SwapRegistrationPositionsRequest) (*entity.GameState, error) {
	return s.updateGame(ctx, req.LobbyID, func(game *entity.GameState) error {
		idxA := findRegistrationIndex(game.Registration.Players, req.PositionA)
		idxB := findRegistrationIndex(game.Registration.Players, req.PositionB)
		if idxA == -1 || idxB == -1 {
			return errors.New("registration player not found")
		}
		game.Registration.Players[idxA].Position, game.Registration.Players[idxB].Position =
			game.Registration.Players[idxB].Position, game.Registration.Players[idxA].Position
		return nil
	})
}
