package connections

import (
	"context"
	"device/internal/database/command"
	"device/internal/database/devicedatabase"
	database "device/internal/database/userdevice"
	"log"
	"time"

	"github.com/rabbitmq/amqp091-go"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DeviceMongo() *database.Mongo {
	opt, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://mongodb:27017"))
	if err != nil {
		log.Fatal(err)
	}
	if err := opt.Ping(context.Background(), options.Client().ReadPreference); err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	data := opt.Database("device").Collection("devices")
	return &database.Mongo{D: data, C: ctx}
}

func DeviceCommand() *command.Command {
	opt, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://mongodb:27017"))
	if err != nil {
		log.Fatal(err)
	}
	if err := opt.Ping(context.Background(), options.Client().ReadPreference); err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	data := opt.Database("device").Collection("commands")
	return &command.Command{D: data, C: ctx}
}

func DeviceLIST() *devicedatabase.LIST {
	opt, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://mongodb:27017"))
	if err != nil {
		log.Fatal(err)
	}
	if err := opt.Ping(context.Background(), options.Client().ReadPreference); err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	data := opt.Database("device").Collection("list")
	return &devicedatabase.LIST{D: data, C: ctx}
}

func Rabbitmq() (*amqp091.Channel, context.Context) {
	var conn *amqp091.Connection
	var err error

	for i := 0; i < 5; i++ {
		conn, err = amqp091.Dial("amqp://guest:guest@rabbitmq:5672/")
		if err == nil {
			break
		}
		log.Printf("RabbitMQ connection failed: %v. Retrying...", err)
		time.Sleep(5 * time.Second)
	}

	if err != nil {
		log.Fatal("Failed to connect to RabbitMQ after multiple attempts:", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal("Failed to open a channel:", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return ch, ctx
}