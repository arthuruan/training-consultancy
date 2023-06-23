package students

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h handler) AddStudent(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{})
}
