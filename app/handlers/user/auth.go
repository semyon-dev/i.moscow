package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"i-moscow-backend/app/db"
	"i-moscow-backend/app/session"
	"net/http"
	"strings"
)

func Auth(c *gin.Context) {

	jsonInput := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
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

	token, err := session.Create(user.Email)
	if err != nil {
		fmt.Println("Error in generating JWT: " + err.Error())
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"token":   token,
	})
}

func ParseBearer(c *gin.Context) (email string, isValid bool) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.AbortWithStatus(http.StatusUnauthorized)
		return "", false
	}

	headerParts := strings.Split(authHeader, " ")
	if len(headerParts) != 2 {
		c.AbortWithStatus(http.StatusUnauthorized)
		return "", false
	}

	if headerParts[0] != "Bearer" {
		c.AbortWithStatus(http.StatusUnauthorized)
		return "", false
	}
	email = session.ParseToken(headerParts[1])
	return email, true
}
