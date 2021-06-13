package db

import (
	"context"
)

func Insert(collection string, document interface{}) (err error) {
	_, err = db.Collection(collection).InsertOne(context.Background(), document)
	if err != nil {
		return err
	}
	return nil
}
