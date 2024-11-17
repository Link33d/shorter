package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func sendMessage(ctx *gin.Context, code int, msg string) {

	ctx.Header("Content-type", "application/json")

	ctx.JSON(code, gin.H{
		"message": msg,
	})
}

func sendInternalServerError(ctx *gin.Context) {
	sendMessage(ctx, http.StatusInternalServerError, "Something went wrong when communicating with the database")
}

func sendData(ctx *gin.Context, code int, msg string, data interface{}) {

	ctx.Header("Content-type", "application/json")

	ctx.JSON(code, gin.H{
		"message": msg,
		"data":    data,
	})
}
