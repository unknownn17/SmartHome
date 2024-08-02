package methods

import (
	"context"
	"device/internal/connections"
	"device/internal/database/devicedatabase"
	database "device/internal/database/userdevice"
	"device/internal/models"
	user "device/internal/userclient"
	"device/protos/deviceproto"
	"device/protos/userproto"
	"fmt"
)

type Methods struct {
	M *database.Mongo
	U userproto.UserClient
	D *devicedatabase.LIST
}

func NewMethods() *Methods {
	a := connections.DeviceMongo()
	b := user.UserClinet()
	c := connections.DeviceLIST()
	return &Methods{M: a, U: b, D: c}
}

func (u *Methods) Adddevice(req *deviceproto.AddDeviceRequest) error {
	a := userproto.UserProfileRequest{Email: req.Email}
	res, err := u.U.UserProfile(context.Background(), &a)
	if err != nil {
		return err
	}
	dev, err := u.D.Findone(req.Name)
	if err != nil {
		return err
	}
	if res.LogOutAt == "" {
		req1 := models.Device{
			User:        res.Username,
			Device_name: dev.Device_name,
			ListCommand: dev.ListCommand,
			Status:      "added",
			Color:       dev.Color,
		}
		if err := u.M.AddDevice(&req1); err != nil {
			return err
		}
		return nil
	}
	return nil
}

func (u *Methods) AllDevices(req *deviceproto.Deviceslist) (*deviceproto.DevicesTop, error) {
	res, err := u.D.FindAll()
	if err != nil {
		return nil, err
	}
	var res1 []*deviceproto.Devicesl

	for _, v := range res {
		var all = deviceproto.Devicesl{
			Name:     v.Device_name,
			Commands: v.ListCommand,
			Color:    v.Color,
		}
		res1 = append(res1, &all)
	}
	fmt.Println("alldevices",res1)
	return &deviceproto.DevicesTop{Devices: res1}, nil
}
