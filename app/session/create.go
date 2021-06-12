package session

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"i-moscow-backend/app/config"
	"log"
	"time"
)

func Create(email string) (string, error) {
	var err error
	//Creating Access Token
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["email"] = email
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
	jwt.StandardClaims
}

func ParseToken(tokenString string) (email string) {
	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.AccessSecret), nil
	})

	if token == nil {
		log.Println("empty token")
		return
	}

	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		email = claims.Email
		fmt.Printf("%v %v", claims.Email, claims.StandardClaims.ExpiresAt)
	} else {
		fmt.Println(err)
	}
	return
}
