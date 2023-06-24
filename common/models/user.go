package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var UserType = struct {
	Professor string
	Student   string
}{
	Professor: "professor",
	Student:   "student",
}

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Email     string             `bson:"email,omitempty" json:"email"`
	Password  string             `bson:"password,omitempty" json:"password"`
	Type      string             `bson:"type,omitempty" json:"type"`
	Name      string             `bson:"name,omitempty" json:"name"`
	CreatedAt time.Time          `bson:"createdAt,omitempty" json:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt,omitempty" json:"updatedAt"`
}
