package services

import (
	"context"
	"device/internal/alarm"
	"device/internal/connections"
	"device/internal/database/devicedatabase"
	"device/internal/door"
	"device/internal/speaker"
	"device/internal/vaccum"
	"device/protos/deviceproto"
	"device/services/methods"
	"errors"
)

type Service struct {
	deviceproto.UnimplementedDevicesServer
	S  *speaker.Speaker
	V  *vaccum.Vaccum
	A  *alarm.Alarm
	D  *door.Door
	DL *devicedatabase.LIST
	M  *methods.Methods
}

func NewService() *Service {
	a := speaker.NewSpeaker()
	b := vaccum.NewVaccum()
	c := alarm.NewAlarm()
	d := door.NewDoor()
	e := connections.DeviceLIST()
	f := methods.NewMethods()
	return &Service{S: a, V: b, A: c, D: d, DL: e, M: f}
}

func (u *Service) AddDevice(ctx context.Context, req *deviceproto.AddDeviceRequest) (*deviceproto.AddDeviceResponse, error) {
	if err := u.M.Adddevice(req); err != nil {
		return nil, err
	}
	res, err := u.M.M.Returndevice(req.Name)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (u *Service) ListDevices(ctx context.Context, req *deviceproto.Deviceslist) (*deviceproto.DevicesTop, error) {
	res, err := u.M.AllDevices(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (u *Service) GetDevices(ctx context.Context, req *deviceproto.GetdevicesRequest) (*deviceproto.GetdevicesResponse, error) {
	res, err := u.M.M.ListDevices(req.User)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (u *Service) DeleteDevice(ctxx context.Context, req *deviceproto.DeleteDeviceRequest) (*deviceproto.DeleteDeviceResponse, error) {
	if err := u.M.M.DeleteDevice(req); err != nil {
		return nil, err
	}
	return &deviceproto.DeleteDeviceResponse{Status: "deleted"}, nil
}

func (u *Service) Speaker(ctx context.Context, req *deviceproto.SpeakerCommandRequest) (*deviceproto.Notification, error) {
	if err := u.S.Commands(req); err != nil {
		return nil, err
	}
	return &deviceproto.Notification{Message: req.Dtype}, nil
}

func (u *Service) SpeakerGet(ctx context.Context, req *deviceproto.GETDeviceRequest) (*deviceproto.SpeakerDeviceResponse, error) {
	res, err := u.S.C.SFind(req.Device)
	if err != nil {
		return nil, errors.New("something went wrong")
	}
	return &deviceproto.SpeakerDeviceResponse{Turn: res.Turn, Volume: int32(res.Volume), Song: res.Song, SongId: int32(res.SongId), Dname: res.Dname}, nil
}

func (u *Service) VaccumClenaer(ctx context.Context, req *deviceproto.VaccumCleanaer) (*deviceproto.Notification, error) {
	req.Status = "in progress"
	if err := u.V.Commands(req); err != nil {
		return nil, err
	}
	return &deviceproto.Notification{Message: "Succesfully"}, nil
}

func (u *Service) VaccumCleanerGet(ctx context.Context, req *deviceproto.GETDeviceRequest) (*deviceproto.VaccumCleanaer, error) {
	res, err := u.V.C.Vfind(req.Device)
	if err != nil {
		return nil, err
	}
	return &deviceproto.VaccumCleanaer{Dname: res.Dname, Location: res.Location, TurnCommand: res.Turn_command, Time: int64(res.Timer), CreatedAt: res.Created_at.Unix(), RemainingTime: int64(res.Remaining_time), Status: res.Status}, nil
}

func (u *Service) SmartAlarms(ctx context.Context, req *deviceproto.SmartAlarm) (*deviceproto.Notification, error) {
	if err := u.A.AlarmCommands(req); err != nil {
		return nil, err
	}
	return &deviceproto.Notification{Message: "Succesfully"}, nil
}

func (u *Service) SmartAlarmGet(ctx context.Context, req *deviceproto.GETDeviceRequest) (*deviceproto.SmartAlarm, error) {
	res, err := u.A.C.Afind(req.Device)
	if err != nil {
		return nil, err
	}
	return &deviceproto.SmartAlarm{LiftingCurtain: res.Lift_curtain, LightCommand: res.Lamp, LightColor: res.Color, Dname: res.Dname, Status: res.Status, Setalarm: int32(res.Alarm), RemainingTime: int32(res.Remaining_time)}, nil
}

func (u *Service) Door(ctx context.Context, req *deviceproto.LockDoor) (*deviceproto.Notification, error) {
	if err := u.D.DoorCommands(req); err != nil {
		return nil, err
	}
	return &deviceproto.Notification{Message: "succesfully"}, nil
}

func (u Service) GetDoor(ctx context.Context, req *deviceproto.GETDeviceRequest) (*deviceproto.LockDoor, error) {
	res, err := u.D.C.Dfind(req.Device)
	if err != nil {
		return nil, err
	}
	return &deviceproto.LockDoor{Door: res.Door, Command: res.Command, Status: res.Status, Dname: res.Dname, Time: int32(res.Timer), RemainingTime: int32(res.Remaining_time)}, nil
}
