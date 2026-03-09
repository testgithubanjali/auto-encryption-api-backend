package main
import(
	"os"

	"auto-encryption-api-backend/database"
	"auto-encryption-api-backend/routers"

	"github.com/gin-contrib/cors"
	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
)

func main(){
	godotenv.Load()
	database.ConnectDB()
	router := gin.Default()
router.Use(cors.New(cors.Config{
    AllowOrigins:     []string{"http://localhost:3001"},
    AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
    AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
    AllowCredentials: true,
}))
routers.RegisterRoutes(router)
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	router.Run(":" + port)

	
}