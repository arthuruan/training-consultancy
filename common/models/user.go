package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Email     string             `bson:"email,omitempty" json:"email"`
	Password  string             `bson:"password,omitempty" json:"password"`
	Type      string             `bson:"type,omitempty" json:"type"`
	Name      string             `bson:"name,omitempty" json:"name"`
	CreatedAt time.Time          `bson:"createdAt,omitempty" json:"createdAt"`
}
