package handlers

import "github.com/gin-gonic/gin"

type LeadHandler struct {
	// Repository will be added here later
}

func NewLeadHandler() *LeadHandler {
	return &LeadHandler{}
}

func (h *LeadHandler) GetLead(c *gin.Context) {
	c.JSON(200, gin.H{"message": "lead"})
}
