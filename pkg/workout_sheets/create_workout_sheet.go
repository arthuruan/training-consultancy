package workoutsheets

import (
	"net/http"
	"time"

	"github.com/arthuruan/training-consultancy/common/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateWorkoutSheetBody struct {
	StudentID      string `json:"studentId" validate:"required"`
	Type           string `json:"type" validate:"required"`
	StartTimestamp string `json:"startTimestamp" validate:"required"`
	EndTimestamp   string `json:"endTimestamp" validate:"required"`
	Observation    string `json:"observation"`
}

func (h handler) CreateWorkoutSheet(ctx *gin.Context) {
	body := CreateWorkoutSheetBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errorMessage": err.Error(),
		})
		return
	}

	// Validate body
	var validate = validator.New()
	if err := validate.Struct(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errorMessage": err.Error(),
		})
		return
	}

	// Validate StudentID
	var student models.Student
	studentId, _ := primitive.ObjectIDFromHex(body.StudentID)
	if err := h.usersCollection.FindOne(ctx, bson.M{"_id": studentId, "type": models.UserType.Student}).Decode(&student); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"errorMessage": "StudentId was not found.",
		})
		return
	}

	startTimestamp, err := time.Parse(time.RFC3339, body.StartTimestamp)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errorMessage": "Invalid startTimestamp format, you should use ISO 8601 format.",
		})
		return
	}

	endTimestamp, err := time.Parse(time.RFC3339, body.EndTimestamp)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errorMessage": "Invalid endTimestamp format, you should use ISO 8601 format.",
		})
		return
	}

	// Create Student
	workoutSheet := models.WorkoutSheet{
		ID:             primitive.NewObjectID(),
		PersonalID:     student.PersonalID,
		StudentID:      body.StudentID,
		Type:           body.Type,
		StartTimestamp: startTimestamp,
		EndTimestamp:   endTimestamp,
		Observation:    body.Observation,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	if _, err := h.workoutSheetsCollection.InsertOne(ctx, workoutSheet); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errorMessage": "Failed to register the workout sheet.",
		})
		return
	}

	ctx.JSON(http.StatusCreated, workoutSheet)
}
