package redisserver

import (
	"api/internal/models"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type Redis struct {
	R *redis.Client
	C context.Context
}

func (u *Redis) RegisterRequest(req *models.RegistrationRequest) error {
	req1, err := json.Marshal(req)
	if err != nil {
		return err
	}
	if err := u.R.Set(u.C, req.Email, req1, 0).Err(); err != nil {
		return err
	}
	return nil
}

func (u *Redis) RegisterResponse(req string) (*models.RegistrationRequest, error) {
	var res models.RegistrationRequest
	val, err := u.R.Get(u.C, req).Result()
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal([]byte(val), &res); err != nil {
		return nil, err
	}
	fmt.Printf("user %v\n age: %v\n email: %v\n password: %v\n",res.Username,res.Age,res.Email,res.Password)
	return &res, nil
}

func (u *Redis) VerifyRequest(code, email string) error {
	req := models.VerifyRequest{Code: code, Email: email}
	if err := u.R.Set(u.C, req.Email+"1", req.Code, 5*time.Minute).Err(); err != nil {
		return err
	}
	return nil
}

func (u *Redis) VerifyResponse(email string) (string, error) {
	var res string
	if err := u.R.Get(u.C, email+"1").Scan(&res); err != nil {
		return ``, err
	}
	return res, nil
}

func (u *Redis) SaveUser(req *models.UserProfileResponse) error {
	email := req.Email
	req1, err := json.Marshal(req)
	if err != nil {
		return err
	}
	if err := u.R.Set(u.C, email+"profile", req1, 0).Err(); err != nil {
		return err
	}
	return nil
}

func (u *Redis) Getuser(req string) (*models.UserProfileResponse, error) {
	var res models.UserProfileResponse
	val, err := u.R.Get(u.C, req+"profile").Result()
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal([]byte(val), &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (u *Redis) UpdateUser(req *models.UserProfileResponse) error {
	if err := u.SaveUser(req); err != nil {
		return err
	}
	return nil
}

func (u *Redis) SaveDevices(dname string, req []byte) error {
	if err := u.R.Set(u.C, dname, req, 0).Err(); err != nil {
		return err
	}
	return nil
}

func (u *Redis) GetDevice(dname string) (string, error) {
	val, err := u.R.Get(u.C, dname).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}
