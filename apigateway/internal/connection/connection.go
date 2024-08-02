package connection

import (
	redisserver "api/internal/redis"
	"context"
	"log"
	"time"

	"github.com/rabbitmq/amqp091-go"
	"github.com/redis/go-redis/v9"
)

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
func Redis()*redisserver.Redis{
	client := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})
	ctx:=context.Background()
	_, err := client.Ping(ctx).Result()
	if err != nil {
		log.Fatal(err)
	}
	return &redisserver.Redis{R: client,C: ctx}
}
