package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `json:"id,omitempty"`
	Email    string             `json:"email"`
	Password string             `json:"password"`
	Type     string             `json:"type"`
	Name     string             `json:"name"`
}
