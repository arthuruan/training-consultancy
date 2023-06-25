package users

import (
	"net/http"

	"github.com/arthuruan/training-consultancy/common/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (h handler) GetUsers(ctx *gin.Context) {
	var users []models.User

	filters := []primitive.E{}
	if userType := ctx.Query("type"); userType != "" {
		filters = append(filters, primitive.E{"type", userType})
	}
	if personalId := ctx.Query("personalId"); personalId != "" {
		filters = append(filters, primitive.E{"personalId", personalId})
	}

	cursor, err := h.usersCollection.Find(ctx, filters)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"errorMessage": err.Error(),
		})
		return
	}

	defer cursor.Close(ctx)
	if err := cursor.All(ctx, &users); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errorMessage": "Failed to list users.",
		})
		return
	}

	ctx.JSON(http.StatusOK, users)
}
