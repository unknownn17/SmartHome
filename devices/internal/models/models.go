package models

import "time"

type Device struct {
	User        string `json:"user" bson:"user"`
	Device_name string `json:"name" bson:"name"`
	ListCommand string `json:"commands" bson:"commands"`
	Status      string `json:"status" bson:"status"`
	Color       string `bson:"color"`
}

type CommandsList struct {
	Commands string `json:"Commands" bson:"commands"`
}

type SpeakerCommandUpdate struct {
	Turn   string `bson:"turn"`
	Volume string `bson:"volume"`
	Song   string `bson:"song"`
}

type SpeakerCommand struct {
	Turn   string `json:"turn" bson:"turn"`
	Volume int    `json:"volume" bson:"volume"`
	SongId int    `json:"sonid" bson:"songid"`
	Song   string `json:"song" bson:"song"`
	Dname  string `json:"dname" bson:"dname"`
}

type Vaccum struct {
	Turn_command   string    `bson:"turn"`
	Location       string    `bson:"location"`
	Timer          int       `bson:"timer"`
	Created_at     time.Time `bson:"created_at"`
	Dname          string    `bson:"dname"`
	Remaining_time int       `bson:"remaining_time"`
	Status         string    `bson:"status"`
}

type Alarm struct {
	Lift_curtain   string    `bson:"curtain"`
	Lamp           string    `bson:"lamp"`
	Color          string    `bson:"color"`
	Status         string    `bson:"status"`
	Dname          string    `bson:"dname"`
	Remaining_time int       `bson:"remaining_time"`
	Alarm          int       `bson:"alarm"`
	Created_at     time.Time `bson:"created_at"`
}

type Door struct {
	Door           string    `bson:"door"`
	Command        string    `bson:"command"`
	Status         string    `bson:"status"`
	Dname          string    `bson:"dname"`
	Remaining_time int       `bson:"remaining_time"`
	Created_at     time.Time `bson:"created_at"`
	Timer          int       `bson:"timer"`
}

type Devicelist struct {
	Device_name string `bson:"name"`
	ListCommand string `bson:"commands"`
	Color       string `bson:"color"`
}

type Speaker struct {
	Turn   string `json:"turn"`
	Song   string `json:"song"`
	Volume string `json:"volume"`
	Add    string `json:"add"`
	Dname  string `json:"dname"`
}
