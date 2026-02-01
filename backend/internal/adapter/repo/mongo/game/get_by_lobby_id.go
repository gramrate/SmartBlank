package game

import (
	"backend/internal/domain/entity"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r *GameRepository) GetGameByLobbyID(ctx context.Context, lobbyID string) (*entity.GameState, error) {
	var game entity.GameState
	err := r.collection.FindOne(ctx, bson.M{"lobby_id": lobbyID}).Decode(&game)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("game not found")
		}
		return nil, fmt.Errorf("failed to get game: %w", err)
	}
	return &game, nil
}
