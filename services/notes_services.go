package services

import (
"context"
"time"


"auto-encryption-api-backend/database"
"auto-encryption-api-backend/models"

"go.mongodb.org/mongo-driver/bson"
"go.mongodb.org/mongo-driver/bson/primitive"


)

func CreateNote(note models.Note) error {


ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

note.CreatedAt = time.Now()

_, err := database.NoteCollection.InsertOne(ctx, note)

return err


}

func GetUserNotes(userID primitive.ObjectID) ([]models.Note, error) {


ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

filter := bson.M{
	"user_id": userID,
}

cursor, err := database.NoteCollection.Find(ctx, filter)
if err != nil {
	return nil, err
}

defer cursor.Close(ctx)

var notes []models.Note

if err = cursor.All(ctx, &notes); err != nil {
	return nil, err
}

return notes, nil


}

func DeleteNote(noteID primitive.ObjectID) error {


ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

filter := bson.M{
	"_id": noteID,
}

_, err := database.NoteCollection.DeleteOne(ctx, filter)

return err


}
