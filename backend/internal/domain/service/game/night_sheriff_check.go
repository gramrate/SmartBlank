package game

import (
	"backend/internal/domain/dto"
	"backend/internal/domain/entity"
	"context"
	"errors"
	"fmt"
)

func (s *Service) SheriffCheck(ctx context.Context, req *dto.SheriffCheckRequest) (*entity.GameState, error) {
	return s.updateGame(ctx, req.LobbyID, func(game *entity.GameState) error {
		nightIdx := lastNightIndex(game)
		if nightIdx == -1 {
			return errors.New("no night started")
		}
		idx := findPlayerIndex(game.Players.Players, req.Position)
		if idx == -1 {
			return fmt.Errorf("player not found: %d", req.Position)
		}
		isMafia := game.Players.Players[idx].Role == entity.RoleMafia || game.Players.Players[idx].Role == entity.RoleDon
		game.Nights[nightIdx].SheriffCheck = &req.Position
		game.Nights[nightIdx].SheriffCheckResult = &isMafia
		return nil
	})
}
