package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"i-moscow-backend/app/model"
)

func UpdateUser(user model.User) (isExist bool) {
	filter := bson.M{"_id": user.Id}

	update := bson.D{
		{"$set", bson.D{
			{"email", user.Email},
			{"password", user.Password},
			{"fio", user.FIO},
		}},
	}

	_, err := usersCollection.UpdateOne(context.Background(), update, filter)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false
		}
	}
	return true
}

func AddRegisteredEventToUser(email string, eventID primitive.ObjectID) {
	filter := bson.M{"email": email}
	update := bson.M{"$push": bson.M{"registeredEvents": eventID}}
	_, err := usersCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		fmt.Println(err)
	}
}
