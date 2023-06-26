package main

import (
	"github.com/arthuruan/training-consultancy/common/configs"
	"github.com/arthuruan/training-consultancy/common/db"
	"github.com/arthuruan/training-consultancy/pkg/exercises"
	"github.com/arthuruan/training-consultancy/pkg/users"
	workoutsheets "github.com/arthuruan/training-consultancy/pkg/workout_sheets"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	app := gin.Default()
	configs.LoadEnv()

	// DB Client
	var dbClient *mongo.Client = db.ConnectDB()

	// routes
	users.RegisterRoutes(app, dbClient)
	workoutsheets.RegisterRoutes(app, dbClient)
	exercises.RegisterRoutes(app, dbClient)

	app.Run()
}
