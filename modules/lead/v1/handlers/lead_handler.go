package handlers

import (
	"crm-go/modules/auth/v1/utils"
	"crm-go/modules/lead/v1/models"
	"crm-go/modules/lead/v1/repositories"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LeadHandler struct {
	// Repository will be added here later
	leadRepo *repositories.LeadRepository
}

func NewLeadHandler() *LeadHandler {
	return &LeadHandler{
		leadRepo: repositories.NewLeadRepository(),
	}
}

func (h *LeadHandler) GetLead(c *gin.Context) {
	c.JSON(200, gin.H{"message": "lead"})
}

func (h *LeadHandler) CreateLead(c *gin.Context) {
	var req models.CreateLeadRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	println("Creating lead with request:")
	// Validate authentication token
	decodedToken, err := utils.ValidateToken(req.AuthenticationToken)
	println("Decoded token:", decodedToken)
	if err != nil {
		c.JSON(401, gin.H{"error": "Invalid token"})
		return
	}

	lead := &models.Lead{
		Email:       req.Email,
		Phone:       req.Phone,
		ClientName:  req.ClientName,
		ProductName: req.ProductName,
		LeadSource:  req.LeadSource,
		LeadStatus:  "New",
		Discount:    req.Discount,
		Notes:       req.Notes,
		LeadOwnerID: decodedToken.UserID,
		ID:          primitive.NewObjectID(),
		TotalPrice:  req.TotalPrice,
	}
	println("Lead to be created:", lead)
	// Insert lead into the database
	generatedLead, err := h.leadRepo.CreateLead(lead)
	println("Generated lead:")
	if err != nil {
		println("Error creating lead:", err)
		c.JSON(500, gin.H{"error": "Error creating lead"})
		return
	}
	println("Lead created successfully:")
	c.JSON(200, gin.H{"message": "lead created successfully", "lead": generatedLead})
}
