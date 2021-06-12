package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

var (
	AccessSecret string
	MongoUrl     string
	Port         string
)

func Load() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error in loading configs: " + err.Error())
	} else {
		AccessSecret = os.Getenv("ACCESS_SECRET")
		MongoUrl = os.Getenv("MONGO_URL")
		Port = os.Getenv("PORT")
	}
}
