package game

import (
	"backend/internal/domain/entity"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *GameRepository) UpdateGame(ctx context.Context, lobbyID string, game *entity.GameState) error {
	game.UpdatedAt = time.Now()

	_, err := r.collection.ReplaceOne(
		ctx,
		bson.M{"lobby_id": lobbyID},
		game,
		options.Replace().SetUpsert(false),
	)

	return err
}
