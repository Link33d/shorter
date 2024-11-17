package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RedirectShortUrl(ctx *gin.Context) {
	code := ctx.Param("code")

	ctx.JSON(http.StatusOK, gin.H{
		"message": code,
	})
}
