package main

import (
	"github.com/gin-gonic/gin"
	"github.com/link33d/shorter/config"
	"github.com/link33d/shorter/router"
)

func main() {

	// Initialize the server
	server := gin.Default()

	// Initialize the database
	_, err := config.ConnectDB()
	if err != nil {
		panic(err)
	}

	// Initialize Router
	router.Initialize(server)

	// Run the server
	server.Run(":8080")
}
