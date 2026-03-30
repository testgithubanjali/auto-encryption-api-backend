package handlers

import (
	"log"
	"net/http"
	"time"

	"auto-encryption-api-backend/models"
	"auto-encryption-api-backend/services"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SaveEntry(c *gin.Context) {
	log.Println("SaveEntry: start")

	var req models.EntryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("SaveEntry: getting invalid input", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if req.Text == "" {
		log.Println("SaveEntry: text missing")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Text required"})
		return
	}

	if req.KeyID == "" {
		log.Println("SaveEntry: key missing")
		c.JSON(http.StatusBadRequest, gin.H{"error": "KeyID required"})
		return
	}

	userIDStr := c.MustGet("user_id").(string)
	userID, err := primitive.ObjectIDFromHex(userIDStr)
	log.Printf("SaveEntry: Processing for userID: %s", userIDStr)
	if err != nil {
		log.Println("SaveEntry: invalid user Id", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user ID"})
		return
	}

	keyID, err := primitive.ObjectIDFromHex(req.KeyID)
	if err != nil {
		log.Println("SaveEntry: getting invalid key")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid key ID"})
		return
	}
	log.Println("SaveEntry: Encrypting text")
	encryptedText, err := services.EncryptUserText(req.Text, req.KeyID)
	if err != nil {
		log.Println("SaveEntry: Encryption failed", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Encryption failed"})
		return
	}

	entry := models.Entry{
		UserID:        userID,
		KeyID:         keyID,
		Type:          "encryption",
		OriginalText:  req.Text,
		ProcessedText: encryptedText,
		CreatedAt:     time.Now(),
	}
	log.Println("SaveEntry: Saving Entry to database")
	err = services.SaveEntry(entry)
	if err != nil {
		log.Println("SaveEntry: failed to save Entry", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save entry"})
		return
	}
	log.Println("SaveEntry: encryption saved sucessfull")

	c.JSON(http.StatusOK, gin.H{
		"message":   "Encryption saved",
		"encrypted": encryptedText,
	})
}
