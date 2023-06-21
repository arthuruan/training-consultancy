package user

import (
	"github.com/arthuruan/training-consultancy/common/db"
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

	routes := router.Group("/users")
	routes.POST("/", h.AddUser)
}
