package alarm

import (
	"device/internal/connections"
	"device/internal/database/command"
	"device/internal/models"
	"device/protos/deviceproto"
	"errors"
	"time"
)

type Alarm struct {
	C *command.Command
}

func NewAlarm() *Alarm {
	b := connections.DeviceCommand()
	return &Alarm{C: b}
}

func (u *Alarm) AlarmCommands(req *deviceproto.SmartAlarm) error {
	if req.Status != "ON" && req.Status != "OFF" && req.Status != "U" {
		return errors.New("unknown command")
	}

	switch req.Status {
	case "ON":
		if err := u.AddAlarm(req); err != nil {
			return err
		}
	case "OFF", "U":
		if err := u.UpdateOne(req); err != nil {
			return err
		}
	}

	return nil
}

func (u *Alarm) AddAlarm(req *deviceproto.SmartAlarm) error {
	_, err := u.C.Afind(req.Dname)
	if err != nil {
		time1 := time.Duration(int64(req.Setalarm) * int64(time.Minute))
		created := time.Now()
		remaining_time := int(time1.Minutes())

		newAlarm := models.Alarm{
			Lift_curtain:   req.LiftingCurtain,
			Lamp:           req.LightCommand,
			Color:          req.LightColor,
			Dname:          req.Dname,
			Status:         "ON",
			Alarm:          int(time1.Minutes()),
			Created_at:     created,
			Remaining_time: remaining_time,
		}

		if err := u.C.Aadd(&newAlarm); err != nil {
			return err
		}
		return nil
	}
	if err := u.UpdateOne(req); err != nil {
		return err
	}
	return nil

}

func (u *Alarm) UpdateOne(req *deviceproto.SmartAlarm) error {
	existingAlarm, err := u.C.Afind(req.Dname)
	if err != nil {
		return err
	}

	if req.LiftingCurtain != "" {
		existingAlarm.Lift_curtain = req.LiftingCurtain
	}
	if req.LightCommand != "" {
		existingAlarm.Lamp = req.LightCommand
	}
	if req.LightColor != "" {
		existingAlarm.Color = req.LightColor
	}
	if req.Status != "" {
		existingAlarm.Status = req.Status
	}
	if req.Setalarm > 0 {
		time1 := time.Duration(int64(req.Setalarm) * int64(time.Minute))
		elapsed := time.Since(existingAlarm.Created_at)
		remaining_time := int(time1.Minutes()) - int(elapsed.Minutes())
		if remaining_time < 0 {
			remaining_time = 0
		}
		existingAlarm.Alarm = int(time1.Minutes())
		existingAlarm.Remaining_time = remaining_time
	}else{
		a:=existingAlarm.Alarm
		existingAlarm.Alarm=a
	}

	if existingAlarm.Remaining_time == 0 && existingAlarm.Status == "ON" {
		existingAlarm.Status = "Ringing"
	}

	if err := u.C.Aupdate(existingAlarm); err != nil {
		return err
	}

	return nil
}
