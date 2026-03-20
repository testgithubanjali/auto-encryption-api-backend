package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type File struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID   primitive.ObjectID `bson:"user_id"`
	FileName string             `bson:"file_name"`
	Path     string             `bson:"path"`
}