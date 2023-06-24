package exercises

import (
	"net/http"
	"time"

	"github.com/arthuruan/training-consultancy/common/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UpdateExerciseBody struct {
	Name     string `json:"name" validate:"required"`
	Category string `json:"category" validate:"required"`
	VideoURL string `json:"videoUrl" validate:"required"`
}

func (h handler) UpdateExercise(ctx *gin.Context) {
	id := ctx.Param("id")
	objId, _ := primitive.ObjectIDFromHex(id)

	body := UpdateExerciseBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errorMessage": err.Error(),
		})
		return
	}

	// Validate body
	var validate = validator.New()
	if err := validate.Struct(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errorMessage": err.Error(),
		})
		return
	}

	// Update exercise
	update := bson.M{
		"name":      body.Name,
		"category":  body.Category,
		"videoURL":  body.VideoURL,
		"updatedAt": time.Now(),
	}
	result, err := h.exercisesCollection.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": update})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"errorMessage": "Failed to update exercise.",
		})
		return
	}

	var updatedExercise models.Exercise
	if result.MatchedCount == 1 {
		if err := h.exercisesCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&updatedExercise); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"errorMessage": err.Error(),
			})
			return
		}
	}

	ctx.JSON(http.StatusOK, updatedExercise)
}
