package students

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h handler) DeleteStudent(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{})
}
