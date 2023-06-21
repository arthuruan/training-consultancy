package user

import (
	"net/http"

	"github.com/arthuruan/training-consultancy/common/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserBody struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
	Type     string `json:"type" validate:"required"`
	Name     string `json:"name" validate:"required"`
}

type ErrorResponse struct {
	ErrorMessage string `json:"errorMessage"`
}

var validate = validator.New()

func (h handler) AddUser(ctx *gin.Context) {
	body := UserBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// use the validator library to validate required fields
	if validationErr := validate.Struct(&body); validationErr != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse{
			ErrorMessage: validationErr.Error(),
		})
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
