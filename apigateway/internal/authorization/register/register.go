package register1

import (
	email1 "api/internal/authorization/email"
	jwttoken "api/internal/authorization/jwt"
	"api/internal/connection"
	"api/internal/models"
	redisserver "api/internal/redis"
	"api/internal/useclient"
	"api/protos/userproto"
	"context"
	"errors"
	"fmt"
)

type Register struct {
	R *redisserver.Redis
	U userproto.UserClient
	C context.Context
	Can context.CancelFunc
}

func NewReg() *Register {
	s := connection.Redis()
	u,c,c1 := useclient.UserClinet()
	return &Register{R: s, U: u,C: c,Can: c1}
}

func (u *Register) Register(req *models.RegistrationRequest) error {
	if req.Username == "" || req.Password == "" || req.Email == "" || req.Age == 0 {
		return errors.New("missing arguments")
	}
	ok, err := u.U.Find(u.C, &userproto.CheckRequest{Email: req.Email})
	if err != nil {
		return err
	}
	if ok.Check {
		return errors.New("user already existed")
	}
	if err := u.R.RegisterRequest(req); err != nil {
		return err
	}
	if err := u.Email(req.Email); err != nil {
		return err
	}
	return nil
}

func (u *Register) Verify(req *models.VerifyRequest) error {
	if req.Email == "" || req.Code == "" {
		return errors.New("missing parameter")
	}
	if err := u.R.VerifyRequest(req.Code, req.Email); err != nil {
		return err
	}
	return nil
}

func (u *Register) Email(to string) error {
	code := email1.Sent(to)
	if err := u.Verify(&models.VerifyRequest{Code: code, Email: to}); err != nil {
		return err
	}
	return nil
}

func (u *Register) UserVerify(req *models.VerifyRequest) error {
	res, err := u.Secretcode(req.Email)
	if err != nil {
		return err
	}
	if res != req.Code {
		return errors.New("password doesn't match")
	}
	return nil
}

func (u *Register) Secretcode(email string) (string, error) {
	code, err := u.R.VerifyResponse(email)
	if err != nil {
		return "", err
	}
	return code, nil
}

func (u *Register) Sendtodatabase(req string) error {
	res, err := u.R.RegisterResponse(req)
	if err != nil {
		return err
	}
	fmt.Println("redis responseeee ",res)
	a:=userproto.RegistrationRequest{
		Username: res.Username,
		Age:      res.Age,
		Email:    res.Email,
		Password: res.Password}
		fmt.Println(a.Age)
	_, err = u.U.Register(u.C, &a)
	if err != nil {
		return err
	}
	fmt.Println(a.Username)
	return nil
}

func (u *Register) Login(req *models.LoginRequest) (string, error) {
	res, err := u.R.RegisterResponse(req.Email)
	if err != nil {
		return "", err
	}
	if res.Password != req.Password {
		return "", errors.New("password isn't match")
	}
	token, err := jwttoken.CreateToken(res)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (u *Register) UserProfile(req *models.UserProfileRequest) (*models.UserProfileResponse, error) {
	res, err := u.R.Getuser(req.Email)
	if err == nil {
		return res, nil
	}
	res1, err := u.U.UserProfile(u.C, &userproto.UserProfileRequest{Email: req.Email})
	if err != nil {
		return nil, err
	}
	fmt.Println("user service get ", res)
	var userprof = models.UserProfileResponse{
		Id:        res1.Id,
		Username:  res1.Username,
		Age:       res1.Age,
		Email:     res1.Email,
		CreatedAt: res1.CreatedAt,
		UpdatedAt: res1.UpdatedAt,
		LogOutAt:  res1.LogOutAt,
	}
	fmt.Println("struct", userprof)
	if err := u.R.SaveUser(&userprof); err != nil {
		return nil, nil
	}
	fmt.Println(userprof)
	return &userprof, nil

}

func (u *Register) Updaterequest(req *models.UpdateRequest) (*models.UserProfileResponse, error) {
	req1 := userproto.UpdateRequest{
		Id:       req.Id,
		Username: req.Username,
		Age:      req.Age,
		Email:    req.Email,
		Password: req.Password,
	}
	res, err := u.U.UpdateUser(u.C, &req1)
	if err != nil {
		return nil, err
	}
	var res1 = models.UserProfileResponse{
		Id:        res.Id,
		Username:  res.Username,
		Age:       res.Age,
		Email:     res.Email,
		CreatedAt: res.CreatedAt,
		UpdatedAt: res.LogOutAt,
	}
	return &res1, nil
}

func (u *Register) Logout(req *models.LogoutRequest)(*models.LogoutResponse,error){
	res,err:=u.U.Logout(u.C,&userproto.LogoutRequest{Email: req.Email})
	if err!=nil{
		return nil,err
	}
	return &models.LogoutResponse{Message: res.Message},nil
}

func (u *Register) Delete(req *models.LogoutRequest)(*models.LogoutResponse,error){
	res,err:=u.U.DeleteUser(u.C,&userproto.LogoutRequest{Email: req.Email})
	if err!=nil{
		return nil,err
	}
	return &models.LogoutResponse{Message: res.Message},nil
}