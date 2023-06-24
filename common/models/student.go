package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Student struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	ProfessorID   string             `bson:"professorId" json:"professorId"`
	Name          string             `bson:"name" json:"name"`
	BirthDate     time.Time          `bson:"birthDate" json:"birthDate"`
	Objective     string             `bson:"objective" json:"objective"`
	Gender        string             `bson:"gender" json:"gender"`
	Email         string             `bson:"email" json:"email"`
	Phone         string             `bson:"phone" json:"phone"`
	PlanType      string             `bson:"planType" json:"planType"`
	Frequence     string             `bson:"frequence" json:"frequence"`
	TrainingPlace string             `bson:"trainingPlace" json:"trainingPlace"`
	CreatedAt     time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt     time.Time          `bson:"updatedAt" json:"updateAt"`
}
