package workoutsheets

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h handler) AddWorkoutSheet(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{})
}
