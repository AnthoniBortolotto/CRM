package routes

import (
	"crm-go/modules/auth/v1/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(router *gin.RouterGroup) {
	handler := handlers.NewAuthHandler()

	auth := router.Group("/auth/v1")
	{
		auth.POST("/login", handler.Login)
		auth.POST("/register", handler.Register)
	}
}
