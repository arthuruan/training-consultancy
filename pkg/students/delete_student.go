package students

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (h handler) DeleteStudent(ctx *gin.Context) {
	studentId := ctx.Param("id")
	objId, _ := primitive.ObjectIDFromHex(studentId)

	result, err := h.studentsCollection.DeleteOne(ctx, bson.M{"_id": objId})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"errorMessage": "Failed to delete student.",
		})
		return
	}

	fmt.Println(result)

	if result.DeletedCount == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"errorMessage": "StudentId not found.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Student successfully deleted!",
	})
}
