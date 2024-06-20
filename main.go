package main

import (
	"os"
	"restaurant-management/database"
	"restaurant-management/middleware"
	"restaurant-management/routes"
	"restaurant-management/controller"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())
	routes.UserRoutes(router)
	router.Use(middleware.Authentication())

	// Routes is a package, FoodRoutes is a function

	routes.FoodRoutes(router)
	routes.MenuRoutes(router)
	routes.TableRoutes(router)
	routes.OrderRoutes(router)
	routes.OrderItemRoutes(router)
	routes.InvoiceRoutes(router)

	// all the function are called here, simply import package "routes" for every file

	router.Run(":" + port)
}
