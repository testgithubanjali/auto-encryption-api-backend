package handlers
import (
	"net/http"

	"auto-encryption-api-backend/models"
	"auto-encryption-api-backend/services"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EntryRequest struct {
	KeyID      string `json:"key_id"`
	Ciphertext string `json:"ciphertext"`
}
func SaveEntry(c *gin.Context){
	var req EntryRequest
	if err:= c.ShouldBindJSON(&req); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Input"})
			return
	}
	userIDStr := c.MustGet("user_id").(string)
	userID, _:= primitive.ObjectIDFromHex(userIDStr)
	keyID, _:= primitive.ObjectIDFromHex(req.KeyID)
	entry := models.Entry{
		UserID: userID,
		KeyID: keyID,
		Ciphertext: req.Ciphertext,
	}
	err := services.SavaEntry(entry)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save entry"	})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Entry Saved"})
}