package game

import (
	"backend/internal/domain/dto"
	"backend/internal/domain/entity"
	"context"
	"errors"
	"fmt"
)

func (s *Service) AddRegistrationPlayer(ctx context.Context, req *dto.AddRegistrationPlayerRequest) (*entity.GameState, error) {
	if req.Name == "" {
		return nil, errors.New("name is required")
	}
	return s.updateGame(ctx, req.LobbyID, func(game *entity.GameState) error {
		position := req.Position
		if position <= 0 {
			position = nextAvailablePosition(game.Registration.Players)
		}
		if hasRegistrationPosition(game.Registration.Players, position) {
			return fmt.Errorf("position %d already taken", position)
		}
		game.Registration.Players = append(game.Registration.Players, entity.RegistrationPlayerState{
			Name:     req.Name,
			Position: position,
		})
		return nil
	})
}
