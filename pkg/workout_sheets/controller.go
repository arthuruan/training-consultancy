package workoutsheets

import (
	"github.com/arthuruan/training-consultancy/common/db"
	"github.com/arthuruan/training-consultancy/common/middleware"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type handler struct {
	workoutSheetsCollection *mongo.Collection
	usersCollection         *mongo.Collection
}

func RegisterRoutes(router *gin.Engine, client *mongo.Client) {
	workoutSheetsCollection := db.GetCollection(client, "workout_sheets")
	usersCollection := db.GetCollection(client, "users")

	h := &handler{
		workoutSheetsCollection,
		usersCollection,
	}

	v1 := router.Group("/v1")

	workoutSheets := v1.Group("/workout-sheets")
	workoutSheets.POST("/", middleware.RequireAuth, h.AddWorkoutSheet)
	workoutSheets.GET("/:id", middleware.RequireAuth, h.GetWorkoutSheetById)
}
