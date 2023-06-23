package workouts

import (
	"github.com/arthuruan/training-consultancy/common/db"
	"github.com/arthuruan/training-consultancy/common/middleware"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type handler struct {
	workoutsCollection *mongo.Collection
}

func RegisterRoutes(router *gin.Engine, client *mongo.Client) {
	workoutsCollection := db.GetCollection(client, "workout")

	h := &handler{
		workoutsCollection,
	}

	v1 := router.Group("/v1")

	students := v1.Group("/workouts")
	students.POST("/", middleware.RequireAuth, h.AddWorkout)
}
