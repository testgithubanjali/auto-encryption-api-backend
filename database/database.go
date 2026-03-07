package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
var userCollection *mongo.Collection
func ConnectDB() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.NewClient(context.TODO(),clientOptions)
	if err !=nil{
		log.Fatal(err)	
	}
	ctx , cancel := context.WithTimeout(context.Background(),10*time.Second)
	defer cancel()
	err = client.Ping(ctx,nil)
	If err != nil{
		log.Fatal(err)
	}
	log.Println("Connected to MongoDBD!")
	db := client.Database("auto-encryption-api-backend")
	UserCollection = db.Collection("users")
}