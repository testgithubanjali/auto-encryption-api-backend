package services

import (
	"context"
	"time"

	"auto-encryption-api-backend/database"
	"auto-encryption-api-backend/models"
)

func CreateKey(key models.Key) error {

	key.CreatedAt = time.Now()

	_, err := database.KeyCollection.InsertOne(context.TODO(), key)

	return err
}