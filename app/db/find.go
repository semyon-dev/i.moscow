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
	err := db.Collection("events").FindOne(context.Background(), filter).Decode(&event)
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
	cursor, err := db.Collection("events").Find(context.Background(), bson.M{})
	if err != nil {
		log.Println(err)
	}
	if err = cursor.All(context.Background(), &events); err != nil {
		log.Println(err)
	}
	return
}

func GetProjects(id primitive.ObjectID) (projects []model.Project) {
	filter := bson.M{"$and": bson.A{
		bson.M{"teamCapitan": bson.M{"$ne": id}},
		bson.M{"teamIDs": bson.M{"$nin": bson.A{id}}},
	}}
	projects = make([]model.Project, 0, 0)
	cursor, err := db.Collection("projects").Find(context.Background(), filter)
	if err != nil {
		log.Println(err)
	}
	if err = cursor.All(context.Background(), &projects); err != nil {
		log.Println(err)
	}
	return
}

func GetMyProjects(capitanId string) (projects []model.Project) {
	id, _ := primitive.ObjectIDFromHex(capitanId)
	filter := bson.M{"teamCapitan": id}
	cursor, err := db.Collection("projects").Find(context.Background(), filter)
	if err != nil {
		log.Println(err)
	}
	if err = cursor.All(context.Background(), &projects); err != nil {
		log.Println(err)
	}
	return
}

func FindUserByEmail(email string) (user model.User, isExist bool) {
	filter := bson.M{"email": email}
	err := db.Collection("users").FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return model.User{}, false
		}
		return
	}
	return user, true
}

func GetProjectById(id primitive.ObjectID) (project model.Project, isExist bool) {
	filter := bson.M{"_id": id}
	err := db.Collection("projects").FindOne(context.Background(), filter).Decode(&project)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return model.Project{}, false
		}
		return
	}
	return project, true
}

func FindUserById(id primitive.ObjectID) (user model.User, isExist bool) {
	filter := bson.M{"_id": id}
	err := db.Collection("users").FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return model.User{}, false
		}
		return
	}
	return user, true
}

type reply struct {
	Name string `json:"Name" bson:"Name"`
}

func FullTextSearch(text string, limit int64) (endRep []string) {
	filter := bson.M{"$text": bson.M{"$search": text}}
	opts := options.Find().SetLimit(limit).SetMaxTime(time.Second * 3)
	cursor, err := db.Collection("skills").Find(context.Background(), filter, opts)
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
