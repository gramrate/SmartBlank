package entity

import (
	"time"

	"github.com/google/uuid"
)

type Lobby struct {
	ID        uuid.UUID `bson:"_id" json:"id"`
	Name      string    `bson:"name,omitempty" json:"name,omitempty"`
	IsActive  bool      `bson:"is_active" json:"is_active"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	ClosedAt  *time.Time `bson:"closed_at,omitempty" json:"closed_at,omitempty"`
}
