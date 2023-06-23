package students

import (
	"github.com/arthuruan/training-consultancy/common/db"
	"github.com/arthuruan/training-consultancy/common/middleware"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type handler struct {
	studentsCollection *mongo.Collection
}

func RegisterRoutes(router *gin.Engine, client *mongo.Client) {
	var studentsCollection *mongo.Collection = db.GetCollection(client, "students")

	h := &handler{
		studentsCollection,
	}

	v1 := router.Group("/v1")

	students := v1.Group("/students")
	students.POST("/", middleware.RequireAuth, h.AddStudent)
	students.GET("/:id", middleware.RequireAuth, h.GetStudent)
	students.GET("/", middleware.RequireAuth, h.GetStudents)
	students.PUT("/:id", middleware.RequireAuth, h.UpdateStudent)
	students.DELETE("/:id", middleware.RequireAuth, h.DeleteStudent)
}
