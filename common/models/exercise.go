package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Exercise struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name      string             `bson:"name" json:"name"`
	Category  string             `bson:"category" json:"category"`
	VideoURL  string             `bson:"videoUrl" json:"videoUrl"`
	CreatedAt time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt" json:"updatedAt"`
}
