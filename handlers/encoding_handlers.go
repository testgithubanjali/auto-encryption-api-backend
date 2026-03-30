package handlers

import (
	"encoding/base64"
	"log"
	"net/http"
	"time"

	"auto-encryption-api-backend/models"
	"auto-encryption-api-backend/services"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func EncodeText(c *gin.Context) {

	log.Println("EncodeText: encoding started")
	var req models.EncodeRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("EncodeText: invalid request", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if req.Text == "" {
		log.Println("EncodeText: text is required")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Text required"})
		return
	}
	log.Println("EncodeText: Encoded text")
	encoded := base64.StdEncoding.EncodeToString([]byte(req.Text))

	userIDStr := c.MustGet("user_id").(string)
	log.Printf("EncodeText: processing requiring for userID: %s", userIDStr)
	userID, err := primitive.ObjectIDFromHex(userIDStr)
	if err != nil {
		log.Println("EncodeText: invalid user id ")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user ID"})
		return
	}

	entry := models.Entry{
		UserID:        userID,
		Type:          "encoding",
		OriginalText:  req.Text,
		ProcessedText: encoded,
		CreatedAt:     time.Now(),
	}
	log.Println("EncodeText: saving entry to database")
	err = services.SaveEntry(entry)
	if err != nil {
		log.Println("EncodeText: failed to save entry", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save entry"})
		return
	}
	log.Println("EncodeText: Encoded text successfully")

	c.JSON(http.StatusOK, gin.H{
		"encoded": encoded,
	})
}

func DecodeText(c *gin.Context) {
	log.Println("DecodeText API called")
	var req models.DecodeRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("DecodeText: invalid request", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if req.Encoded == "" {
		log.Println("EncodeText: encode text field is empty")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Encoded text required"})
		return
	}
	log.Println("DecodeText: Decoding text")
	decodedBytes, err := base64.StdEncoding.DecodeString(req.Encoded)
	if err != nil {
		log.Println("DecodeText: Decoding failed", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Decode failed"})
		return
	}

	decoded := string(decodedBytes)

	// 👤 Get user
	userIDStr := c.MustGet("user_id").(string)
	log.Printf("DecodeText: Processing for userID: %s", userIDStr)
	userID, err := primitive.ObjectIDFromHex(userIDStr)
	if err != nil {
		log.Println("DecodeText: getting invalid user ID", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user ID"})
		return
	}

	entry := models.Entry{
		UserID:        userID,
		Type:          "decoding",
		OriginalText:  req.Encoded,
		ProcessedText: decoded,
		CreatedAt:     time.Now(),
	}
	log.Printf("DecodeText: Saving decoding to database")
	err = services.SaveEntry(entry)
	if err != nil {
		log.Println("DecodeText: failed to save entry", err)
	} else {
		log.Printf("DecodeText: Successfully saved decoding entry for user %s", userIDStr)
	}
	log.Println(" DecodeText: Decoding successfull")
	c.JSON(http.StatusOK, gin.H{
		"text": decoded,
	})
}
