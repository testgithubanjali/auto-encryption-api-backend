package utils

import (
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GetJWTSecret() []byte {
	secret := []byte(os.Getenv("JWT_SECRET"))
	log.Printf("GetJWTSecret: retrieved JWT secret, length=%d", len(secret))
	return secret
}

func GenerateAccessToken(userID string) (string, error) {
	log.Printf("GenerateAccessToken: generating access token for user_id=%s", userID)

	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Minute * 15).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(GetJWTSecret())
	if err != nil {
		log.Printf("GenerateAccessToken: failed to sign token for user_id=%s: %v", userID, err)
		return "", err
	}

	log.Printf("GenerateAccessToken: success, generated token for user_id=%s", userID)
	return tokenString, nil
}

func GenerateRefreshToken(userID string) (string, error) {
	log.Printf("GenerateRefreshToken: generating refresh token for user_id=%s", userID)

	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24 * 7).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(GetJWTSecret())
	if err != nil {
		log.Printf("GenerateRefreshToken: failed to sign token for user_id=%s: %v", userID, err)
		return "", err
	}

	log.Printf("GenerateRefreshToken: success, generated token for user_id=%s", userID)
	return tokenString, nil
}
