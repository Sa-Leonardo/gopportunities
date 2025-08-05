package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowOpeningHandlerOpeningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg" : "Show opening",
	})
}