package game

import (
	"backend/internal/domain/dto"
	"backend/internal/domain/entity"
	"context"
	"errors"
	"math/rand"
	"time"
)

func (s *Service) GenerateSeating(ctx context.Context, req *dto.GenerateSeatingRequest) (*entity.GameState, error) {
	return s.updateGame(ctx, req.LobbyID, func(game *entity.GameState) error {
		if len(game.Registration.Players) == 0 {
			return errors.New("no registration players")
		}

		players := make([]entity.RegistrationPlayerState, len(game.Registration.Players))
		copy(players, game.Registration.Players)

		rng := rand.New(rand.NewSource(time.Now().UnixNano()))
		rng.Shuffle(len(players), func(i, j int) {
			players[i], players[j] = players[j], players[i]
		})

		for i := range players {
			players[i].Position = i + 1
		}
		game.Registration.Players = players

		gamePlayers := make([]entity.PlayerState, 0, len(players))
		for _, p := range players {
			gamePlayers = append(gamePlayers, entity.PlayerState{
				Name:           p.Name,
				Position:       p.Position,
				Fouls:          0,
				IsAlive:        true,
				IsDisqualified: false,
				Role:           entity.RoleCivilian,
				Card:           entity.CardNone,
				PendingDeath:   false,
			})
		}

		game.Players = entity.PlayersState{
			AlivePlayers: len(gamePlayers),
			Players:      gamePlayers,
		}
		game.StageType = entity.StageDeal
		return nil
	})
}
