package command

import (
	"context"
	"device/internal/models"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Command struct {
	D *mongo.Collection
	C context.Context
}

func (u *Command) SAddCommand(req *models.SpeakerCommand) error {
	a := bson.M{
		"turn":   req.Turn,
		"volume": req.Volume,
		"songid": req.SongId,
		"song":   req.Song,
		"dname":  req.Dname,
	}
	_, err := u.D.InsertOne(u.C, a)
	if err != nil {
		return err
	}
	return nil
}

func (u *Command) SUpdate(req *models.SpeakerCommand) error {
	_, err := u.D.UpdateOne(u.C, bson.M{"dname": req.Dname}, bson.M{"$set": bson.M{
		"turn":   req.Turn,
		"song":   req.Song,
		"volume": req.Volume,
		"songid": req.SongId}})
	if err != nil {
		return err
	}
	return nil
}

func (u *Command) SFind(req string) (*models.SpeakerCommand, error) {
	var res models.SpeakerCommand
	if err := u.D.FindOne(u.C, bson.M{"dname": req}).Decode(&res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (u *Command) Vadd(req *models.Vaccum) error {
	a := bson.M{
		"dname":          req.Dname,
		"turn":           req.Turn_command,
		"location":       req.Location,
		"timer":          req.Timer,
		"created_at":     req.Created_at,
		"remaining_time": req.Remaining_time,
		"status":         req.Turn_command,
	}
	_, err := u.D.InsertOne(u.C, a)
	if err != nil {
		return err
	}
	return nil
}

func (u *Command) VUpdate(req *models.Vaccum) error {
	_, err := u.D.UpdateOne(u.C, bson.M{"dname": req.Dname}, bson.M{"$set": bson.M{
		"turn":           req.Turn_command,
		"location":       req.Location,
		"timer":          req.Timer,
		"created_at":     req.Created_at,
		"remaining_time": req.Remaining_time,
		"status":         req.Status}})
	if err != nil {
		return err
	}
	return nil
}

func (u *Command) Vfind(req string) (*models.Vaccum, error) {
	var res models.Vaccum
	if err := u.D.FindOne(u.C, bson.M{"dname": req}).Decode(&res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (u *Command) Aadd(req *models.Alarm) error {
	a := bson.M{
		"curtain":        req.Lift_curtain,
		"lamp":           req.Lamp,
		"color":          req.Color,
		"dname":          req.Dname,
		"status":         req.Status,
		"alarm":          req.Alarm,
		"created_at":     req.Created_at,
		"remaining_time": req.Remaining_time,
	}
	_, err := u.D.InsertOne(u.C, a)
	if err != nil {
		return err
	}
	return nil
}

func (u *Command) Aupdate(req *models.Alarm) error {
	_, err := u.D.UpdateOne(u.C, bson.M{"dname": req.Dname}, bson.M{"$set": bson.M{
		"curtain":        req.Lift_curtain,
		"lamp":           req.Lamp,
		"color":          req.Color,
		"alarm":          req.Alarm,
		"created_at":     req.Created_at,
		"remaining_time": req.Remaining_time}})
	if err != nil {
		return err
	}
	return nil
}

func (u *Command) Afind(req string) (*models.Alarm, error) {
	var res models.Alarm
	if err := u.D.FindOne(u.C, bson.M{"dname": req}).Decode(&res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (u *Command) Dadd(req *models.Door) error {
	a := bson.M{
		"door":           req.Door,
		"command":        req.Command,
		"status":         req.Status,
		"dname":          req.Dname,
		"timer":          req.Timer,
		"remaining_time": req.Remaining_time,
		"created_at":     req.Created_at}
	_, err := u.D.InsertOne(u.C, a)
	if err != nil {
		return err
	}
	return nil
}

func (u *Command) Dupdate(req *models.Door) error {
	_, err := u.D.UpdateOne(u.C, bson.M{"dname": req.Dname}, bson.M{"$set": bson.M{
		"door":           req.Door,
		"commmand":       req.Command,
		"status":         req.Status,
		"time":           req.Timer,
		"remaining_time": req.Remaining_time,
		"created_at":     req.Created_at}})
	if err != nil {
		return err
	}
	fmt.Println("update", req)
	return nil
}

func (u *Command) Dfind(req string) (*models.Door, error) {
	if err := u.TIMER(req); err != nil {
		return nil, err
	}
	var res models.Door
	if err := u.D.FindOne(u.C, bson.M{"dname": req}).Decode(&res); err != nil {
		return nil, err
	}
	fmt.Println("Dfind ", res)
	return &res, nil
}

func (u *Command) GDelete(req string) error {
	_, err := u.D.DeleteOne(u.C, bson.M{"dname": req})
	if err != nil {
		return err
	}
	return nil
}

func (u *Command) TIMER(req string) error {
	var res models.Door
	if err := u.D.FindOne(u.C, bson.M{"dname": req}).Decode(&res);err != nil {
		return err
	}
	time1 := time.Duration(int64(res.Timer) * int64(time.Minute))
	elapsed := time.Since(res.Created_at)
	remaining_time := int(time1.Minutes()) - int(elapsed.Minutes())
	_, err := u.D.UpdateOne(u.C, bson.M{"dname": req}, bson.M{"$set": bson.M{"remaining_time": remaining_time}})
	if err != nil {
		return err
	}
	return nil
}
