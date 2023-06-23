package workouts

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h handler) AddWorkout(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{})
}
