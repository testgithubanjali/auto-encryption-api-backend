package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Entry struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID     primitive.ObjectID `bson:"user_id" json:"user_id"`
	KeyID      primitive.ObjectID `bson:"key_id" json:"key_id"`
	Ciphertext string             `bson:"ciphertext" json:"ciphertext"`
	CreatedAt  time.Time          `bson:"created_at" json:"created_at"`
}