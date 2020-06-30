package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID        string               `json:"id,omitempty" bson:"_id,omitempty"`
	Email     string               `json:"email" bson:"email"`
	UserName  string               `json:"username" bson:"username"`
	Password  string               `json:"-" bson:"password,omitempty"`
	Posts     []primitive.ObjectID `json:"posts,omitempty" bson:"posts"`
	Following []primitive.ObjectID `json:"following,omitempty" bson:"following"`
	Followers []primitive.ObjectID `json:"followers,omitempty" bson:"followers"`
}

type Author struct {
	ID       string `json:"id,omitempty" bson:"_id,omitempty"`
	UserName string `json:"username" bson:"username"`
}
