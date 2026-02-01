package lobby

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewRepo(db *mongo.Database) *Repo {
	collection := db.Collection("lobbies")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, _ = collection.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys:    bson.D{{Key: "is_active", Value: 1}, {Key: "created_at", Value: -1}},
		Options: options.Index(),
	})

	return &Repo{collection: collection}
}
