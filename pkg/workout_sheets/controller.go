package workoutsheets

import (
	"github.com/arthuruan/training-consultancy/common/db"
	"github.com/arthuruan/training-consultancy/common/middleware"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type handler struct {
	workoutSheetsCollection *mongo.Collection
	studentsCollection      *mongo.Collection
}

func RegisterRoutes(router *gin.Engine, client *mongo.Client) {
	workoutSheetsCollection := db.GetCollection(client, "workout_sheets")
	studentsCollection := db.GetCollection(client, "students")

	h := &handler{
		workoutSheetsCollection,
		studentsCollection,
	}

	v1 := router.Group("/v1")

	students := v1.Group("/workout-sheets")
	students.POST("/", middleware.RequireAuth, h.AddWorkoutSheet)
	students.GET("/:id", middleware.RequireAuth, h.GetWorkoutSheetById)
}
