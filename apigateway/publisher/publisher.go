package publisher

import (
	"api/internal/connection"
	"api/internal/models"
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/rabbitmq/amqp091-go"
)

type Publisher struct {
	Ch  *amqp091.Channel
	Ctx context.Context
}

func NewPub() *Publisher {
	ch, ctx := connection.Rabbitmq()
	return &Publisher{Ch: ch, Ctx: ctx}
}

func (u *Publisher) Publish(que string, msg models.Message) error {
	body, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	if err = u.Ch.Publish(
		"",    // exchange
		que,   // routing key
		false, // mandatory
		false, // immediate
		amqp091.Publishing{
			ContentType: "application/json",
			Body:        body,
		}); err != nil {
		return err
	}
	fmt.Println(string(body))
	return nil
}

func (u *Publisher) Adjust(msg models.Message) error {
	q, err := u.Ch.QueueDeclare(
		"deviceQueue", // name
		true,          // durable
		false,         // delete when unused
		false,         // exclusive
		false,         // no-wait
		nil,           // arguments
	)
	if err != nil {
		log.Println(err)
	}
	if err = u.Publish(q.Name, msg); err != nil {
		return err
	}
	return nil
}
