package users

import (
	"net/http"
	"time"

	"github.com/arthuruan/training-consultancy/common/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CreatePersonalBody struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
	Name     string `json:"name" validate:"required"`
	Phone    string `json:"phone" validate:"required"`
	Birthday string `json:"birthday" validate:"required"`
	Gender   string `json:"gender" validate:"required"`
}

func (h handler) CreatePersonal(ctx *gin.Context) {
	body := CreatePersonalBody{}

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

	// Create a unique index on the email field
	if _, err := h.usersCollection.Indexes().CreateOne(
		ctx,
		mongo.IndexModel{
			Keys:    bson.D{{Key: "email", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
	); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"errorMessage": err.Error(),
		})
		return
	}

	birthday, err := time.Parse(time.RFC3339, body.Birthday)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errorMessage": "Invalid birthday format, you should use ISO 8601 format.",
		})
		return
	}

	// Hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"errorMessage": "Failed to hash password",
		})
		return
	}

	// Insert in the database
	personal := models.Personal{
		ID:        primitive.NewObjectID(),
		Name:      body.Name,
		Email:     body.Email,
		Birthday:  birthday,
		Gender:    body.Gender,
		Phone:     body.Phone,
		Password:  string(hash),
		Type:      models.UserType.Personal,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if _, err := h.usersCollection.InsertOne(ctx, personal); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errorMessage": "Failed to register the personal.",
		})
		return
	}

	ctx.JSON(http.StatusCreated, personal)
}
