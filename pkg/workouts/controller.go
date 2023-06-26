package workouts

import (
	"github.com/arthuruan/training-consultancy/common/db"
	"github.com/arthuruan/training-consultancy/common/middleware"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type handler struct {
	workoutsCollection  *mongo.Collection
	exercisesCollection *mongo.Collection
}

func RegisterRoutes(router *gin.Engine, client *mongo.Client) {
	workoutsCollection := db.GetCollection(client, "workouts")
	exercisesCollection := db.GetCollection(client, "exercises")

	h := &handler{
		workoutsCollection,
		exercisesCollection,
	}

	v1 := router.Group("/v1")

	workoutSheets := v1.Group("/workouts")
	workoutSheets.POST("/", middleware.RequireAuth, h.CreateWorkouts)
	workoutSheets.GET("/", middleware.RequireAuth, h.GetWorkouts)
	workoutSheets.GET("/:id", middleware.RequireAuth, h.GetWorkoutById)
	workoutSheets.PUT("/:id", middleware.RequireAuth, h.UpdateWorkout)
	workoutSheets.DELETE("/workout-sheet/:workoutSheetId", middleware.RequireAuth, h.DeleteWorkoutsByWorkoutSheetId)
}
