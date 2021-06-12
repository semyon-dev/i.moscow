package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"i-moscow-backend/app/config"
	"log"
	"time"
)

var (
	client           *mongo.Client
	usersCollection  *mongo.Collection
	eventsCollection *mongo.Collection
	skillsCollection *mongo.Collection
)

func Connect() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var err error
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(config.MongoUrl))
	if err != nil {
		log.Fatal(err)
	}

	// Create connect
	err = client.Connect(ctx)
	if err != nil {
		fmt.Println("client MongoDB:", err)
	} else {
		fmt.Println("✔ Подключение client MongoDB успешно")
	}

	database := client.Database("main")
	usersCollection = database.Collection("users")
	skillsCollection = database.Collection("skills")
	eventsCollection = database.Collection("events")

	err = Ping()
	if err == nil {
		fmt.Println("Connected to MongoDB!")
		return
	}
	fmt.Println(err.Error())
}

func Ping() error {
	return client.Ping(context.Background(), readpref.Primary())
}
