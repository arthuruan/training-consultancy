package models

type Program struct {
	Sheet    WorkoutSheet `bson:"sheet,omitempty" json:"sheet,omitempty"`
	Workouts []Workout    `bson:"workouts,omitempty" json:"workouts,omitempty"`
}
