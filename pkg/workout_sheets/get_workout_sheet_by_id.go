package workoutsheets

import (
	"net/http"

	"github.com/arthuruan/training-consultancy/common/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (h handler) GetWorkoutSheetById(ctx *gin.Context) {
	id := ctx.Param("id")
	objId, _ := primitive.ObjectIDFromHex(id)

	var workoutSheet models.WorkoutSheet

	if err := h.workoutSheetsCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&workoutSheet); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"errorMessage": "Workout sheet not found.",
		})
		return
	}

	ctx.JSON(http.StatusOK, workoutSheet)
}
