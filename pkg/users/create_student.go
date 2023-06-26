package users

import (
	"net/http"
	"time"

	"github.com/arthuruan/training-consultancy/common/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CreateStudentBody struct {
	PersonalID    string `json:"personalId" validate:"required"`
	PlanType      string `json:"planType" validate:"required"`
	Name          string `json:"name" validate:"required"`
	Email         string `json:"email" validate:"required"`
	Phone         string `json:"phone" validate:"required"`
	Birthday      string `json:"birthday" validate:"required"`
	Gender        string `json:"gender" validate:"required"`
	Objective     string `json:"objective" validate:"required"`
	Frequence     string `json:"frequence" validate:"required"`
	TrainingPlace string `json:"trainingPlace" validate:"required"`
}

func (h handler) CreateStudent(ctx *gin.Context) {
	body := CreateStudentBody{}

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

	// Validate PersonalID
	var personal models.User
	personalId, _ := primitive.ObjectIDFromHex(body.PersonalID)
	if err := h.usersCollection.FindOne(ctx, bson.M{"_id": personalId, "type": models.UserType.Personal}).Decode(&personal); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"errorMessage": "PersonalId was not found.",
		})
		return
	}

	// Insert in the database
	student := models.Student{
		ID:            primitive.NewObjectID(),
		Name:          body.Name,
		Email:         body.Email,
		Type:          models.UserType.Student,
		PersonalID:    personalId,
		Birthday:      birthday,
		Objective:     body.Objective,
		Gender:        body.Gender,
		Phone:         body.Phone,
		PlanType:      body.PlanType,
		Frequence:     body.Frequence,
		TrainingPlace: body.TrainingPlace,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}
	if _, err := h.usersCollection.InsertOne(ctx, student); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errorMessage": "Failed to register the student.",
		})
		return
	}

	ctx.JSON(http.StatusCreated, student)
}
