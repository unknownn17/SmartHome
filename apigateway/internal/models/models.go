package models

import (
	"encoding/json"
	"time"
)

type Message struct {
	Type    string          `json:"type"`
	Payload json.RawMessage `json:"payload"`
}

type Speaker struct {
	Turn   string `json:"turn"`
	Song   string `json:"song"`
	Volume string `json:"volume"`
	Add    string `json:"add"`
	Dname  string `json:"dname"`
}

type Vaccum struct {
	Turn_command   string    `json:"turn"`
	Location       string    `json:"location"`
	Timer          int       `json:"timer"`
	Created_at     time.Time `bson:"created_at"`
	Dname          string    `json:"dname"`
	Remaining_time int       `json:"remaining_time"`
	Status         string    `json:"status"`
}

type Alarm struct {
	Lift_curtain   string    `json:"curtain"`
	Lamp           string    `json:"lamp"`
	Color          string    `json:"color"`
	Status         string    `json:"status"`
	Dname          string    `json:"dname"`
	Remaining_time int       `json:"remaining_time"`
	Alarm          int       `json:"alarm"`
	Created_at     time.Time `json:"created_at"`
}

type Door struct {
	Door           string    `json:"door"`
	Command        string    `json:"command"`
	Status         string    `json:"status"`
	Dname          string    `json:"dname"`
	Remaining_time int       `json:"remaining_time"`
	Created_at     time.Time `json:"created_at"`
	Timer          int       `json:"timer"`
}

type RegistrationRequest struct {
	Username string `redis:"usernanem" json:"username"`
	Age      int32  `redis:"age" json:"age"`
	Email    string `redis:"email" json:"email"`
	Password string `redis:"password" json:"password"`
}

type RegistrationResponse struct {
	Message string `json:"message"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Message string `json:"message"`
}

type UserProfileRequest struct {
	Email string `json:"email"`
}

type UserProfileResponse struct {
	Id        int64  `json:"id"`
	Username  string `json:"username"`
	Age       int32  `json:"age"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	LogOutAt  string `json:"log_out_at"`
}

type UpdateRequest struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Age      int32  `json:"age"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LogoutRequest struct {
	Email string `json:"email"`
}

type LogoutResponse struct {
	Message string `json:"message"`
}

type CheckRequest struct {
	Email string `json:"email"`
}

type CheckResponse struct {
	Check bool `json:"check"`
}

type VerifyRequest struct {
	Email string `json:"email" redis:"email"`
	Code  string `json:"code" redis:"code"`
}

type VerifyResponse struct {
	Message string `json:"verify"`
}

type AddDeviceRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type AddDeviceResponse struct {
	User     string `json:"user"`
	Name     string `json:"name"`
	Commands string `json:"commands"`
	Status   string `json:"status"`
	Color    string `json:"color"`
}

type DeleteDeviceRequest struct {
	Name string `json:"name"`
}

type DeleteDeviceResponse struct {
	Status string `json:"status"`
}

type GetDevicesRequest struct {
	User string `json:"user"`
}

type GetDevicesResponse struct {
	Devices []AddDeviceResponse `json:"devices"`
}

type GetDevice struct {
	Device string `json:"device"`
}

type SpeakerGet struct {
	Turn   string `json:"turn"`
	Song   string `json:"song"`
	Songid int    `json:"songid"`
	Volume int    `json:"volume"`
	Add    string `json:"add"`
	Dname  string `json:"dname"`
}

type AllDevices struct {
	Name     string `json:"name"`
	Commands string `json:"commands"`
	Color    string `json:"color"`
}
