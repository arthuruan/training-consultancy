package students

import (
	"net/http"
	"time"

	"github.com/arthuruan/training-consultancy/common/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UpdateStudentBody struct {
	ProfessorID   string `json:"professorId" validate:"required"`
	Name          string `json:"name" validate:"required"`
	BirthDate     string `json:"birthDate" validate:"required"`
	Objective     string `json:"objective" validate:"required"`
	Gender        string `json:"gender" validate:"required"`
	Phone         string `json:"phone" validate:"required"`
	PlanType      string `json:"planType" validate:"required"`
	Frequence     string `json:"frequence" validate:"required"`
	TrainingPlace string `json:"trainingPlace" validate:"required"`
}

func (h handler) UpdateStudent(ctx *gin.Context) {
	id := ctx.Param("id")
	objId, _ := primitive.ObjectIDFromHex(id)

	body := UpdateStudentBody{}

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

	birthDate, err := time.Parse(time.RFC3339, body.BirthDate)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errorMessage": "Invalid birthDate format, you should use ISO 8601 format.",
		})
		return
	}

	// Update student
	update := bson.M{
		"professorId":   body.ProfessorID,
		"name":          body.Name,
		"birthDate":     birthDate,
		"objective":     body.Objective,
		"gender":        body.Gender,
		"phone":         body.Phone,
		"planType":      body.PlanType,
		"frequence":     body.Frequence,
		"trainingPlace": body.TrainingPlace,
		"updatedAt":     time.Now(),
	}
	result, err := h.studentsCollection.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": update})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"errorMessage": "Failed to update student.",
		})
		return
	}

	var updatedStudent models.Student
	if result.MatchedCount == 1 {
		if err := h.studentsCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&updatedStudent); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"errorMessage": err.Error(),
			})
			return
		}
	}

	ctx.JSON(http.StatusOK, updatedStudent)
}
