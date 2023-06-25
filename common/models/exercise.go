package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Exercise struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name      string             `bson:"name,omitempty" json:"name,omitempty"`
	Category  string             `bson:"category,omitempty" json:"category,omitempty"`
	VideoURL  string             `bson:"videoUrl,omitempty" json:"videoUrl,omitempty"`
	CreatedAt time.Time          `bson:"createdAt,omitempty" json:"createdAt,omitempty"`
	UpdatedAt time.Time          `bson:"updatedAt,omitempty" json:"updatedAt,omitempty"`
}
