package routes

import (
	"crm-go/modules/lead/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterLeadRoutes(router *gin.RouterGroup) {
	handler := handlers.NewLeadHandler()

	lead := router.Group("/lead")
	{
		lead.GET("", handler.GetLead)
	}
}
