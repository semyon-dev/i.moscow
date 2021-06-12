package session

import (
	"github.com/dgrijalva/jwt-go"
	"i-moscow-backend/app/config"
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
