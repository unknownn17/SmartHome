package database

import (
	"context"
	"device/internal/models"
	"device/protos/deviceproto"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Mongo struct {
	D *mongo.Collection
	C context.Context
}

func (u *Mongo) AddDevice(req *models.Device) error {
	_, err := u.D.InsertOne(u.C, req)
	if err != nil {
		return err
	}
	return nil
}

func (u *Mongo) DeleteDevice(req *deviceproto.DeleteDeviceRequest) error {
	_, err := u.D.DeleteOne(u.C, bson.M{"name": req.Name})
	if err != nil {
		return err
	}
	return nil
}

func (u *Mongo) ListDevices(user string) (*deviceproto.GetdevicesResponse, error) {
	resp, err := u.D.Find(u.C, bson.M{"user": user})
	if err != nil {
		return nil, err
	}
	var res []*deviceproto.AddDeviceResponse
	for resp.Next(u.C) {
		var all deviceproto.AddDeviceResponse
		if err := resp.Decode(&all); err != nil {
			return nil, err
		}
		res = append(res, &all)
	}
	return &deviceproto.GetdevicesResponse{Devices: res}, nil
}

func (u *Mongo) Find(req string) bool {
	var res models.Device
	if err := u.D.FindOne(u.C, bson.M{"name": req}).Decode(&res); err != nil {
		if err == mongo.ErrNoDocuments {
			return false
		}
		log.Println("Error finding device:", err)
		return false
	}
	return true
}

func (u *Mongo) Returndevice(req string)(*deviceproto.AddDeviceResponse,error){
	var res deviceproto.AddDeviceResponse
	if err:=u.D.FindOne(u.C,bson.M{"name":req}).Decode(&res);err!=nil{
		return nil,err
	}
	return &res,nil
}