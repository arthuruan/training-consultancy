package main

import (
	"github.com/arthuruan/training-consultancy/common/db"
	"github.com/arthuruan/training-consultancy/pkg/user"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	app := gin.Default()

	// DB Client
	var dbClient *mongo.Client = db.ConnectDB()

	// routes
	user.RegisterRoutes(app, dbClient)

	app.Run()
}
