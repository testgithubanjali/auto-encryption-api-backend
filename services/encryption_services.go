package services
import (
	"os"

	
	"auto-encryption-api-backend/utils"

)
func EncryptUserText(text string) (string, error){
	key := []byte(os.Getenv("ENCRYPTION_KEY"))
	return utils.EncryptText(text, key)

}
func DecryptUserText(cipherText string) (string, error){
	key := []byte(os.Getenv("ENCRYPTION_KEY"))
	return utils.DecryptText(cipherText, key)
	
}