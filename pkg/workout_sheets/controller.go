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
	workoutsCollection      *mongo.Collection
	exercisesCollection     *mongo.Collection
}

func RegisterRoutes(router *gin.Engine, client *mongo.Client) {
	workoutSheetsCollection := db.GetCollection(client, "workout_sheets")
	usersCollection := db.GetCollection(client, "users")
	workoutsCollection := db.GetCollection(client, "workouts")
	exercisesCollection := db.GetCollection(client, "exercises")

	h := &handler{
		workoutSheetsCollection,
		usersCollection,
		workoutsCollection,
		exercisesCollection,
	}

	v1 := router.Group("/v1")

	workoutSheets := v1.Group("/workout-sheets")
	workoutSheets.POST("/", middleware.RequireAuth, h.CreateWorkoutSheet)
	workoutSheets.GET("/", middleware.RequireAuth, h.GetWorkoutSheets)
	workoutSheets.GET("/:id", middleware.RequireAuth, h.GetWorkoutSheetById)
	workoutSheets.PUT("/:id", middleware.RequireAuth, h.UpdateWorkoutSheet)
	workoutSheets.DELETE("/:id", middleware.RequireAuth, h.DeleteWorkoutSheet)
}
