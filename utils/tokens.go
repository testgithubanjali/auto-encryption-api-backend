package utils
import(
	"time"
	"os"

	"github.com/golang-jwt/jwt/v5"
)
func GetJWTSecret() []byte {
	return []byte(os.Getenv("JWT_SECRET"))
}
func GenerateToken(userID string) (string, error){
	claims := jwt.MapClaims{
		"user_id" : userID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(GetJWTSecret())
	if err!=nil{
		return "", err
	}
	return tokenString, nil
}