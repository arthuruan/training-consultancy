package workoutsheets

import (
	"fmt"
	"net/http"
	"time"

	"github.com/arthuruan/training-consultancy/common/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type WorkoutBody struct {
	ExerciseID    string `json:"exerciseId" validate:"required"`
	Type          string `json:"type" validate:"required"`
	SequenceIndex int64  `json:"sequenceIndex" validate:"required"`
	Series        int64  `json:"series" validate:"required"`
	Repetition    string `json:"repetition" validate:"required"`
	Duration      int64  `json:"duration"`
	Rest          int64  `json:"rest" validate:"required"`
	Load          string `json:"load" validate:"required"`
	Method        string `json:"method" validate:"required"`
}

type CreateWorkoutSheetBody struct {
	StudentID      string        `json:"studentId" validate:"required"`
	Type           string        `json:"type" validate:"required"`
	StartTimestamp string        `json:"startTimestamp" validate:"required"`
	EndTimestamp   string        `json:"endTimestamp" validate:"required"`
	Observation    string        `json:"observation"`
	Workouts       []WorkoutBody `json:"workouts" validate:"required"`
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
		fmt.Println(err.Error())
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

	if endTimestamp.Unix() < startTimestamp.Unix() {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errorMessage": "End timestamp should be gratter than start timestamp.",
		})
		return
	}

	// Create Workout Sheet
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

	var workouts []models.Workout
	var documents []interface{}
	for _, item := range body.Workouts {
		// Validate workout body
		if err := validate.Struct(&item); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"errorMessage": err.Error(),
			})
			return
		}
		// Validate ExerciseID
		var exercise models.Exercise
		exerciseId, _ := primitive.ObjectIDFromHex(item.ExerciseID)
		if err := h.exercisesCollection.FindOne(ctx, bson.M{"_id": exerciseId}).Decode(&exercise); err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"errorMessage": "ExerciseId was not found.",
			})
			return
		}

		workout := models.Workout{
			ID:             primitive.NewObjectID(),
			WorkoutSheetID: workoutSheet.ID.String(),
			ExerciseID:     item.ExerciseID,
			Type:           item.Type,
			SequenceIndex:  item.SequenceIndex,
			Series:         item.Series,
			Repetition:     item.Repetition,
			Duration:       item.Duration,
			Rest:           item.Rest,
			Load:           item.Load,
			Method:         item.Method,
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
		}
		workouts = append(workouts, workout)

		document := bson.D{
			{Key: "id", Value: workout.ID},
			{Key: "workoutSheetId", Value: workout.WorkoutSheetID},
			{Key: "exerciseID", Value: workout.ExerciseID},
			{Key: "type", Value: workout.Type},
			{Key: "sequenceIndex", Value: workout.SequenceIndex},
			{Key: "series", Value: workout.Series},
			{Key: "repetition", Value: workout.Repetition},
			{Key: "duration", Value: workout.Duration},
			{Key: "rest", Value: workout.Rest},
			{Key: "load", Value: workout.Load},
			{Key: "method", Value: workout.Method},
			{Key: "createdAt", Value: workout.CreatedAt},
			{Key: "updatedAt", Value: workout.UpdatedAt},
		}
		documents = append(documents, document)
	}

	if _, err := h.workoutSheetsCollection.InsertOne(ctx, workoutSheet); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errorMessage": "Failed to register the workout sheet.",
		})
		return
	}

	if _, err := h.workoutsCollection.InsertMany(ctx, documents); err != nil {
		h.workoutSheetsCollection.DeleteOne(ctx, bson.M{"_id": workoutSheet.ID})
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errorMessage": "Failed to register the workouts.",
		})
		return
	}

	response := models.WorkoutSheetResponse{
		ID:             primitive.NewObjectID(),
		PersonalID:     student.PersonalID,
		StudentID:      body.StudentID,
		Type:           body.Type,
		StartTimestamp: startTimestamp,
		EndTimestamp:   endTimestamp,
		Observation:    body.Observation,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
		Workouts:       workouts,
	}

	ctx.JSON(http.StatusCreated, response)
}
