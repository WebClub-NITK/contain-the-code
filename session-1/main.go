package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

const (
	listenIP    = "0.0.0.0:"
	defaultPort = "8080"
	portKey     = "PORT"
)

func main() {
	s := gin.Default()

	// Get port from PORT envariable. Use default is not set
	port := os.Getenv(portKey)
	if port == "" {
		port = defaultPort
	}
	// Setup routes
	s.GET("ping", pingHandler)
	s.NoRoute(defaultHandler)

	// Start listening
	s.Run(listenIP + port)
}

func pingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "pong",
	})
	return
}

func defaultHandler(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
		"status": http.StatusNotFound,
	})
}
