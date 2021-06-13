package session

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"i-moscow-backend/app/config"
	"log"
	"net/http"
	"strings"
	"time"
)

func Create(email string, id string) (string, error) {
	var err error
	//Creating Access Token
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["email"] = email
	atClaims["id"] = id
	atClaims["exp"] = time.Now().Add(time.Minute * 24 * 60).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(config.AccessSecret))
	if err != nil {
		return "", err
	}
	return token, nil
}

type MyCustomClaims struct {
	Email string `json:"email"`
	Id    string `json:"id"`
	jwt.StandardClaims
}

func ParseToken(tokenString string) (id, email string) {
	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.AccessSecret), nil
	})

	if token == nil {
		log.Println("empty token")
		return
	}

	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		email = claims.Email
		id = claims.Id
		//fmt.Printf("%v %v", claims.Email, claims.StandardClaims.ExpiresAt)
	} else {
		fmt.Println(err)
	}
	return
}

func ParseBearer(c *gin.Context) (id, email string, isValid bool) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	headerParts := strings.Split(authHeader, " ")
	if len(headerParts) != 2 {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if headerParts[0] != "Bearer" {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	id, email = ParseToken(headerParts[1])
	isValid = true
	return
}
