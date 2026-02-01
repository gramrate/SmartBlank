package game

import (
	"backend/internal/domain/entity"
	"context"
)

type gameRepo interface {
	CreateGame(ctx context.Context, game *entity.GameState) (string, error)
	GetGameByLobbyID(ctx context.Context, lobbyID string) (*entity.GameState, error)
	UpdateGame(ctx context.Context, lobbyID string, game *entity.GameState) error
}

type Service struct {
	repo gameRepo
}

func NewService(repo gameRepo) *Service {
	return &Service{repo: repo}
}
