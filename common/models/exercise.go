package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Exercise struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name     string             `bson:"name" json:"name"`
	Category string             `bson:"category" json:"category"`
	VideoURL string             `bson:"videoUrl" json:"videoUrl"`
}
