package handlers

import (
	"log"
	"net/http"

	"auto-encryption-api-backend/models"
	"auto-encryption-api-backend/services"
	"auto-encryption-api-backend/utils"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateKey(c *gin.Context) {
	log.Println("key creation started")
	var req models.Key
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("CreateKey: getting invalid request", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	log.Printf("CreateKey: request received, algorithm: %s", req.Algorithm)
	if req.Algorithm == "" || req.KeyValue == "" {
		log.Printf("CreateKey: algorithm and key value are missing ")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Algorithm and key value are required"})
		return
	}

	userIDStr := c.MustGet("user_id").(string)
	userID, err := primitive.ObjectIDFromHex(userIDStr)
	log.Printf("CreateKey: Processing for userID: %s", userIDStr)
	if err != nil {
		log.Println("CreateKey: getting invalid userID", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user ID"})
		return
	}
	hashedKey := utils.HashData([]byte(req.KeyValue))
	log.Println("CreateKey: Key hashed using SHA-256")
	key := models.Key{
		UserID:    userID,
		Algorithm: req.Algorithm,
		KeyValue:  hashedKey,
	}
	log.Println("CreateKey: saving key to database")
	err = services.CreateKey(key)
	if err != nil {
		log.Println("CreateKey: Key creation failed", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "key creation failed"})
		return
	}
	log.Println("CreateKey: Key creation Successfull")
	c.JSON(http.StatusOK, gin.H{"message": "Key created"})
}
