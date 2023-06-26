package programs

import (
	"github.com/arthuruan/training-consultancy/common/db"
	"github.com/arthuruan/training-consultancy/common/middleware"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type handler struct {
	workoutSheetsCollection *mongo.Collection
	workoutsCollection      *mongo.Collection
	exercisesCollection     *mongo.Collection
	usersCollection         *mongo.Collection
}

func RegisterRoutes(router *gin.Engine, client *mongo.Client) {
	workoutSheetsCollection := db.GetCollection(client, "workout_sheets")
	workoutsCollection := db.GetCollection(client, "workouts")
	exercisesCollection := db.GetCollection(client, "exercises")
	usersCollection := db.GetCollection(client, "users")

	h := &handler{
		workoutSheetsCollection,
		workoutsCollection,
		exercisesCollection,
		usersCollection,
	}

	v1 := router.Group("/v1")

	workoutSheets := v1.Group("/programs")
	workoutSheets.POST("/", middleware.RequireAuth, h.CreateProgram)
	workoutSheets.GET("/sheet/:workoutSheetId", middleware.RequireAuth, h.GetProgramBySheetId)
	workoutSheets.DELETE("/sheet/:workoutSheetId", middleware.RequireAuth, h.DeleteProgram)
}
