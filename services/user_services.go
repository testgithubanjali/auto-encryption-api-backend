package services

import (
	"auto-encryption-api-backend/database"
	"auto-encryption-api-backend/models"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateUser(user models.User) error {
	log.Println("Inserting user into database")
	result, err := database.UserCollection.InsertOne(context.TODO(), user)
	if err != nil {
		log.Println("CreateUser: insert failed", err)
		return err
	}
	log.Printf("CreateUser: user created successfully with id:  %v", result.InsertedID)
	return nil
}

func GetUserByEmail(email string) (*models.User, error) {
	log.Printf("GetUserByEmail: get user by email: %s", email)
	filter := bson.M{
		"email": email,
	}

	var user models.User

	err := database.UserCollection.FindOne(context.TODO(), filter).Decode(&user)

	if err != nil {
		log.Println("GetUserByEmail: user not found or error: ", err)
		return nil, err
	}
	log.Printf("user found with ID: %s", user.ID.Hex())
	return &user, nil
}

func GetUserByID(id primitive.ObjectID) (*models.User, error) {
	log.Printf("GetUserByID: Searching user by id: %s", id.Hex())
	var user models.User

	filter := bson.M{"_id": id}

	err := database.UserCollection.FindOne(context.TODO(), filter).Decode(&user)

	if err != nil {
		log.Println("GetUserByID: user not found or error", err)
		return nil, err
	}
	log.Printf("GetUserByID: userfound with email: %s", user.Email)
	return &user, nil
}
