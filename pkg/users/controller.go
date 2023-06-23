package users

import (
	"github.com/arthuruan/training-consultancy/common/db"
	"github.com/arthuruan/training-consultancy/common/middleware"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type handler struct {
	userCollection *mongo.Collection
}

func RegisterRoutes(router *gin.Engine, client *mongo.Client) {
	var userCollection *mongo.Collection = db.GetCollection(client, "users")

	h := &handler{
		userCollection,
	}

	v1 := router.Group("/v1")
	users := v1.Group("/users")
	users.POST("/", h.AddUser)
	users.POST("/login", h.Login)
	users.GET("/", middleware.RequireAuth, h.GetUsers)
	users.GET("/:id", middleware.RequireAuth, h.GetUser)
}
