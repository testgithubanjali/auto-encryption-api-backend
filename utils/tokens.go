package utils
import(
	"time"

	"github.com/golang-jwt/jwt/v5"
)
var JWT_SECRET = []byte("my secret  key")
func GenerateToken(userID string) (string, error){
	claims := jwt.MapClaims{
		"user_id" : userID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString , err := token.SignedString(JWT_SECRET)
	if err!=nil{
		return "", err
	}
	return tokenString, nil
}