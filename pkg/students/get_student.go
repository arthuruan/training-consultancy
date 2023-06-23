package students

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h handler) GetStudent(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{})
}
