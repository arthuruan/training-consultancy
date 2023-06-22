package users

import (
	"net/http"

	"github.com/arthuruan/training-consultancy/common/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func (h handler) GetUsers(ctx *gin.Context) {
	var users []models.User

	cursor, err := h.userCollection.Find(ctx, bson.D{})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	defer cursor.Close(ctx)
	if err := cursor.All(ctx, &users); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
	}

	ctx.JSON(http.StatusOK, users)
}
