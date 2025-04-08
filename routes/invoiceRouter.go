package routes

import (
	controller "go-resturant-management/controllers"

	"github.com/gin-gonic/gin"
)

func InvoiceRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/invoices", controller.GetInvoices())
	incomingRoutes.GET("/invoices/id/:id", controller.GetInvoice())
	incomingRoutes.POST("/innvoices", controller.CreateInvoice())
	incomingRoutes.PATCH("invoices/:id", controller.UpdateInvoice())
}
