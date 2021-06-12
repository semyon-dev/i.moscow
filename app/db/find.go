package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"i-moscow-backend/app/model"
	"log"
	"time"
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

type reply struct {
	Name string `json:"Name" bson:"Name"`
}

func FullTextSearch(text string) (endRep []string) {
	filter := bson.M{"$text": bson.M{"$search": text}}
	opts := options.Find().SetLimit(50).SetMaxTime(time.Second * 3)
	cursor, err := skillsCollection.Find(context.Background(), filter, opts)
	if err != nil {
		log.Println(err)
	}
	// выделяем память заранее
	rep := make([]reply, 0, 50)
	endRep = make([]string, 0, 50)
	err = cursor.All(context.Background(), &rep)
	if err != nil {
		log.Fatalln(err)
	}
	for _, v := range rep {
		if v.Name == "" {
			continue
		}
		endRep = append(endRep, v.Name)
	}
	return
}

//func RegexSearch(text string) (endRep []string) {
//	filter := bson.M{"positionName": bson.M{"$regex": text, "$options": "i"}}
//	cursor, err := skillsCollection.Find(context.Background(), filter)
//	if err != nil {
//		log.Println(err)
//	}
//	rep := make([]reply, 0, 50)
//	err = cursor.All(context.Background(), &rep)
//	if err != nil {
//		log.Fatalln(err)
//	}
//	for i, v := range rep {
//		if i == 50 {
//			break
//		}
//		endRep = append(endRep, v.Name)
//	}
//	return
//}
