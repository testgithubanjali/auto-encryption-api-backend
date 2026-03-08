package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
var UserCollection *mongo.Collection
func ConnectDB() {
	clientOptions := options.Client().ApplyURI("mongodb://mongodb:27017")
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
	db := client.Database("auto-encryption-api-backend")
	UserCollection = db.Collection("users")
}