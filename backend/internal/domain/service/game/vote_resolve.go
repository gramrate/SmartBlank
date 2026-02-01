package game

import (
	"backend/internal/domain/dto"
	"backend/internal/domain/entity"
	"context"
	"errors"
)

func (s *Service) ResolveVote(ctx context.Context, req *dto.ResolveVoteRequest) (*entity.GameState, error) {
	return s.updateGame(ctx, req.LobbyID, func(game *entity.GameState) error {
		dayIdx := lastDayIndex(game)
		if dayIdx == -1 {
			return errors.New("no day started")
		}
		day := &game.Days[dayIdx]
		vote := &day.Vote
		if req.IsReVote {
			vote = &day.ReVote
		}
		if len(vote.SuggestedPlayers) == 0 {
			return errors.New("no vote candidates")
		}
		maxVotes, tied := resolveVoteTies(vote)
		if maxVotes == 0 && len(tied) == 0 {
			return errors.New("no votes to resolve")
		}
		if len(tied) == 1 {
			kickPlayers(game, tied)
			vote.KickedPlayers = tied
			day.TodayHaveVote = false
			return nil
		}

		if !req.IsReVote {
			day.ReVote = entity.VoteStateState{
				SuggestedPlayers: tied,
				Voting:           initVoteCounts(tied),
			}
			return nil
		}

		if len(tied) < len(day.Vote.SuggestedPlayers) {
			day.KickAllCandidates = tied
			day.KickAllVotes = 0
			return nil
		}

		day.KickAllCandidates = tied
		day.KickAllVotes = 0
		return nil
	})
}
