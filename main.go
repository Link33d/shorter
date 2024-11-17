package main

import (
	"github.com/gin-gonic/gin"
	"github.com/link33d/shorter/router"
)

func main() {

	// Initialize the server
	server := gin.Default()

	// Initialize Router
	router.Initialize(server)

	// Run the server
	server.Run(":8080")
}
