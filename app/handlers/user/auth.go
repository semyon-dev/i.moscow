package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"i-moscow-backend/app/db"
	"i-moscow-backend/app/session"
	"log"
	"net/http"
)

func Auth(c *gin.Context) {

	jsonInput := struct {
		Email       string `json:"email"`
		Password    string `json:"password"`
		DeviceToken string `json:"deviceToken"`
	}{}

	if err := c.ShouldBindJSON(&jsonInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "not all parameters are specified",
		})
		return
	}

	user, exist := db.FindUserByEmail(jsonInput.Email)
	if !exist {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid email",
		})
		return
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(jsonInput.Password)) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid password",
		})
		return
	}

	err := db.UpdateDeviceToken(user.Id, jsonInput.DeviceToken)
	if err != nil {
		log.Println("can't update token: ", err)
	}

	token, err := session.Create(user.Email, user.Id.Hex())
	if err != nil {
		fmt.Println("Error in generating JWT: " + err.Error())
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"token":   token,
	})
}
