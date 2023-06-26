package workoutsheets

import (
	"net/http"

	"github.com/arthuruan/training-consultancy/common/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (h handler) GetWorkoutSheets(ctx *gin.Context) {
	var workoutSheets []models.WorkoutSheet

	filters := []primitive.E{}
	if studentId := ctx.Query("studentId"); studentId != "" {
		objId, _ := primitive.ObjectIDFromHex(studentId)
		filters = append(filters, primitive.E{Key: "studentId", Value: objId})
	}

	cursor, err := h.workoutSheetsCollection.Find(ctx, filters)
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
