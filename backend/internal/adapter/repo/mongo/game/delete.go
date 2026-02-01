package game

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

func (r *GameRepository) DeleteGame(ctx context.Context, lobbyID string) error {
	_, err := r.collection.DeleteOne(ctx, bson.M{"lobby_id": lobbyID})
	return err
}
