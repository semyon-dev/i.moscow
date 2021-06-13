package db

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
)

func Delete(collection string, id string) (err error) {
	_, err = db.Collection(collection).DeleteOne(context.Background(), bson.M{"_id": id})
	return
}
