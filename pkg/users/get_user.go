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

	if err := h.userCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&user); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"errorMessage": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, user)
}
