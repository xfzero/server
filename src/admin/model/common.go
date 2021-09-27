package model

import (
	"context"
	"gm/admin/service"

	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/mongo"
)

var dbName = "ngame"

func getCollection(cName string) *mongo.Collection {
	return service.GetMgoClient().Database(dbName).Collection(cName)
}

func FindOne(cName string, filter interface{}, ret interface{}, opts ...*options.FindOneOptions) (not bool, err error) {
	err = getCollection(cName).FindOne(context.TODO(), filter).Decode(ret)
	if err != nil && err == mongo.ErrNoDocuments {
		not = true
		err = nil
	}

	return
}
