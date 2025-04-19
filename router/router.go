package router

import (
	"crm-go/config"

	"github.com/gin-gonic/gin"
)

func Initialize() {
	// Initialize Router
	router := gin.Default()

	// Initialize Routes
	InitializeRoutes(router)

	// Get the port from environment variables
	port := config.GetEnv("PORT", "8080")

	// Run the server
	router.Run("0.0.0.0:" + port)
}
