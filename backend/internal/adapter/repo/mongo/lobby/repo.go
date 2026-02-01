package lobby

import "go.mongodb.org/mongo-driver/mongo"

type Repo struct {
	collection *mongo.Collection
}
