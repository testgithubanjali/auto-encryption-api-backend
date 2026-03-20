package services

import (
	"auto-encryption-api-backend/database"
	"auto-encryption-api-backend/models"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateUser(user models.User) error {
	_, err := database.UserCollection.InsertOne(context.TODO(), user)
	return err
}

func GetUserByEmail(email string) (*models.User, error) {

	filter := bson.M{
		"email": email,
	}

	var user models.User

	err := database.UserCollection.FindOne(context.TODO(), filter).Decode(&user)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
func GetUserByID(id primitive.ObjectID) (*models.User, error) {

	var user models.User

	filter := bson.M{"_id": id}

	err := database.UserCollection.FindOne(context.TODO(), filter).Decode(&user)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
