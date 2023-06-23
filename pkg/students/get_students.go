package students

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h handler) GetStudents(ctx *gin.Context) {
	// should accept a query
	ctx.JSON(http.StatusOK, gin.H{})
}
