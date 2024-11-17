package handler

import "github.com/gin-gonic/gin"

func sendMessage(ctx *gin.Context, code int, msg string) {

	ctx.Header("Content-type", "application/json")

	ctx.JSON(code, gin.H{
		"message": msg,
	})
}

func sendData(ctx *gin.Context, code int, msg string, data interface{}) {

	ctx.Header("Content-type", "application/json")

	ctx.JSON(code, gin.H{
		"message": msg,
		"data":    data,
	})
}
