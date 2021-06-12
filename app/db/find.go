package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"i-moscow-backend/app/model"
	"log"
)

func GetEventByID(id primitive.ObjectID) (event model.Event, isExist bool) {
	filter := bson.M{"_id": id}

	err := eventsCollection.FindOne(context.Background(), filter).Decode(&event)
	if err != nil {
		fmt.Println(err)
		if err == mongo.ErrNoDocuments {
			return model.Event{}, false
		}
		return
	}
	return event, true
}

func GetEvents() (events []model.Event) {
	cursor, err := eventsCollection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Println(err)
	}
	if err = cursor.All(context.Background(), &events); err != nil {
		log.Println(err)
	}
	return
}

func FindUserByEmail(email string) (User model.User, isExist bool) {
	filter := bson.M{"email": email}

	err := usersCollection.FindOne(context.Background(), filter).Decode(&User)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return model.User{}, false
		}
		return
	}
	return User, true
}
