package mongodb

import (
	"context"
	"crypto/rand"
	"errors"
	"fmt"
	"log"
	"math/big"
	"time"
	"user/internal/models"
	"user/protos/userproto"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type Mongo struct {
	D *mongo.Collection
	C context.Context
}
func (u *Mongo) AddUser(req *userproto.RegistrationRequest) error {
	password := req.Password
	var res = models.UserProfile{
		Id:         generateRandomCode(),
		Username:   req.Username,
		Age:        int(req.Age),
		Email:      req.Email,
		Password:   Hashing(password),
		Created_at: time.Now().Format(time.RFC3339),
		Updated_at: time.Now().Format(time.RFC3339),
		Logout_at:  "",
	}
	if req.Password == "" {
		return errors.New("password hashing error")
	}
	_, err := u.D.InsertOne(u.C, res)

	if err != nil {
		return err
	}
	fmt.Println(res.Password)
	return nil
}

func (u *Mongo) Userprofile(req *userproto.UserProfileRequest) (*models.UserProfile, error) {
	resp := u.D.FindOne(u.C, bson.M{"email": req.Email})
	var res models.UserProfile

	if err := resp.Decode(&res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (u *Mongo) Login(req *userproto.LoginRequest)bool{
	var res models.UserProfile
	if err:=u.D.FindOne(u.C,bson.M{"email":req.Email}).Decode(&res);err!=nil{
		return false
	}
	fmt.Println(res.Password)
	fmt.Println(req.Password)
	return ComparePassword(res.Password,req.Password)
}
func (u *Mongo) Update(req *userproto.UpdateRequest) (*userproto.UserProfileResponse, error) {
	_, err := u.D.UpdateOne(u.C, bson.M{"id": req.Id}, bson.M{"$set": bson.M{
		"username":   req.Username,
		"age":        req.Age,
		"email":      req.Email,
		"password":   Hashing(req.Password),
		"updated_at": time.Now().Format(time.RFC3339)}})
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (u *Mongo) Logout(req *userproto.LogoutRequest) error {
	_, err := u.D.UpdateOne(u.C, bson.M{"email": req.Email}, bson.M{"$set": bson.M{"logout_at": time.Now().Format(time.RFC3339)}})
	if err != nil {
		return err
	}
	return nil
}

func (u *Mongo) Delete(req *userproto.LogoutRequest) error {
	_, err := u.D.DeleteOne(u.C, bson.M{"email": req.Email})
	if err != nil {
		return err
	}
	return nil
}

func (u *Mongo) Find(req *userproto.CheckRequest) bool {
	var res models.UserProfile
	if err := u.D.FindOne(u.C, bson.M{"email": req.Email}).Decode(&res); err != nil {
		return false
	}
	if res.Email!=""{
		return true
	}
	return false
}

func (u *Mongo) Findone(email string) (*models.UserProfile, error) {
	var res models.UserProfile
	resp := u.D.FindOne(u.C, bson.M{"email": email})
	if err := resp.Decode(&res); err != nil {
		return nil, err
	}
	return &res, nil
}


func ComparePassword(hashed,password string)bool{
	if err:=bcrypt.CompareHashAndPassword([]byte(hashed),[]byte(password));err!=nil{
		log.Fatal(err)
		return false
	}
	return true
}

func generateRandomCode() int64 {
	max := big.NewInt(1000000)
	n, err := rand.Int(rand.Reader, max)
	if err != nil {
		return 0
	}
	return n.Int64()
}

func Hashing(password string) string {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return ""
	}
	return string(hashed)
}
