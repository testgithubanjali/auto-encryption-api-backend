package middleware

import (
"fmt"
"net/http"
"strings"


"auto-encryption-api-backend/utils"

"github.com/gin-gonic/gin"
"github.com/golang-jwt/jwt/v5"


)

func AuthMiddleware() gin.HandlerFunc {


return func(c *gin.Context) {

	authHeader := c.GetHeader("Authorization")
	fmt.Println("Authorization Header:", authHeader)

	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Authorization header missing",
		})
		c.Abort()
		return
	}

	// Expect: Bearer TOKEN
	parts := strings.Split(authHeader, " ")

	if len(parts) != 2 || parts[0] != "Bearer" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid authorization format",
		})
		c.Abort()
		return
	}

	tokenString := parts[1]
	fmt.Println("Token received:", tokenString)

	claims := jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(
		tokenString,
		claims,
		func(token *jwt.Token) (interface{}, error) {

			// Validate signing method
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method")
			}

			return utils.GetJWTSecret(), nil
		},
	)

	if err != nil || !token.Valid {
		fmt.Println("JWT validation error:", err)

		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid or expired token",
		})
		c.Abort()
		return
	}

	fmt.Println("Token claims:", claims)

	userID, ok := claims["user_id"].(string)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "User ID missing in token",
		})
		c.Abort()
		return
	}

	// Store user_id in request context
	c.Set("user_id", userID)

	c.Next()
}


}
