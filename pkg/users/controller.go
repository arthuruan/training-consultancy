package users

import (
	"github.com/arthuruan/training-consultancy/common/db"
	"github.com/arthuruan/training-consultancy/common/middleware"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type handler struct {
	usersCollection *mongo.Collection
}

func RegisterRoutes(router *gin.Engine, client *mongo.Client) {
	var usersCollection *mongo.Collection = db.GetCollection(client, "users")

	h := &handler{
		usersCollection,
	}

	v1 := router.Group("/v1")

	users := v1.Group("/users")
	users.POST("/personal", h.CreatePersonal)
	users.POST("/student", h.CreateStudent)
	users.POST("/login", h.Login)
	users.GET("/", middleware.RequireAuth, h.GetUsers)
	users.GET("/:id", middleware.RequireAuth, h.GetUserById)
	users.PUT("/:id", middleware.RequireAuth, h.UpdateUser)
	users.DELETE("/:id", middleware.RequireAuth, h.DeleteUser)
}
