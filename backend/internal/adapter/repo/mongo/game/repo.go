package game

import "go.mongodb.org/mongo-driver/mongo"

type GameRepository struct {
	collection *mongo.Collection
}
