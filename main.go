package main

import (
	"github.com/arthuruan/training-consultancy/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()

	// routes
	routes.UserRoute(app)

	app.Run()
}
