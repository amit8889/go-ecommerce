package database

import (
	"log"
	"log/slog"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBSet() *mongo.client {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal("=========Error in mongodb connection==>", err)
	}
	context, cancle := context.WithTimeOut(context.Background(), 10*time.Secound)
	defer cancle()
	err = client.Connect(context)
	if err != nil {
		log.Fatal("=========Error1 in mongodb connection==>", err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal("=========Error2 in mongodb connection==>", err)
	}
	slog.Info("===mongodb connected successfully====")
	return client

}

var Client *mongo.client = DBSet()

func UserData(client *mongo.Client, collectionName string) *mongo.Collection {

}
func ProductData(client *mongo.Client, collectionName string) *mongo.Collection {

}
