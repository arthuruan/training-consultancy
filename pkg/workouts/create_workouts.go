package workouts

import (
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
	Series        string `json:"series" validate:"required"`
	Repetition    string `json:"repetition" validate:"required"`
	Duration      int64  `json:"duration"`
	Rest          int64  `json:"rest" validate:"required"`
	Load          string `json:"load" validate:"required"`
	Method        string `json:"method" validate:"required"`
}

type CreateWorkoutsBody struct {
	WorkoutSheetID string        `json:"workoutSheetID" validate:"required"`
	Workouts       []WorkoutBody `json:"workouts" validate:"required"`
}

func (h handler) CreateWorkouts(ctx *gin.Context) {
	body := CreateWorkoutsBody{}

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

	var workouts []models.Workout
	var documents []interface{}

	for _, item := range body.Workouts {
		// Validate Workout
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

		workoutSheetId, _ := primitive.ObjectIDFromHex(body.WorkoutSheetID)

		workout := models.Workout{
			ID:             primitive.NewObjectID(),
			WorkoutSheetID: workoutSheetId,
			ExerciseID:     exerciseId,
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
			{Key: "_id", Value: workout.ID},
			{Key: "workoutSheetId", Value: workout.WorkoutSheetID},
			{Key: "exerciseId", Value: workout.ExerciseID},
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

	if _, err := h.workoutsCollection.InsertMany(ctx, documents); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errorMessage": "Failed to register the workouts.",
		})
		return
	}

	ctx.JSON(http.StatusCreated, workouts)
}
