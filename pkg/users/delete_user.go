package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (h handler) DeleteUser(ctx *gin.Context) {
	userId := ctx.Param("id")
	objId, _ := primitive.ObjectIDFromHex(userId)

	result, err := h.usersCollection.DeleteOne(ctx, bson.M{"_id": objId})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"errorMessage": "Failed to delete user.",
		})
		return
	}

	if result.DeletedCount == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"errorMessage": "UserId not found.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "User successfully deleted!",
	})
}
