package exercises

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (h handler) DeleteExercise(ctx *gin.Context) {
	exerciseId := ctx.Param("id")
	objId, _ := primitive.ObjectIDFromHex(exerciseId)

	result, err := h.exercisesCollection.DeleteOne(ctx, bson.M{"_id": objId})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"errorMessage": "Failed to delete exercise.",
		})
		return
	}

	if result.DeletedCount == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"errorMessage": "ExerciseId not found.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Exercise successfully deleted!",
	})
}
