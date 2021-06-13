package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"i-moscow-backend/app/db"
	"i-moscow-backend/app/model"
	"i-moscow-backend/app/session"
	"net/http"
	"time"
)

func Register(c *gin.Context) {

	var userToken string

	var user model.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "not all parameters are specified",
		})
		return
	}

	if user.Password == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "password can't be empty",
		})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
		})
		return
	}

	user.Id = primitive.NewObjectID()
	user.Password = string(hashedPassword)
	user.CreatedAt = time.Now()

	err = db.Insert("users", user)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
		})
		return
	}

	userToken, err = session.Create(user.Email, user.Id.Hex())
	if err != nil {
		fmt.Println("Error in generating JWT: " + err.Error())
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "ok",
		"token":   userToken,
	})
}
