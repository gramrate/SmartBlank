package lobby

import (
	"backend/internal/domain/entity"
	"context"

	"github.com/google/uuid"
)

type lobbyRepo interface {
	Create(ctx context.Context) (*entity.Lobby, error)
	Close(ctx context.Context, id uuid.UUID) error
}

type Service struct {
	lobbyRepo lobbyRepo
}

func NewService(repo lobbyRepo) *Service {
	return &Service{
		lobbyRepo: repo}
}
