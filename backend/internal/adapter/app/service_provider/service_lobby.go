package service_provider

import (
	"backend/internal/domain/dto"
	"backend/internal/domain/service/lobby"
	"context"
)

type lobbyService interface {
	Close(ctx context.Context, req *dto.CloseLobbyRequest) (*dto.CloseLobbyResponse, error)
	Create(ctx context.Context, req *dto.CreateLobbyRequest) (*dto.CreateLobbyResponse, error)
}

func (s *ServiceProvider) LobbyService() lobbyService {
	if s.lobbyService == nil {
		s.lobbyService = lobby.NewService(s.LobbyRepository())
	}

	return s.lobbyService
}
