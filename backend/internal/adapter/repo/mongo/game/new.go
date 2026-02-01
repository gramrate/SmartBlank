package game

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewGameRepository(db *mongo.Database) *GameRepository {
	collection := db.Collection("games")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, _ = collection.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys:    bson.D{{Key: "lobby_id", Value: 1}},
		Options: options.Index().SetUnique(true),
	})

	_, _ = collection.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys: bson.D{
			{Key: "is_active", Value: 1},
			{Key: "updated_at", Value: -1},
		},
	})

	return &GameRepository{collection: collection}
}
