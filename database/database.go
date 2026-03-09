package database

import (
	"os"
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
var UserCollection *mongo.Collection
func ConnectDB() {
		mongoURI := os.Getenv("MONGO_URI")
	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(context.TODO(),clientOptions)
	if err !=nil{
		log.Fatal(err)	
	}
	ctx , cancel := context.WithTimeout(context.Background(),10*time.Second)
	defer cancel()
	err = client.Ping(ctx,nil)
	if err != nil{
		log.Fatal(err)
	}
	log.Println("Connected to MongoDBD!")
		dbName := os.Getenv("DB_NAME")
	db := client.Database(dbName)
	UserCollection = db.Collection("users")
}