package handlers

import (
	"auto-encryption-api-backend/models"
	"auto-encryption-api-backend/services"
	"auto-encryption-api-backend/utils"

	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func SignUpUser(c *gin.Context) {

	var req models.SignupRequest

	// Parse request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request",
		})
		return
	}

	// Check if user already exists
	existingUser, _ := services.GetUserByEmail(req.Email)
	if existingUser != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "User already exists",
		})
		return
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Password hashing failed",
		})
		return
	}

	// Create user
	user := models.User{
		Email:    req.Email,
		Password: string(hashedPassword),
	}

	// Save user
	err = services.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "User creation failed",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Signup successful",
	})
}

func LoginUser(c *gin.Context) {

	var req models.LoginRequest

	
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request",
		})
		return
	}

	
	user, err := services.GetUserByEmail(req.Email)
	if err != nil || user == nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid email or password",
		})
		return
	}

	
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid email or password",
		})
		return
	}

	
	accessToken, err := utils.GenerateAccessToken(user.ID.Hex())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to generate access token",
		})
		return
	}

	refreshToken, err := utils.GenerateRefreshToken(user.ID.Hex())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to generate refresh token",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}
func UserProfile(c *gin.Context) {

	
	userID := c.MustGet("user_id")
	userID2, err := primitive.ObjectIDFromHex(userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Invalid user ID",
		})
		return
	}
	
	user, err := services.GetUserByID(userID2)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "User not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": gin.H{
			"id":    user.ID.Hex(),
			"email": user.Email,
		},
	})
}
