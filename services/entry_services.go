package services
import (
	"context"
	"time"

	"auto-encryption-api-backend/database"
	"auto-encryption-api-backend/models"
)
func SavaEntry(entry models.Entry) error{
	entry.CreatedAt = time.Now()
	_, err := database.EntryCollection.InsertOne(context.TODO(), entry)
	return err
}
