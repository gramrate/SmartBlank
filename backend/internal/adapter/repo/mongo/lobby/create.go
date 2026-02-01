package lobby

import (
	"backend/internal/domain/entity"
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
)

func (r *Repo) Create(ctx context.Context) (*entity.Lobby, error) {
	lobby := &entity.Lobby{
		ID:        uuid.New(),
		IsActive:  true,
		CreatedAt: time.Now(),
	}
	if _, err := r.collection.InsertOne(ctx, lobby); err != nil {
		return nil, fmt.Errorf("failed to create lobby: %w", err)
	}
	return lobby, nil
}
