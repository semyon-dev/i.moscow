package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
	"i-moscow-backend/app/model"
	"log"
)

func UpdateUser(user model.User) (err error) {
	filter := bson.M{"_id": user.Id}

	if user.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		user.Password = string(hashedPassword)
	} else {
		var thisUser model.User
		result := usersCollection.FindOne(context.Background(), filter)
		err = result.Decode(&thisUser)
		if err != nil {
			log.Println(err)
			return err
		}
		user.Password = thisUser.Password
	}

	result := usersCollection.FindOneAndReplace(context.Background(), filter, user)
	if result.Err() != nil {
		log.Println(err)
		if err == mongo.ErrNoDocuments {
			log.Println("no docs")
		}
		return
	}
	return nil
}

func AddRegisteredEventToUser(email string, eventID primitive.ObjectID) {
	filter := bson.M{"email": email}
	update := bson.M{"$push": bson.M{"registeredEvents": eventID}}
	_, err := usersCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		fmt.Println(err)
	}
}
