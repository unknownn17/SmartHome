package speaker

import (
	"device/internal/connections"
	"device/internal/database/command"
	"device/internal/models"
	"device/protos/deviceproto"
	"errors"
	"fmt"
)

type Speaker struct {
	C *command.Command
}

func NewSpeaker() *Speaker {
	b := connections.DeviceCommand()
	return &Speaker{C: b}
}

var songs = []string{"ILTIJO", "SHAYTANAT", "ARZIMASKAN", "NO LOVE", "RAP ELECTROSHOCK", "ANDIJON UCHUN", "KECHIRING OYIJON", "FACT", "VASSABI", "DEZDEMONA", "YOLG'IZLIK"}

func (u *Speaker) Commands(req *deviceproto.SpeakerCommandRequest) error {
	fmt.Println(req)
	if req.TurnCommand == "" {
		return errors.New("there is no command")
	}
	if req.TurnCommand != "ON" && req.TurnCommand != "OFF" && req.TurnCommand != "U" {
		return errors.New("unknown command")
	}

	switch req.TurnCommand {
	case "ON":
		if err := u.AddOne(req.Dtype); err != nil {
			return err
		}
		return nil
	case "OFF":
		return u.turnOff(req.Dtype)
	case "U":
		return u.processDoneCommand(req)
	}

	return nil
}

func (u *Speaker) AddOne(req string) error {
	_, err := u.C.SFind(req)
	if err != nil {
		updaterequest := models.SpeakerCommand{
			Dname:  req,
			Song:   songs[0],
			SongId: 0,
			Turn:   "ON",
			Volume: 5,
		}
		if err := u.C.SAddCommand(&updaterequest); err != nil {
			return err
		}
		return nil
	}
	updaterequest := models.SpeakerCommand{
		Dname:  req,
		Song:   songs[0],
		SongId: 0,
		Turn:   "ON",
		Volume: 5,
	}
	if err := u.C.SUpdate(&updaterequest); err != nil {
		return err
	}
	return nil
}

func (u *Speaker) turnOff(dname string) error {
	updaterequest := models.SpeakerCommand{
		Dname:  dname,
		Song:   songs[0],
		SongId: 0,
		Turn:   "OFF",
		Volume: 5,
	}
	if err := u.C.SUpdate(&updaterequest); err != nil {
		return err
	}
	return nil
}

func (u *Speaker) processDoneCommand(req *deviceproto.SpeakerCommandRequest) error {
	updaterequest := models.SpeakerCommand{
		Dname: req.Dtype,
		Turn:  req.TurnCommand,
	}

	if req.Control != "" {
		id, err := u.SongCommand(req.Dtype, req.Control)
		if err != nil {
			return err
		}
		updaterequest.SongId = id
		updaterequest.Song = songs[id]
	}

	if req.VolumeCommand != "" {
		vol, err := u.VolumeCommand(req.Dtype, req.VolumeCommand)
		if err != nil {
			return err
		}
		updaterequest.Volume = vol
	}

	if req.AddSong != "" {
		songs = append(songs, req.AddSong)
	}

	fmt.Printf("before updating %v\n", updaterequest)
	if err := u.C.SUpdate(&updaterequest); err != nil {
		return errors.New("updating error")
	}
	fmt.Println("update successful")
	return nil
}

func (u *Speaker) SongCommand(req, command string) (int, error) {
	res, err := u.C.SFind(req)
	fmt.Printf("result %v\nreq %v\n", res, req)
	if err != nil {
		return -1, err
	}

	switch command {
	case "NEXT":
		if res.SongId == len(songs)-1 {
			return 0, nil
		}
		return res.SongId + 1, nil
	case "PREVIOUS":
		if res.SongId == 0 {
			return len(songs) - 1, nil
		}
		return res.SongId - 1, nil
	default:
		return res.SongId, nil
	}
}

func (u *Speaker) VolumeCommand(req, command string) (int, error) {
	res, err := u.C.SFind(req)
	if err != nil {
		return -1, err
	}

	switch command {
	case "UP":
		if res.Volume < 7 {
			return res.Volume + 1, nil
		}
		return res.Volume, nil
	case "DOWN":
		if res.Volume > 0 {
			return res.Volume - 1, nil
		}
		return res.Volume, nil
	default:
		return res.Volume, nil
	}
}
