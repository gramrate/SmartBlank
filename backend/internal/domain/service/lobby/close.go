package lobby

import (
	"backend/internal/domain/dto"
	"context"
)

func (s *Service) Close(ctx context.Context, req *dto.CloseLobbyRequest) (*dto.CloseLobbyResponse, error) {
	return &dto.CloseLobbyResponse{}, s.lobbyRepo.Close(ctx, req.LobbyID)
}
