package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var WorkoutType = struct {
	AB   string
	ABC  string
	ABCD string
}{
	AB:   "AB",
	ABC:  "ABC",
	ABCD: "ABCD",
}

type Workout struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	WorkoutSheetID string             `bson:"workoutSheetId,omitempty" json:"workoutSheetId,omitempty"`
	ExerciseID     string             `bson:"exerciseId,omitempty" json:"exerciseId,omitempty"`
	Type           string             `bson:"type,omitempty" json:"type,omitempty"`
	SequenceIndex  int64              `bson:"sequenceIndex,omitempty" json:"sequence,omitempty"`
	Series         int64              `bson:"series,omitempty" json:"series,omitempty"`
	Repetition     string             `bson:"repetition,omitempty" json:"repetition,omitempty"`
	Duration       int64              `bson:"duration,omitempty" json:"duration,omitempty"`
	Rest           int64              `bson:"rest,omitempty" json:"rest,omitempty"`
	Load           string             `bson:"load,omitempty" json:"load,omitempty"`
	Method         string             `bson:"method,omitempty" json:"method,omitempty"`
	CreatedAt      time.Time          `bson:"createdAt,omitempty" json:"createdAt,omitempty"`
	UpdatedAt      time.Time          `bson:"updatedAt,omitempty" json:"updatedAt,omitempty"`
}
