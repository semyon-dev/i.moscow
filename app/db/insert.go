package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"i-moscow-backend/app/model"
)

func InsertUser(User model.User) (err error) {
	_, err = usersCollection.InsertOne(context.Background(), User)
	if err != nil {
		return err
	}
	return nil
}

func Insert(collection *mongo.Collection, document interface{}) (err error) {
	_, err = collection.InsertOne(context.Background(), document)
	if err != nil {
		return err
	}
	return nil
}
