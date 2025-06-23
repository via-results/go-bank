package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the Gin router
	// This is a simple web server that responds with a welcome message
	// when accessed at the root URL ("/").
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.String(200, "Welcome to Go Customer!")
	})
	router.Run()
}
