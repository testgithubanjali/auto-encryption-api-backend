package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Entry struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	UserID        primitive.ObjectID `bson:"user_id"`
	KeyID         primitive.ObjectID `bson:"key_id"` 
	Type          string             `bson:"type"`
	OriginalText  string             `bson:"original_text"`
	ProcessedText string             `bson:"processed_text"`
	CreatedAt     time.Time          `bson:"created_at"`
}
