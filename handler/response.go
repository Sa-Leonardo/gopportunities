package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)


func sendError(ctx *gin.Context, code int, msg string) {

	ctx.Header("Content-type", "aplication/json")
	ctx.JSON(code, gin.H{
		"messege": msg,
		"errorCode": code,
	})

}



func sendSuccess(ctx *gin.Context, op string, data interface{}) {

	ctx.Header("Content-type", "aplication/json")
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Operation from handler: %s sucessfull", op),
		"data": data,
	})
}