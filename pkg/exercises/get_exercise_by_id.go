package exercises

import (
	"net/http"

	"github.com/arthuruan/training-consultancy/common/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (h handler) GetExerciseById(ctx *gin.Context) {
	id := ctx.Param("id")
	objId, _ := primitive.ObjectIDFromHex(id)

	var exercise models.Exercise

	if err := h.exercisesCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&exercise); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"errorMessage": "exercise not found.",
		})
		return
	}

	ctx.JSON(http.StatusOK, exercise)
}
