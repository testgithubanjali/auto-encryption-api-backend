package database

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// 🔹 Global collections
var UserCollection *mongo.Collection
var KeyCollection *mongo.Collection
var EntryCollection *mongo.Collection
var RefreshTokenCollection *mongo.Collection
var SessionCollection *mongo.Collection

func ConnectDB() {
	log.Println("ConnectDB: Starting MongoDB connection...")
	mongoURI := os.Getenv("MONGO_URI")

	clientOptions := options.Client().ApplyURI(mongoURI)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// ⏱ Timeout context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	log.Println("ConnectDB: Pinging mongodb")
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to MongoDB!")

	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		log.Println("ConnectDB: DB_NAME not found in environment")
	}
	log.Printf("ConnectDB: using database: %s\n", dbName)
	db := client.Database(dbName)

	// 🔹 Collections
	UserCollection = db.Collection("users")
	KeyCollection = db.Collection("keys")
	EntryCollection = db.Collection("entries")
	RefreshTokenCollection = db.Collection("refresh_tokens")
	SessionCollection = db.Collection("sessions")
	log.Println("ConnectDB: Collections initialized:")
	log.Println(" - users")
	log.Println(" - keys")
	log.Println(" - entries")
	log.Println(" - refresh_tokens")

	log.Println("ConnectDB: Database setup completed successfully ")
}
