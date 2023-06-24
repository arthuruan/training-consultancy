package students

import (
	"net/http"

	"github.com/arthuruan/training-consultancy/common/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (h handler) GetStudent(ctx *gin.Context) {
	id := ctx.Param("id")
	objId, _ := primitive.ObjectIDFromHex(id)

	var student models.Student

	if err := h.studentsCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&student); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"errorMessage": "Student not found.",
		})
		return
	}

	ctx.JSON(http.StatusOK, student)
}
