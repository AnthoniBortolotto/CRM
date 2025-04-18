package router

import "github.com/gin-gonic/gin"

func InitializeRoutes(router *gin.Engine) {
	router.GET("/health-check", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "OK"})
	})

	router.GET("/lead", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "lead"})
	})

	router.GET("/auth/login", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "login"})
	})

	router.GET("/auth/register", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "register"})
	})

}
