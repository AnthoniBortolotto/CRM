package handlers

import "github.com/gin-gonic/gin"

type AuthHandler struct {
	// Repository will be added here later
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{}
}

func (h *AuthHandler) Login(c *gin.Context) {
	c.JSON(200, gin.H{"message": "login"})
}

func (h *AuthHandler) Register(c *gin.Context) {
	c.JSON(200, gin.H{"message": "register"})
}
