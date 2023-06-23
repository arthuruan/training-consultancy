package exercises

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h handler) AddExercie(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{})
}
