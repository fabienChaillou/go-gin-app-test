package main

import (
	"github.com/fabienChaillou/go-gin-app-test/internal/handler"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	// Set the router as the default one provided by Gin
	router = gin.Default()

	// Process the templates at the start so that they don't have to be loaded
	// from the disk again. This makes serving HTML pages very fast.
	router.LoadHTMLGlob("templates/*")

	// Initialize the routes
	handler.RegisterRoutes(router)

	// Start serving the application
	router.Run()
}
