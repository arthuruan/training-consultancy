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

type UpdateWorkoutSheetBody struct {
	Type           string `json:"type" validate:"required"`
	StartTimestamp string `json:"startTimestamp" validate:"required"`
	EndTimestamp   string `json:"endTimestamp" validate:"required"`
	Observation    string `json:"observation"`
}

func (h handler) UpdateWorkoutSheet(ctx *gin.Context) {
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

	startTimestamp, err := time.Parse(time.RFC3339, body.StartTimestamp)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errorMessage": "Invalid start timestamp format, you should use ISO 8601 format.",
		})
		return
	}

	endTimestamp, err := time.Parse(time.RFC3339, body.EndTimestamp)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errorMessage": "Invalid start timestamp format, you should use ISO 8601 format.",
		})
		return
	}

	if endTimestamp.Unix() < startTimestamp.Unix() {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errorMessage": "End timestamp should be greater than start timestamp.",
		})
		return
	}

	// Update workout sheet
	update := bson.M{
		"type":           body.Type,
		"startTimestamp": startTimestamp,
		"endTimestamp":   endTimestamp,
		"observation":    body.Observation,
		"updatedAt":      time.Now(),
	}
	result, err := h.workoutSheetsCollection.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": update})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"errorMessage": "Failed to update workout sheet.",
		})
		return
	}

	var updatedWorkoutSheet models.WorkoutSheet
	if result.MatchedCount == 1 {
		if err := h.workoutSheetsCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&updatedWorkoutSheet); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"errorMessage": err.Error(),
			})
			return
		}
	}

	ctx.JSON(http.StatusOK, updatedWorkoutSheet)
}
