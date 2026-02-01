package game

import (
	"backend/internal/domain/entity"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *GameRepository) CreateGame(ctx context.Context, game *entity.GameState) (string, error) {
	now := time.Now()
	game.CreatedAt = now
	game.UpdatedAt = now

	result, err := r.collection.InsertOne(ctx, game)
	if err != nil {
		return "", fmt.Errorf("failed to create game: %w", err)
	}

	id, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return "", fmt.Errorf("failed to read inserted id")
	}

	return id.Hex(), nil
}
