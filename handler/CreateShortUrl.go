package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateShortUrl(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}