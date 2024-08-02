package models

type UserProfile struct {
	Id         int64  `json:"id" bson:"id"`
	Username   string `json:"username" bson:"username"`
	Age        int    `json:"age" bson:"age"`
	Email      string `json:"email" bson:"email"`
	Password   string `json:"password" bson:"password"`
	Created_at string `json:"created_at" bson:"created_at"`
	Updated_at string `json:"updated_at" bson:"updated_at"`
	Logout_at  string `json:"logout_at" bson:"logout_at"`
}
