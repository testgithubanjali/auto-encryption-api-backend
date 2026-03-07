package services
import(
	"auto-encryption-api-backend/database"
	"auto-encryption-api-backend/models"
	"context"

	"go.mongodb.org/mongo-driver/bson"
)
func CreateUser(user models.user) error{
	_, err := database.UserCollection.InsertOne(context.TODO(),user)
	return err
}

func GetUserByEmail(email string) (models.user, error){
	var user models.user
	filter := bson.M("email": email)
	err := database.UserCollection.FindOne(context.TODO(),filter).Decode(&user)
	if err!=nil{
		return nil,err
	}
	return &user, nil
}
func GetUserByID(id string) (models.user, error){
	var user models.user
	filer := bson.M("_id": id)
	err := database.UserCollection.FindOne(context.TODO(),filter).Decode(&user)
	if err!=nil{
		return nil,err
	}
	return &user,nil
}