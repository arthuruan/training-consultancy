package exercises

import (
	"net/http"

	"github.com/arthuruan/training-consultancy/common/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ExerciseBody struct {
	Name     string `json:"name" validate:"required"`
	Category string `json:"category" validate:"required"`
	VideoURL string `json:"videoUrl" validate:"required"`
}

func (h handler) AddExercie(ctx *gin.Context) {
	body := ExerciseBody{}

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

	// Insert in the database
	exercise := models.Exercise{
		ID:       primitive.NewObjectID(),
		Name:     body.Name,
		Category: body.Category,
		VideoURL: body.VideoURL,
	}
	if _, err := h.exercisesCollection.InsertOne(ctx, exercise); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errorMessage": "Failed to create exercise.",
		})
		return
	}

	ctx.JSON(http.StatusCreated, exercise)
}
