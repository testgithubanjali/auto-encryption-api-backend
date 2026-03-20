package handlers

import (
	"net/http"

	"auto-encryption-api-backend/models"
	"auto-encryption-api-backend/services"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type KeyRequest struct {
	Algorithm string `json:"algorithm"`
	KeyValue  string `json:"key_value"`
}
func CreateKey(c *gin.Context){
	var req KeyRequest
	if err := c.ShouldBindJSON(&req); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	userIDStr := c.MustGet("user_id").(string)
	userID,_ := primitive.ObjectIDFromHex(userIDStr)
	key := models.Key{
		UserID: userID,
		Algorithm: req.Algorithm,
		KeyValue: req.KeyValue,

}
err := services.CreateKey(key)
if err != nil {
	c.JSON(http.StatusInternalServerError, gin.H{"error": "key creation failed"})
	return
}
c.JSON(http.StatusOK, gin.H{"message": "Key created"})
}