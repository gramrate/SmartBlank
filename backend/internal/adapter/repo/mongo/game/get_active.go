package game

import (
	"backend/internal/domain/entity"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *GameRepository) GetActiveGames(ctx context.Context) ([]entity.GameState, error) {
	filter := bson.M{"is_active": true}
	opts := options.Find().SetSort(bson.M{"updated_at": -1}).SetLimit(100)

	cursor, err := r.collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var games []entity.GameState
	if err := cursor.All(ctx, &games); err != nil {
		return nil, err
	}

	return games, nil
}
