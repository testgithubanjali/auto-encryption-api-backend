package services

import (
	"auto-encryption-api-backend/database"
	"auto-encryption-api-backend/models"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetSessionByID(sessionID string) (*models.Session, error) {

	var session models.Session

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := database.SessionCollection.FindOne(ctx, bson.M{
		"session_id": sessionID,
	}).Decode(&session)

	if err != nil {
		return nil, err
	}

	return &session, nil
}
func CreateSession(userID string, sessionID string, refreshToken string) error {

	objID, _ := primitive.ObjectIDFromHex(userID)

	session := models.Session{
		UserID:       objID,
		SessionID:    sessionID,
		RefreshToken: refreshToken,
		IsExpired:    false,
	}

	_, err := database.SessionCollection.InsertOne(context.TODO(), session)
	return err
}
func ExpireSession(refreshToken string) error {

	_, err := database.SessionCollection.UpdateOne(
		context.TODO(),
		bson.M{"refresh_token": refreshToken},
		bson.M{"$set": bson.M{"is_expired": true}},
	)

	return err
}
