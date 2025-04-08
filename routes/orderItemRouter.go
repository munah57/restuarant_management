package routes

import (
	"github.com/gin-gonic/gin"
	controller "go-resturant-management/controllers"
)

func OrderItemRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/orderItems", controller.GetOrderItems())
	incomingRoutes.GET("/orderItems/:order_item", controller.GetOrderItem())
	incomingRoutes.GET("/orderItems-order/:order_id",controller.GetOrderItemsByOrder())
	incomingRoutes.POST("/orderItems", controller.CreateOrderItem())
	incomingRoutes.PATCH("/orderItems/:order_item", controller.UpdateOrderItem())
}
