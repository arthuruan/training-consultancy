package exercises

import (
	"github.com/arthuruan/training-consultancy/common/db"
	"github.com/arthuruan/training-consultancy/common/middleware"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type handler struct {
	exercisesCollection *mongo.Collection
}

func RegisterRoutes(router *gin.Engine, client *mongo.Client) {
	exercisesCollection := db.GetCollection(client, "exercises")

	h := &handler{
		exercisesCollection,
	}

	v1 := router.Group("/v1")

	exercises := v1.Group("/exercises")
	exercises.POST("/", middleware.RequireAuth, h.AddExercie)
	exercises.GET("/", middleware.RequireAuth, h.GetExercises)
	exercises.GET("/:id", middleware.RequireAuth, h.GetExerciseById)
	exercises.PUT("/:id", middleware.RequireAuth, h.UpdateExercise)
}
