package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Key struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID    primitive.ObjectID `bson:"user_id" json:"user_id"`
	Algorithm string             `bson:"algorithm" json:"algorithm"`
	KeyValue  string             `bson:"key_value" json:"key_value"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
}