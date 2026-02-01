package lobby

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

func (r *Repo) Close(ctx context.Context, id uuid.UUID) error {
	update := bson.M{
		"$set": bson.M{
			"is_active": false,
			"closed_at": time.Now(),
		},
	}
	result, err := r.collection.UpdateOne(ctx, bson.M{"_id": id}, update)
	if err != nil {
		return fmt.Errorf("failed to close lobby: %w", err)
	}
	if result.MatchedCount == 0 {
		return fmt.Errorf("lobby not found")
	}
	return nil
}
