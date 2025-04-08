package main

import(
	"os"
	"go-resturant-management/database"
	"go-resturant-management/routes"
	"go-resturant-management/middleware"
	"go.mongodb.org/mongo-driver/mongo"
	"github.com/gin-gonic/gin"

)

var foodCollection *mongo.Collection = database.OpenCollection(database.Client,"food")

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080" // Default port
	}

	//initialize routes 
	router := gin.New()
	router.Use(gin.Logger())
	routes.UserRoutes(router)
	router.Use(middleware.Authentication()) 
	
	
	routes.FoodRoutes(router)
	routes.MenuRoutes(router)
	routes.OrderRoutes(router)
	routes.OrderItemRoutes(router)
	routes.TableRoutes(router)
	routes.InvoiceRoutes(router)

	router.Run(":" + port)
}