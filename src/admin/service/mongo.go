package service

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/mongo"
)

var (
	address = "mongodb://localhost:27017"
	timeout = 5 * time.Second
	mgoCli  *mongo.Client
)

func initEngine() {
	var err error
	clinetOptions := options.Client().ApplyURI(address).SetConnectTimeout(timeout)

	//链接到MongoDB
	mgoCli, err = mongo.Connect(context.TODO(), clinetOptions)
	if err != nil {
		log.Fatal(err)
	}

	//检查链接
	err = mgoCli.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
}

func GetMgoClient() *mongo.Client {
	if mgoCli == nil {
		initEngine()
	}

	return mgoCli
}

func CloseMgo() bool {
	if mgoCli == nil {
		return true
	}

	err := mgoCli.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
		return false
	}
	fmt.Println("Connection to MongoDB closed.")

	return true
}
