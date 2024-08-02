package vaccum

import (
	"device/internal/connections"
	"device/internal/database/command"
	"device/internal/models"
	"device/protos/deviceproto"
	"errors"
	"time"
)

type Vaccum struct {
	C *command.Command
}

func NewVaccum() *Vaccum {
	b := connections.DeviceCommand()
	return &Vaccum{C: b}
}

func (u *Vaccum) Commands(req *deviceproto.VaccumCleanaer) error {
	if req.TurnCommand != "ON" && req.TurnCommand != "OFF" && req.TurnCommand != "U" {
		return errors.New("unknown command")
	}

	switch req.TurnCommand {
	case "ON":
		if err := u.AddTask(req); err != nil {
			return err
		}
		return nil
	case "OFF":
		if err := u.UpdateOne(req); err != nil {
			return err
		}
		return nil
	case "U":
		if err := u.UpdateOne(req); err != nil {
			return err
		}
		return nil
	}

	return nil
}

func (u *Vaccum) AddTask(req *deviceproto.VaccumCleanaer) error {
	_, err := u.C.Vfind(req.Dname)
	if err != nil {
		time1 := time.Duration(req.Time * int64(time.Minute))
		created := time.Now()
		remaining_time := int(time1.Minutes())

		var req1 = models.Vaccum{
			Dname:          req.Dname,
			Turn_command:   req.TurnCommand,
			Location:       req.Location,
			Timer:          int(time1.Minutes()),
			Created_at:     created,
			Remaining_time: remaining_time,
			Status:         "ON",
		}

		if err := u.C.Vadd(&req1); err != nil {
			return err
		}
		return nil
	}
	if err := u.UpdateOne(req); err != nil {
		return err
	}
	return nil
}

func (u *Vaccum) UpdateOne(req *deviceproto.VaccumCleanaer) error {
	res, err := u.C.Vfind(req.Dname)
	if err != nil {
		return err
	}

	req1 := res

	if req.Location != "" {
		req1.Location = req.Location
	}

	if req.TurnCommand != "" {
		req1.Turn_command = req.TurnCommand
	}

	if req.Time > 0 {
		time1 := time.Duration(req.Time * int64(time.Minute))
		req1.Timer = int(time1.Minutes())

		elapsed := time.Since(res.Created_at)
		remaining_time := int(time1.Minutes()) - int(elapsed.Minutes())
		if remaining_time < 0 {
			remaining_time = 0
		}
		req1.Remaining_time = remaining_time
	}

	if req1.Remaining_time == 0 {
		req1.Status = "DONE"
	} else {
		req1.Status = "in progress"
	}

	if err := u.C.VUpdate(req1); err != nil {
		return err
	}

	return nil
}
