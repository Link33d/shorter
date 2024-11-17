package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func IndexPage(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func CreateShortUrl(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

func RedirectShortUrl(ctx *gin.Context) {
	code := ctx.Param("code")

	ctx.JSON(http.StatusOK, gin.H{
		"message": code,
	})
}
