package consumer

import (
	"context"
	"device/internal/alarm"
	"device/internal/connections"
	"device/internal/door"
	"device/internal/models"
	"device/internal/speaker"
	"device/internal/vaccum"
	"device/protos/deviceproto"
	"encoding/json"
	"fmt"
	"log"

	"github.com/rabbitmq/amqp091-go"
)

type Subscriber struct {
	Ch  *amqp091.Channel
	Ctx context.Context
	S   *speaker.Speaker
	V *vaccum.Vaccum
	A *alarm.Alarm
	D *door.Door
}

func NewSub() *Subscriber {
	ch, c := connections.Rabbitmq()
	s := speaker.NewSpeaker()
	v:=vaccum.NewVaccum()
	a:=alarm.NewAlarm()
	d:=door.NewDoor()
	return &Subscriber{Ch: ch, Ctx: c, S: s,V: v,A: a,D: d}
}

type Message struct {
	Type    string          `json:"type"`
	Payload json.RawMessage `json:"payload"`
}

func (u *Subscriber) Subscriber() {
	fmt.Println("consumer is ready")
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
	msgs, err := u.Ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		log.Fatal(err)
	}
	forever := make(chan bool)
	go func() {
		fmt.Println(`in the gourouting`)
		for d := range msgs {
			fmt.Printf("message is %v\n", string(d.Body))
			var msg Message
			err := json.Unmarshal(d.Body, &msg)
			if err != nil {
				log.Printf("Failed to unmarshal message: %s", err)
				continue
			}
			if err := u.Adjust(msg); err != nil {
				log.Printf("adjust error %v", err)
			}
		}
	}()
	<-forever
}

func (u *Subscriber) Adjust(msg Message) error {
	switch msg.Type {
	case "speaker":
		var req models.Speaker
		if err := json.Unmarshal(msg.Payload, &req); err != nil {
			return err
		}
		if err := u.Speaker(req); err != nil {
			return err
		}
	case "vaccum":
		var req models.Vaccum
		if err := json.Unmarshal(msg.Payload, &req); err != nil {
			return err
		}
		if err:=u.Vaccum(req);err!=nil{
			return err
		}
	case "alarm":
		var req models.Alarm
		if err := json.Unmarshal(msg.Payload, &req); err != nil {
			return err
		}
		if err:=u.Alarm(req);err!=nil{
			return err
		}
		return nil
	case "door":
		var req models.Door
		if err := json.Unmarshal(msg.Payload, &req); err != nil {
			return err
		}
		if err:=u.Door(req);err!=nil{
			return err
		}
	}
	return nil
}

func (u *Subscriber) Speaker(req models.Speaker) error {
	var req1 = deviceproto.SpeakerCommandRequest{
		VolumeCommand: req.Volume,
		Control:       req.Song,
		TurnCommand:   req.Turn,
		AddSong:       req.Add,
		Dtype:         req.Dname,
	}
	fmt.Printf("speaker is %v", req)
	if err := u.S.Commands(&req1); err != nil {
		return err
	}
	return nil
}

func (u *Subscriber) Vaccum(req models.Vaccum) error {
	var req1 = deviceproto.VaccumCleanaer{
		TurnCommand:   req.Turn_command,
		Location:      req.Location,
		Time:          int64(req.Timer),
		Dname:         req.Dname,
		Status:        req.Status,
		RemainingTime: int64(req.Remaining_time),
	}
	if err:=u.V.Commands(&req1);err!=nil{
		return err
	}
	return nil
}

func (u *Subscriber) Alarm(req models.Alarm)error{
	var req1=deviceproto.SmartAlarm{
		LiftingCurtain: req.Lift_curtain,
		LightCommand: req.Lamp,
		LightColor: req.Color,
		Dname: req.Dname,
		Status: req.Status,
		Setalarm: int32(req.Alarm),
		RemainingTime: int32(req.Remaining_time),
	}
	if err:=u.A.AlarmCommands(&req1);err!=nil{
		return err
	}
	return nil
}

func (u *Subscriber) Door(req models.Door)error{
	var req1=deviceproto.LockDoor{
		Door: req.Door,
		Command: req.Command,
		Time: int32(req.Timer),
		Dname: req.Dname,
		Status: req.Status,
		RemainingTime: int32(req.Remaining_time),
	}
	if err:=u.D.DoorCommands(&req1);err!=nil{
		return err
	}
	return nil
}