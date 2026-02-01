package game

import (
	"backend/internal/domain/dto"
	"backend/internal/domain/entity"
	"context"
	"errors"
)

func (s *Service) StartVote(ctx context.Context, req *dto.StartVoteRequest) (*entity.GameState, error) {
	return s.updateGame(ctx, req.LobbyID, func(game *entity.GameState) error {
		dayIdx := lastDayIndex(game)
		if dayIdx == -1 {
			return errors.New("no day started")
		}
		vote := entity.VoteStateState{
			SuggestedPlayers: req.SuggestedPlayers,
			Voting:           initVoteCounts(req.SuggestedPlayers),
		}
		if req.IsReVote {
			game.Days[dayIdx].ReVote = vote
		} else {
			game.Days[dayIdx].TodayHaveVote = true
			game.Days[dayIdx].Vote = vote
			game.Days[dayIdx].ReVote = entity.VoteStateState{}
			game.Days[dayIdx].KickAllCandidates = nil
			game.Days[dayIdx].KickAllVotes = 0
		}
		return nil
	})
}
