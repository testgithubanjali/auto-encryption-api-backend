package services

import (
	"context"
	"log"
	"time"

	"auto-encryption-api-backend/database"
	"auto-encryption-api-backend/models"
)

func CreateKey(key models.Key) error {
	log.Println("CreateKey: Inserting key into database")
	key.CreatedAt = time.Now()

	result, err := database.KeyCollection.InsertOne(context.TODO(), key)
	if err != nil {
		log.Println("CreateKey: Insert failed: ", err)
		return err
	}

	log.Printf("CreateKey: success, created key with id=%v for user_id=%s", result.InsertedID, key.UserID.Hex())
	return nil
}
