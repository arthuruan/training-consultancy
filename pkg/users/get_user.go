package users

import (
	"net/http"

	"github.com/arthuruan/training-consultancy/common/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (h handler) GetUser(ctx *gin.Context) {
	id := ctx.Param("id")
	objId, _ := primitive.ObjectIDFromHex(id)

	var user models.User

	if err := h.usersCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&user); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"errorMessage": "User not found.",
		})
		return
	}

	ctx.JSON(http.StatusOK, user)
}
