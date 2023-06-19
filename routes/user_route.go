package routes

import (
	"github.com/arthuruan/training-consultancy/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoute(app *gin.Engine) {
	// All routes related to users comes here
	app.POST("/user", controllers.CreateUser)
}
