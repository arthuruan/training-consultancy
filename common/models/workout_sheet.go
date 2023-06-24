package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type WorkoutSheet struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	StudentID      string             `bson:"studentId" json:"studentId"`
	ProfessorID    string             `bson:"professorId" json:"professorId"`
	Type           string             `bson:"type" json:"type"`
	StartTimestamp time.Time          `bson:"startTimestamp" json:"startTimestamp"`
	EndTimestamp   time.Time          `bson:"endTimestamp" json:"endTimestamp"`
	Observation    string             `bson:"observation" json:"observation"`
	CreatedAt      time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt      time.Time          `bson:"updatedAt" json:"updatedAt"`
}
