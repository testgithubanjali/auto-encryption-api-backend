package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type RefreshToken struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	UserID primitive.ObjectID `bson:"user_id"`
	Token  string             `bson:"token"`
}
