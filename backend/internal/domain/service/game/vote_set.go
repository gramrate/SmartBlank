package game

import (
	"backend/internal/domain/dto"
	"backend/internal/domain/entity"
	"context"
	"errors"
	"fmt"
)

func (s *Service) SetVote(ctx context.Context, req *dto.SetVoteRequest) (*entity.GameState, error) {
	return s.updateGame(ctx, req.LobbyID, func(game *entity.GameState) error {
		dayIdx := lastDayIndex(game)
		if dayIdx == -1 {
			return errors.New("no day started")
		}
		vote := &game.Days[dayIdx].Vote
		if req.IsReVote {
			vote = &game.Days[dayIdx].ReVote
		}
		if !containsInt(vote.SuggestedPlayers, req.Position) {
			return fmt.Errorf("position %d not in suggested list", req.Position)
		}
		setVoteCount(vote, req.Position, req.VotesCount)
		return nil
	})
}
