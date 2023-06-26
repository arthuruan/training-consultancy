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

type UpdateWorkoutSheetBody struct {
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

func (h handler) UpdateWorkout(ctx *gin.Context) {
	id := ctx.Param("id")
	objId, _ := primitive.ObjectIDFromHex(id)

	body := UpdateWorkoutSheetBody{}

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

	// Validate ExerciseID
	var exercise models.Exercise
	exerciseId, _ := primitive.ObjectIDFromHex(body.ExerciseID)
	if err := h.exercisesCollection.FindOne(ctx, bson.M{"_id": exerciseId}).Decode(&exercise); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"errorMessage": "ExerciseId was not found.",
		})
		return
	}

	// Update workout sheet
	update := bson.M{
		"exerciseId":    body.ExerciseID,
		"type":          body.Type,
		"sequenceIndex": body.SequenceIndex,
		"series":        body.Series,
		"repetition":    body.Repetition,
		"rest":          body.Rest,
		"load":          body.Load,
		"method":        body.Method,
		"duration":      body.Duration,
		"updatedAt":     time.Now(),
	}
	result, err := h.workoutsCollection.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": update})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"errorMessage": "Failed to update workouts.",
		})
		return
	}

	var updatedWorkouts models.Workout
	if result.MatchedCount == 1 {
		if err := h.workoutsCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&updatedWorkouts); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"errorMessage": err.Error(),
			})
			return
		}
	}

	ctx.JSON(http.StatusOK, updatedWorkouts)
}
