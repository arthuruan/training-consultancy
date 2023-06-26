package workouts

import (
	"net/http"

	"github.com/arthuruan/training-consultancy/common/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (h handler) GetWorkoutById(ctx *gin.Context) {
	id := ctx.Param("id")
	objId, _ := primitive.ObjectIDFromHex(id)

	var workout models.Workout

	if err := h.workoutsCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&workout); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"errorMessage": "Workout not found.",
		})
		return
	}

	ctx.JSON(http.StatusOK, workout)
}
