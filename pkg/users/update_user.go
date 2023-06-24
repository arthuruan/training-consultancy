package users

import (
	"net/http"
	"time"

	"github.com/arthuruan/training-consultancy/common/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UpdateUserBody struct {
	Name string `json:"name" validate:"required"`
}

func (h handler) UpdateUser(ctx *gin.Context) {
	id := ctx.Param("id")
	objId, _ := primitive.ObjectIDFromHex(id)

	body := UpdateUserBody{}

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

	// Update user
	update := bson.M{"name": body.Name, "updatedAt": time.Now()}
	result, err := h.usersCollection.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": update})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"errorMessage": "Failed to update user.",
		})
		return
	}

	var updatedUser models.User
	if result.MatchedCount == 1 {
		if err := h.usersCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&updatedUser); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"errorMessage": err.Error(),
			})
			return
		}
	}

	ctx.JSON(http.StatusOK, updatedUser)
}
