package game

import (
	"backend/internal/domain/dto"
	"backend/internal/domain/entity"
	"context"
	"errors"
)

func (s *Service) KickAllVote(ctx context.Context, req *dto.KickAllVoteRequest) (*entity.GameState, error) {
	return s.updateGame(ctx, req.LobbyID, func(game *entity.GameState) error {
		dayIdx := lastDayIndex(game)
		if dayIdx == -1 {
			return errors.New("no day started")
		}
		day := &game.Days[dayIdx]
		if len(day.KickAllCandidates) == 0 {
			return errors.New("no kick-all candidates")
		}
		day.KickAllVotes = req.Votes
		if req.Votes > game.Players.AlivePlayers/2 {
			kickPlayers(game, day.KickAllCandidates)
			day.ReVote.KickedPlayers = day.KickAllCandidates
			day.TodayHaveVote = false
		}
		return nil
	})
}
