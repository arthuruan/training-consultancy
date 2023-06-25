package workoutsheets

import (
	"net/http"

	"github.com/arthuruan/training-consultancy/common/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func (h handler) GetWorkoutSheets(ctx *gin.Context) {
	var workoutSheets []models.WorkoutSheet

	cursor, err := h.workoutSheetsCollection.Find(ctx, bson.D{})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"errorMessage": err.Error(),
		})
		return
	}

	defer cursor.Close(ctx)
	if err := cursor.All(ctx, &workoutSheets); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errorMessage": "Failed to list workout sheets.",
		})
		return
	}

	ctx.JSON(http.StatusOK, workoutSheets)
}
