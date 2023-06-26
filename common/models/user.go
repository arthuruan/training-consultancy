package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var UserType = struct {
	Personal string
	Student  string
}{
	Personal: "personal",
	Student:  "student",
}

type User struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Email         string             `bson:"email,omitempty" json:"email,omitempty"`
	Password      string             `bson:"password,omitempty" json:"-"`
	Type          string             `bson:"type,omitempty" json:"type,omitempty"`
	Name          string             `bson:"name,omitempty" json:"name,omitempty"`
	PersonalID    string             `bson:"personalId,omitempty" json:"personalId,omitempty"`
	Birthday      time.Time          `bson:"birthday,omitempty" json:"birthday,omitempty"`
	Objective     string             `bson:"objective,omitempty" json:"objective,omitempty"`
	Gender        string             `bson:"gender,omitempty" json:"gender,omitempty"`
	Phone         string             `bson:"phone,omitempty" json:"phone,omitempty"`
	PlanType      string             `bson:"planType,omitempty" json:"planType,omitempty"`
	Frequence     string             `bson:"frequence,omitempty" json:"frequence,omitempty"`
	TrainingPlace string             `bson:"trainingPlace,omitempty" json:"trainingPlace,omitempty"`
	CreatedAt     time.Time          `bson:"createdAt,omitempty" json:"createdAt,omitempty"`
	UpdatedAt     time.Time          `bson:"updatedAt,omitempty" json:"updatedAt,omitempty"`
}

type Student struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Email         string             `bson:"email,omitempty" json:"email,omitempty"`
	Type          string             `bson:"type,omitempty" json:"type,omitempty"`
	Name          string             `bson:"name,omitempty" json:"name,omitempty"`
	PersonalID    primitive.ObjectID `bson:"personalId,omitempty" json:"personalId,omitempty"`
	Birthday      time.Time          `bson:"birthday,omitempty" json:"birthday,omitempty"`
	Objective     string             `bson:"objective,omitempty" json:"objective,omitempty"`
	Gender        string             `bson:"gender,omitempty" json:"gender,omitempty"`
	Phone         string             `bson:"phone,omitempty" json:"phone,omitempty"`
	PlanType      string             `bson:"planType,omitempty" json:"planType,omitempty"`
	Frequence     string             `bson:"frequence,omitempty" json:"frequence,omitempty"`
	TrainingPlace string             `bson:"trainingPlace,omitempty" json:"trainingPlace,omitempty"`
	CreatedAt     time.Time          `bson:"createdAt,omitempty" json:"createdAt,omitempty"`
	UpdatedAt     time.Time          `bson:"updatedAt,omitempty" json:"updatedAt,omitempty"`
}

type Personal struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Email     string             `bson:"email,omitempty" json:"email,omitempty"`
	Password  string             `bson:"password,omitempty" json:"-"`
	Type      string             `bson:"type,omitempty" json:"type,omitempty"`
	Name      string             `bson:"name,omitempty" json:"name,omitempty"`
	Birthday  time.Time          `bson:"birthday,omitempty" json:"birthday,omitempty"`
	Gender    string             `bson:"gender,omitempty" json:"gender,omitempty"`
	Phone     string             `bson:"phone,omitempty" json:"phone,omitempty"`
	CreatedAt time.Time          `bson:"createdAt,omitempty" json:"createdAt,omitempty"`
	UpdatedAt time.Time          `bson:"updatedAt,omitempty" json:"updatedAt,omitempty"`
}
