package service_provider

import (
	"backend/internal/domain/dto"
	"backend/internal/domain/entity"
	"backend/internal/domain/service/game"
	"context"
)

type gameService interface {
	Create(ctx context.Context, req *dto.CreateGameRequest) (*entity.GameState, error)
	Get(ctx context.Context, req *dto.LobbyIDRequest) (*entity.GameState, error)
	AddRegistrationPlayer(ctx context.Context, req *dto.AddRegistrationPlayerRequest) (*entity.GameState, error)
	UpdateRegistrationPlayer(ctx context.Context, req *dto.UpdateRegistrationPlayerRequest) (*entity.GameState, error)
	RemoveRegistrationPlayer(ctx context.Context, req *dto.RemoveRegistrationPlayerRequest) (*entity.GameState, error)
	SwapRegistrationPositions(ctx context.Context, req *dto.SwapRegistrationPositionsRequest) (*entity.GameState, error)
	GenerateSeating(ctx context.Context, req *dto.GenerateSeatingRequest) (*entity.GameState, error)
	SetStage(ctx context.Context, req *dto.SetStageRequest) (*entity.GameState, error)
	AssignRole(ctx context.Context, req *dto.AssignRoleRequest) (*entity.GameState, error)
	SetForbiddenRole(ctx context.Context, req *dto.ForbiddenRoleRequest) (*entity.GameState, error)
	AutoDeal(ctx context.Context, req *dto.AutoDealRequest) (*entity.GameState, error)
	AddFoul(ctx context.Context, req *dto.AddFoulRequest) (*entity.GameState, error)
	SetCard(ctx context.Context, req *dto.SetCardRequest) (*entity.GameState, error)
	RemovePlayer(ctx context.Context, req *dto.RemovePlayerRequest) (*entity.GameState, error)
	StartDay(ctx context.Context, req *dto.StartDayRequest) (*entity.GameState, error)
	StartNight(ctx context.Context, req *dto.StartNightRequest) (*entity.GameState, error)
	StartVote(ctx context.Context, req *dto.StartVoteRequest) (*entity.GameState, error)
	SetVote(ctx context.Context, req *dto.SetVoteRequest) (*entity.GameState, error)
	ResolveVote(ctx context.Context, req *dto.ResolveVoteRequest) (*entity.GameState, error)
	KickAllVote(ctx context.Context, req *dto.KickAllVoteRequest) (*entity.GameState, error)
	MafiaTarget(ctx context.Context, req *dto.MafiaTargetRequest) (*entity.GameState, error)
	MafiaMiss(ctx context.Context, req *dto.MafiaMissRequest) (*entity.GameState, error)
	SheriffCheck(ctx context.Context, req *dto.SheriffCheckRequest) (*entity.GameState, error)
	DonCheck(ctx context.Context, req *dto.DonCheckRequest) (*entity.GameState, error)
	ApplyNightResults(ctx context.Context, req *dto.ApplyNightResultsRequest) (*entity.GameState, error)
	SetBestTurn(ctx context.Context, req *dto.SetBestTurnRequest) (*entity.GameState, error)
	UpdateMusic(ctx context.Context, req *dto.UpdateMusicRequest) (*entity.GameState, error)
	EndGame(ctx context.Context, req *dto.EndGameRequest) (*entity.GameState, error)
}

func (s *ServiceProvider) GameService() gameService {
	if s.gameService == nil {
		s.gameService = game.NewService(s.GameRepository())
	}
	return s.gameService
}
