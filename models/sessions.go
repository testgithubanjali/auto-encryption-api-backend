package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Session struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	UserID       primitive.ObjectID `bson:"user_id"`
	SessionID    string             `bson:"session_id"`
	RefreshToken string             `bson:"refresh_token"`
	IsExpired    bool               `bson:"is_expired"`
	CreatedAt    time.Time          `bson:"created_at"`
}
