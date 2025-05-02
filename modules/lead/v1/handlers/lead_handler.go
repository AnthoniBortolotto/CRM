package handlers

import (
	"crm-go/modules/auth/v1/utils"
	"crm-go/modules/lead/v1/models"
	"crm-go/modules/lead/v1/repositories"
	"net/http"

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
	// Get token from header
	leadId := c.Param("leadId")
	leadIdObjID, err := primitive.ObjectIDFromHex(leadId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid lead ID"})
		return
	}
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
		return
	}

	// Validate token
	decodedToken, err := utils.ValidateToken(token)
	println("Decoded token:", decodedToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	lead, err := h.leadRepo.GetLeadByID(leadIdObjID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting lead"})
		return
	}
	if lead == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Lead not found"})
		return
	}
	// Check if the lead belongs to the user
	if lead.LeadOwnerID != decodedToken.UserID {
		c.JSON(http.StatusForbidden, gin.H{"error": "You do not have permission to access this lead"})
		return
	}
	c.JSON(200, gin.H{"message": "lead found", "lead": lead})
}

func (h *LeadHandler) CreateLead(c *gin.Context) {
	var req models.CreateLeadRequest

	// Get token from header
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
		return
	}

	// Validate token
	decodedToken, err := utils.ValidateToken(token)
	println("Decoded token:", decodedToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	// Validate request body
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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

	// Insert lead into the database
	generatedLead, err := h.leadRepo.CreateLead(lead)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating lead"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "lead created successfully", "lead": generatedLead})
}
