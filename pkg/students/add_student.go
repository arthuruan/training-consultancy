package students

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

type StudentBody struct {
	ProfessorID   string `json:"professorId" validate:"required"`
	PlanType      string `json:"planType" validate:"required"`
	Name          string `json:"name" validate:"required"`
	Email         string `json:"email" validate:"required"`
	Phone         string `json:"phone" validate:"required"`
	BirthDate     string `json:"birthDate" validate:"required"`
	Gender        string `json:"gender" validate:"required"`
	Objective     string `json:"objective" validate:"required"`
	Frequence     string `json:"frequence" validate:"required"`
	TrainingPlace string `json:"trainingPlace" validate:"required"`
}

func (h handler) AddStudent(ctx *gin.Context) {
	body := StudentBody{}

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

	// Validate ProfessorID
	var professor models.User
	professorId, _ := primitive.ObjectIDFromHex(body.ProfessorID)
	if err := h.usersCollection.FindOne(ctx, bson.M{"_id": professorId, "type": models.UserType.Professor}).Decode(&professor); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"errorMessage": "The professorId was not found.",
		})
		return
	}

	// Create a unique index on the email field
	if _, err := h.studentsCollection.Indexes().CreateOne(
		ctx,
		mongo.IndexModel{
			Keys:    bson.D{{Key: "email", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
	); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"errorMessage": "Email is already registered.",
		})
		return
	}

	birthDate, err := time.Parse(time.RFC3339, body.BirthDate)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errorMessage": "Invalid birthDate format, you should use ISO 8601 format.",
		})
		return
	}

	// Create Student
	student := models.Student{
		ID:            primitive.NewObjectID(),
		ProfessorID:   body.ProfessorID,
		Name:          body.Name,
		Email:         body.Email,
		Phone:         body.Phone,
		BirthDate:     birthDate,
		Objective:     body.Objective,
		Gender:        body.Gender,
		PlanType:      body.PlanType,
		Frequence:     body.Frequence,
		TrainingPlace: body.TrainingPlace,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	if _, err := h.studentsCollection.InsertOne(ctx, student); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errorMessage": "Failed to register the student.",
		})
		return
	}

	ctx.JSON(http.StatusCreated, student)
}
