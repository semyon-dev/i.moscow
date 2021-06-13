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
		result := db.Collection("users").FindOne(context.Background(), filter)
		err = result.Decode(&thisUser)
		if err != nil {
			log.Println(err)
			return err
		}
		user.Password = thisUser.Password
	}

	result := db.Collection("users").FindOneAndReplace(context.Background(), filter, user)
	if result.Err() != nil {
		log.Println(err)
		if err == mongo.ErrNoDocuments {
			log.Println("no docs")
		}
		return
	}
	return nil
}

func UpdateProject(project model.Project) (err error) {
	filter := bson.M{"_id": project.Id}

	var thisProject model.Project
	result := db.Collection("projects").FindOne(context.Background(), filter)
	err = result.Decode(&thisProject)
	if err != nil {
		log.Println(err)
		return err
	}
	project.Id = thisProject.Id
	project.TeamCapitan = thisProject.TeamCapitan
	project.TeamIDs = thisProject.TeamIDs

	resultReplace := db.Collection("projects").FindOneAndReplace(context.Background(), filter, project)
	if resultReplace.Err() != nil {
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
	_, err := db.Collection("users").UpdateOne(context.Background(), filter, update)
	if err != nil {
		fmt.Println(err)
	}
}

func AddMemberToProject(id, userId primitive.ObjectID) error {
	filter := bson.M{"_id": id}
	update := bson.M{"$push": bson.M{"teamIDs": userId}}
	_, err := db.Collection("projects").UpdateOne(context.Background(), filter, update)
	if err != nil {
		fmt.Println(err)
	}
	return err
}

func DeleteMemberFromProject(id, memberId primitive.ObjectID) error {
	filter := bson.M{"_id": id}
	update := bson.M{"$pull": bson.M{"teamIDs": memberId}}
	_, err := db.Collection("projects").UpdateOne(context.Background(), filter, update)
	if err != nil {
		fmt.Println(err)
	}
	return err
}

func DeleteRequestFromProject(projectId, userId primitive.ObjectID) error {
	filter := bson.M{"_id": projectId}
	update := bson.M{"$pull": bson.M{"requestedIds": userId}}
	_, err := db.Collection("projects").UpdateOne(context.Background(), filter, update)
	if err != nil {
		fmt.Println(err)
	}
	return err
}

func AddRequestMemberToProject(projectId, memberId primitive.ObjectID) error {
	filter := bson.M{"_id": projectId}
	update := bson.M{"$push": bson.M{"requestedIds": memberId}}
	_, err := db.Collection("projects").UpdateOne(context.Background(), filter, update)
	if err != nil {
		fmt.Println(err)
	}
	return err
}
