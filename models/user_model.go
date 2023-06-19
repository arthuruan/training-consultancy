package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `json:"_id,omitempty"`
	Email    string             `json:"email" validate:"required"`
	Password string             `json:"password" validate:"required"`
	Type     string             `json:"type" validate:"required"`
	Name     string             `json:"name" validate:"required"`
}
