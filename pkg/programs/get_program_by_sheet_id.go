package programs

import (
	"net/http"

	"github.com/arthuruan/training-consultancy/common/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (h handler) GetProgramBySheetId(ctx *gin.Context) {
	workoutSheetId := ctx.Param("workoutSheetId")
	objId, _ := primitive.ObjectIDFromHex(workoutSheetId)
	var sheet models.WorkoutSheet

	if err := h.workoutSheetsCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&sheet); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"errorMessage": "Workout sheet not found.",
		})
		return
	}

	var workouts []models.Workout
	cursor, err := h.workoutsCollection.Find(ctx, bson.D{{Key: "workoutSheetId", Value: objId}})
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

	program := models.Program{
		Sheet:    sheet,
		Workouts: workouts,
	}

	ctx.JSON(http.StatusOK, program)
}
