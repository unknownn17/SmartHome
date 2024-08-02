package device

import (
	"api/internal/connection"
	"api/internal/deviceclient"
	"api/internal/models"
	redisserver "api/internal/redis"
	"api/protos/deviceproto"
	"context"
	"encoding/json"
	"fmt"
)

type Devices struct {
	D deviceproto.DevicesClient
	R *redisserver.Redis
}

func NewDevice() *Devices {
	d := deviceclient.UserClinet()
	r := connection.Redis()
	return &Devices{R: r, D: d}
}

func (u *Devices) Alldevice(req *deviceproto.Deviceslist) ([]models.AllDevices, error) {
	fmt.Println("here")
	res, err := u.D.ListDevices(context.Background(),req)
	if err != nil {
		return nil, err
	}
	fmt.Println("result ",res)
	var res1 []models.AllDevices

	for _, v := range res.Devices {
		var all = models.AllDevices{
			Name:     v.Name,
			Commands: v.Commands,
			Color:    v.Color,
		}
		res1 = append(res1, all)
	}
	return res1, nil
}

func (u *Devices) AddDevice(req *models.AddDeviceRequest) (*models.AddDeviceResponse, error) {
	res, err := u.D.AddDevice(context.Background(), &deviceproto.AddDeviceRequest{Name: req.Name, Email: req.Email})
	if err != nil {
		return nil, err
	}
	return &models.AddDeviceResponse{User: res.User, Name: res.Name, Commands: res.Commands, Status: res.Status, Color: res.Color}, nil
}

func (u *Devices) GetDevices(req *models.GetDevicesRequest) (*models.GetDevicesResponse, error) {
	res, err := u.D.GetDevices(context.Background(), &deviceproto.GetdevicesRequest{User: req.User})
	if err != nil {
		return nil, err
	}
	var res1 []models.AddDeviceResponse
	for _, v := range res.Devices {
		var all = models.AddDeviceResponse{
			User:     v.User,
			Name:     v.Name,
			Commands: v.Commands,
			Color:    v.Color,
			Status:   v.Status,
		}
		res1 = append(res1, all)
	}
	return &models.GetDevicesResponse{Devices: res1}, nil
}

func (u *Devices) DeleteDevice(req *models.DeleteDeviceRequest) (*models.DeleteDeviceResponse, error) {
	res, err := u.D.DeleteDevice(context.Background(), &deviceproto.DeleteDeviceRequest{Name: req.Name})
	if err != nil {
		return nil, err
	}
	return &models.DeleteDeviceResponse{Status: res.Status}, nil
}

func (u *Devices) SpeakerGet(req *models.GetDevice) (*models.SpeakerGet, error) {
	res1, err := u.R.GetDevice(req.Device)
	if err == nil {
		var res models.SpeakerGet
		if err := json.Unmarshal([]byte(res1), &res); err != nil {
			return nil, err
		}
		return &res, nil
	}
	res2, err := u.D.SpeakerGet(context.Background(), &deviceproto.GETDeviceRequest{Device: req.Device})
	if err != nil {
		return nil, err
	}
	var speaker = models.SpeakerGet{
		Turn:   res2.Turn,
		Song:   res2.Song,
		Volume: int(res2.Volume),
		Songid: int(res2.SongId),
		Dname:  res2.Dname,
	}
	req1, err := json.Marshal(speaker)
	if err != nil {
		return nil, err
	}
	if err := u.R.SaveDevices(speaker.Dname, req1); err != nil {
		return nil, err
	}
	return &speaker, nil
}

func (u *Devices) VaccumGet(req *models.GetDevice) (*models.Vaccum, error) {
	res1, err := u.R.GetDevice(req.Device)
	if err == nil {
		var res models.Vaccum
		if err := json.Unmarshal([]byte(res1), &res); err != nil {
			return nil, err
		}
		return &res, nil
	}
	res2, err := u.D.VaccumCleanerGet(context.Background(), &deviceproto.GETDeviceRequest{Device: req.Device})
	if err != nil {
		return nil, err
	}
	var vaccum = models.Vaccum{
		Turn_command:   res2.TurnCommand,
		Location:       res2.Location,
		Timer:          int(res2.Time),
		Remaining_time: int(res2.RemainingTime),
		Dname:          res2.Dname,
		Status:         res2.Status,
	}
	req1, err := json.Marshal(vaccum)
	if err != nil {
		return nil, err
	}
	if err := u.R.SaveDevices(vaccum.Dname, req1); err != nil {
		return nil, err
	}
	return &vaccum, nil
}

func (u *Devices) AlarmGet(req *models.GetDevice) (*models.Alarm, error) {
	res1, err := u.R.GetDevice(req.Device)
	if err == nil {
		var res models.Alarm
		if err := json.Unmarshal([]byte(res1), &res); err != nil {
			return nil, err
		}
		return &res, nil
	}
	res2, err := u.D.SmartAlarmGet(context.Background(), &deviceproto.GETDeviceRequest{Device: req.Device})
	if err != nil {
		return nil, err
	}
	var alarm = models.Alarm{
		Lift_curtain:   res2.LiftingCurtain,
		Lamp:           res2.LightCommand,
		Color:          res2.LightColor,
		Alarm:          int(res2.Setalarm),
		Status:         res2.Status,
		Dname:          res2.Dname,
		Remaining_time: int(res2.RemainingTime),
	}
	req1, err := json.Marshal(alarm)
	if err != nil {
		return nil, err
	}
	if err := u.R.SaveDevices(alarm.Dname, req1); err != nil {
		return nil, err
	}
	return &alarm, nil
}

func (u *Devices) DoorGet(req *models.GetDevice) (*models.Door, error) {
	res1, err := u.R.GetDevice(req.Device)
	if err == nil {
		var res models.Door
		if err := json.Unmarshal([]byte(res1), &res); err != nil {
			return nil, err
		}
		return &res, nil
	}
	res2, err := u.D.GetDoor(context.Background(), &deviceproto.GETDeviceRequest{Device: req.Device})
	if err != nil {
		return nil, err
	}
	var door = models.Door{
		Door:           res2.Door,
		Command:        res2.Command,
		Dname:          res2.Dname,
		Timer:          int(res2.Time),
		Remaining_time: int(res2.RemainingTime),
		Status:         res2.Status,
	}
	req1, err := json.Marshal(door)
	if err != nil {
		return nil, err
	}
	if err := u.R.SaveDevices(door.Dname, req1); err != nil {
		return nil, err
	}
	return &door, nil
}
