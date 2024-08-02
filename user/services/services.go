package service

import (
	"context"
	"fmt"
	"user/internal/mongodb"
	"user/protos/userproto"
)

type Service struct {
	userproto.UnimplementedUserServer
	R *mongodb.Mongo
}

	func (u *Service) Register(ctx context.Context, req *userproto.RegistrationRequest) (*userproto.RegistrationResponse, error) {
		fmt.Printf("user %v\nage %v\n email %v\npassword %v\n",req.Username,req.Age,req.Email,req.Password)
		if err := u.R.AddUser(req); err != nil {
			return nil, err
		}
		return &userproto.RegistrationResponse{Message: "done"}, nil
	}

func (u *Service) UpdateUser(ctx context.Context, req *userproto.UpdateRequest) (*userproto.UserProfileResponse, error) {
	_, err := u.R.Update(req)
	if err != nil {
		return nil, err
	}
	res, err := u.R.Findone(req.Email)
	if err != nil {
		return nil, err
	}
	return &userproto.UserProfileResponse{Id: res.Id, Username: res.Username, Age: int32(res.Age), Email: res.Email, CreatedAt: res.Created_at, UpdatedAt: res.Updated_at}, nil
}

func (u *Service) Find(ctx context.Context, req *userproto.CheckRequest) (*userproto.CheckResponse, error) {
	ok := u.R.Find(req)
	if !ok {
		return &userproto.CheckResponse{Check: false}, nil
	}
	return &userproto.CheckResponse{Check: true}, nil
}

func (u *Service) UserProfile(ctx context.Context, req *userproto.UserProfileRequest) (*userproto.UserProfileResponse, error) {
	res, err := u.R.Findone(req.Email)
	if err != nil {
		return nil, err
	}
	return &userproto.UserProfileResponse{Id: res.Id, Username: res.Username, Age: int32(res.Age), Email: res.Email, CreatedAt: res.Created_at, UpdatedAt: res.Updated_at}, nil
}

func (u *Service) Logout(ctx context.Context, req *userproto.LogoutRequest) (*userproto.LogoutResponse, error) {
	if err := u.R.Logout(req); err != nil {
		return nil, err
	}
	var res = userproto.LogoutResponse{Message: "Logged out successfully"}
	return &res, nil
}

func (u *Service) DeleteUser(ctx context.Context, req *userproto.LogoutRequest) (*userproto.LogoutResponse, error) {
	if err := u.R.Delete(req); err != nil {
		return nil, err
	}
	var res = userproto.LogoutResponse{Message: "Deleted successfully"}
	return &res, nil
}

func (u *Service) LogIn(ctx context.Context,req *userproto.LoginRequest) (*userproto.LoginResponse, error){
	ok:=u.R.Login(req)
	if !ok{
		return &userproto.LoginResponse{Message: "Password doesn't match"},nil
	}
	return &userproto.LoginResponse{Message:"You are logged in"},nil
}