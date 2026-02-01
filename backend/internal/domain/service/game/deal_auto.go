package game

import (
	"backend/internal/domain/dto"
	"backend/internal/domain/entity"
	"context"
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func (s *Service) AutoDeal(ctx context.Context, req *dto.AutoDealRequest) (*entity.GameState, error) {
	return s.updateGame(ctx, req.LobbyID, func(game *entity.GameState) error {
		if len(game.Players.Players) == 0 {
			return errors.New("no players to deal")
		}

		for i := range game.Players.Players {
			game.Players.Players[i].Role = entity.RoleCivilian
		}

		for _, forced := range game.Deal.ForcedRoles {
			if forced.Role == entity.RoleCivilian {
				continue
			}
			idx := findPlayerIndex(game.Players.Players, forced.Position)
			if idx == -1 {
				return fmt.Errorf("forced player not found: %d", forced.Position)
			}
			if isRoleForbidden(game.Deal.ForbiddenRoles, forced.Position, forced.Role) {
				return fmt.Errorf("forced role forbidden for position %d", forced.Position)
			}
			game.Players.Players[idx].Role = forced.Role
		}

		roles := buildRoles(req.MafiaCount, req.IncludeSheriff, req.IncludeDon)
		roles = removeForcedRoles(roles, game.Deal.ForcedRoles)

		rng := rand.New(rand.NewSource(time.Now().UnixNano()))
		for _, role := range roles {
			candidates := availableForRole(game.Players.Players, game.Deal.ForbiddenRoles, role)
			if len(candidates) == 0 {
				return fmt.Errorf("no candidates for role %d", role)
			}
			pick := candidates[rng.Intn(len(candidates))]
			game.Players.Players[pick].Role = role
		}
		return nil
	})
}
