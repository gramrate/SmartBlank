package lobby

import (
	"backend/internal/domain/dto"
	"context"
)

func (s *Service) Create(ctx context.Context, _ *dto.CreateLobbyRequest) (*dto.CreateLobbyResponse, error) {
	l, err := s.lobbyRepo.Create(ctx)
	if err != nil {
		return nil, err
	}
	return &dto.CreateLobbyResponse{
		LobbyID: l.ID,
	}, nil
}
