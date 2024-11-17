package router

import (
	"github.com/gin-gonic/gin"
	"github.com/link33d/shorter/handler"
)

func Initialize(server *gin.Engine) {

	// Index page router
	server.GET("/", handler.IndexPage)

	// Create link router
	server.POST("/", handler.CreateShortUrl)

	// Redirect router
	server.GET("/:code", handler.RedirectShortUrl)
}
