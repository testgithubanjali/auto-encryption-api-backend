package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)
type Note struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	UserID     primitive.ObjectID `bson:"user_id" json:"user_id,omitempty"`
	Ciphertext string             `bson:"ciphertext" json:"ciphertext"`
	CreatedAt  time.Time          `bson:"created_at" json:"created_at"`
}