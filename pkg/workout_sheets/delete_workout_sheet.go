package workoutsheets

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (h handler) DeleteWorkoutSheet(ctx *gin.Context) {
	id := ctx.Param("id")
	objId, _ := primitive.ObjectIDFromHex(id)

	result, err := h.workoutSheetsCollection.DeleteOne(ctx, bson.M{"_id": objId})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"errorMessage": "Failed to delete workout sheet.",
		})
		return
	}

	if result.DeletedCount == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"errorMessage": "Workout sheet not found.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Workout sheet successfully deleted!",
	})
}
