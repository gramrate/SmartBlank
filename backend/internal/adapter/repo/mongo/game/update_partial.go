package game

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func (r *GameRepository) UpdateGamePartial(ctx context.Context, lobbyID string, updates bson.M) error {
	updates["updated_at"] = time.Now()

	_, err := r.collection.UpdateOne(
		ctx,
		bson.M{"lobby_id": lobbyID},
		bson.M{"$set": updates},
	)

	return err
}
