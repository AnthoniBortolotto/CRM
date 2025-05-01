package routes

import (
	"crm-go/modules/lead/v1/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterLeadRoutes(router *gin.RouterGroup) {
	handler := handlers.NewLeadHandler()

	lead := router.Group("/lead/v1")
	{
		lead.GET("", handler.GetLead)
		lead.POST("", handler.CreateLead)
	}
}
