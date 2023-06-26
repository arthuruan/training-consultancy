package programs

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (h handler) DeleteProgram(ctx *gin.Context) {
	workoutSheetId := ctx.Param("workoutSheetId")
	objId, _ := primitive.ObjectIDFromHex(workoutSheetId)

	result, err := h.workoutSheetsCollection.DeleteOne(ctx, bson.M{"_id": objId})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"errorMessage": "Failed to delete sheet.",
		})
		return
	}

	if result.DeletedCount == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"errorMessage": "Sheet not found.",
		})
		return
	}

	_result, err := h.workoutsCollection.DeleteMany(ctx, bson.M{"workoutSheetId": objId})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"errorMessage": "Failed to delete workouts.",
		})
		return
	}

	if _result.DeletedCount == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"errorMessage": "Workouts not found.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Workout sheet successfully deleted!",
	})
}
