package workouts

import (
	"net/http"

	"github.com/arthuruan/training-consultancy/common/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (h handler) GetWorkouts(ctx *gin.Context) {
	var workouts []models.Workout

	filters := []primitive.E{}
	if workoutSheetId := ctx.Query("workoutSheetId"); workoutSheetId != "" {
		objId, _ := primitive.ObjectIDFromHex(workoutSheetId)
		filters = append(filters, primitive.E{Key: "workoutSheetId", Value: objId})
	}

	cursor, err := h.workoutsCollection.Find(ctx, filters)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"errorMessage": err.Error(),
		})
		return
	}

	defer cursor.Close(ctx)
	if err := cursor.All(ctx, &workouts); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errorMessage": "Failed to list workouts.",
		})
		return
	}

	ctx.JSON(http.StatusOK, workouts)
}
