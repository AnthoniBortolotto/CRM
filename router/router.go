package router

import (
	"os"

	"github.com/gin-gonic/gin"
)

func Initialize() {
	// Initialize Router
	router := gin.Default()

	// Initialize Routes
	InitializeRoutes(router)

	// Get the port from the environment
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Run the server
	router.Run("localhost:" + port)
}
