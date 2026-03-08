package handlers
import(
	"auto-encryption-api-backend/models"
	"auto-encryption-api-backend/services"
	"auto-encryption-api-backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)
func SignUpUser(c *gin.Context){
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	existingUser, _ := services.GetUserByEmail(user.Email)
	if existingUser != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error" : "User already exists"})
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password),14)
     if err !=nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error":  "Password hashing failed"})
		return
	 }
	 user.Password = string(hashedPassword)

	 err = services.CreateUser(user)
	 if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "user creation failed",
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message" : "user created successfully",
	})
	 }

func LoginUser(c *gin.Context){
		var input models.User
		if err := c.ShouldBindJSON(&input); err!= nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error" : "invalid input",
			})
			return
		}
		user, err := services.GetUserByEmail(input.Email)
		if err!=nil{
		   c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid email or password",
		   })
		   return
		}
		err=bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(input.Password))
		if err!=nil{
			c.JSON(http.StatusUnauthorized, gin.H{
				"error" : "Invalid email or password",
			})
			return
		}
		token, err := utils.GenerateToken(user.ID.Hex())
		if err!=nil{
			c.JSON(http.StatusInternalServerError, gin.H{
				"error" : "Token generation failed",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "Login successful",
			"token" : token,
		})
		
	 }
	 func UserProfile(c *gin.Context){
		userID := c.MustGet("user_id").(string)
		c.Set("user_id", userID)
		user, err := services.GetUserByID(userID)
		if err!=nil{
			c.JSON(http.StatusUnauthorized, gin.H{"error" : "User not found"})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"user": gin.H{
			"id" : user.ID,
			"email" : user.Email,
			},
			
		})
			
		}