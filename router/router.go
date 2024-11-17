package router

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Initialize(server *gin.Engine) {

	// Index page router
	server.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// Create link router
	server.POST("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	// Redirect router
	server.GET("/:code", func(ctx *gin.Context) {

		code := ctx.Param("code")

		fmt.Println(code)

		ctx.JSON(http.StatusOK, gin.H{
			"message": code,
		})
	})
}
