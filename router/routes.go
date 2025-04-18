package router

import (
	auth_routes "crm-go/modules/auth/routes"
	lead_routes "crm-go/modules/lead/routes"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine) {
	// Health check route
	router.GET("/health-check", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "OK"})
	})

	// Register route groups
	api := router.Group("/api")
	{
		auth_routes.RegisterAuthRoutes(api)
		lead_routes.RegisterLeadRoutes(api)
	}
}
