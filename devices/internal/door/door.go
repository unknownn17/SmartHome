package door

import (
	"device/internal/connections"
	"device/internal/database/command"
	"device/internal/models"
	"device/protos/deviceproto"
	"errors"
	"time"
)

type Door struct {
	C *command.Command
}

func NewDoor() *Door {
	b := connections.DeviceCommand()
	return &Door{C: b}
}

func (u *Door) DoorCommands(req *deviceproto.LockDoor) error {
	if req.Status != "ON" && req.Status != "OFF" && req.Status != "U" {
		return errors.New("unsupported status")
	}

	switch req.Status {
	case "ON":
		if err := u.Commands(req); err != nil {
			return err
		}
	case "OFF":
		if err := u.Update(req); err != nil {
			return err
		}
	case "U":
		if err := u.Update(req); err != nil {
			return err
		}
	}

	return nil
}

func (u *Door) Commands(req *deviceproto.LockDoor) error {
	_, err := u.C.Dfind(req.Dname)
	if err == nil {
		if err := u.Update(req); err != nil {
			return err
		}
		return nil
	}
	time1 := time.Duration(int64(req.Time) * int64(time.Minute))
	created := time.Now()
	remaining_time := int(time1.Minutes())

	req1 := models.Door{
		Door:           req.Door,
		Command:        req.Command,
		Dname:          req.Dname,
		Status:         "waiting",
		Timer:          int(time1.Minutes()),
		Remaining_time: remaining_time,
		Created_at:     created,
	}
	if err := u.C.Dadd(&req1); err != nil {
		return err
	}
	return nil
}

func (u *Door) Update(req *deviceproto.LockDoor) error {
	res, err := u.C.Dfind(req.Dname)
	if err != nil {
		return err
	}

	if req.Door != "" {
		res.Door = req.Door
	}
	if req.Command != "" {
		res.Command = req.Command
	}
	if req.Time > 0 {
		time1 := time.Duration(int64(req.Time) * int64(time.Minute))
		elapsed := time.Since(res.Created_at)
		remaining_time := int(time1.Minutes()) - int(elapsed.Minutes())
		if remaining_time < 0 {
			remaining_time = 0
		}
		res.Timer = int(time1.Minutes())
		res.Remaining_time = remaining_time
	}

	if res.Remaining_time == 0 && res.Status == "waiting" {
		res.Status = "done"
	}
	if err := u.C.Dupdate(res); err != nil {
		return err
	}

	return nil
}
