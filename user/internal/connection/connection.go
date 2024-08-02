package connection

import (
	"context"
	"log"
	"user/internal/mongodb"
	service "user/services"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Mongo() (*mongodb.Mongo) {
	opt, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://mongodb:27017"))
	if err != nil {
		log.Fatal(err)
	}
	if err := opt.Ping(context.Background(), options.Client().ReadPreference); err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	data := opt.Database("user").Collection("users")
	return &mongodb.Mongo{D: data, C: ctx}
}

func NewServer() *service.Service {
	mongoInstance:= Mongo()
	return &service.Service{R: mongoInstance}
}
