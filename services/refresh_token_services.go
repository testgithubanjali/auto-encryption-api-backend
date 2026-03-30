package services

import (
	"context"
	"time"

	"auto-encryption-api-backend/database"
	"auto-encryption-api-backend/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SaveRefreshToken(userID string, token string) error {

	objID, _ := primitive.ObjectIDFromHex(userID)

	rt := models.RefreshToken{
		UserID: objID,
		Token:  token,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := database.RefreshTokenCollection.InsertOne(ctx, rt)
	return err
}
func DeleteRefreshToken(token string) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := database.RefreshTokenCollection.DeleteOne(ctx, map[string]interface{}{
		"token": token,
	})

	return err
}
