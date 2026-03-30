package services

import (
	"auto-encryption-api-backend/database"
	"auto-encryption-api-backend/models"
	"context"
	"log"
)

func SaveEntry(entry models.Entry) error {
	log.Println("SaveEntry: Inserting entry into database")
	result, err := database.EntryCollection.InsertOne(context.TODO(), entry)
	if err != nil {
		log.Println("SaveEntry: insertion failed")
		return err
	}
	log.Printf("SaveEntry: Insert Successful, ID: %v", result.InsertedID)
	return nil
}
