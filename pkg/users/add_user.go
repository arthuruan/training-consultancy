package users

import (
	"net/http"

	"github.com/arthuruan/training-consultancy/common/models"
	"github.com/arthuruan/training-consultancy/common/validations"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserBody struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
	Type     string `json:"type" validate:"required"`
	Name     string `json:"name" validate:"required"`
}

func (h handler) AddUser(ctx *gin.Context) {
	body := UserBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if validationErrors := validations.Struct(&body); validationErrors != nil {
		ctx.JSON(http.StatusBadRequest, validationErrors)
		return
	}

	user := models.User{
		ID:       primitive.NewObjectID(),
		Name:     body.Name,
		Email:    body.Email,
		Password: body.Password,
		Type:     body.Type,
	}

	if _, err := h.userCollection.InsertOne(ctx, user); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusCreated, user)
}
