package main

import (
	"os"

	middleware "restaurant-management/middleware"
	routes "restaurant-management/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	router := gin.New()
	router.Use(gin.Logger())
	routes.AuthRoutes(router)
	router.Use(middleware.Authentification())

	routes.UserRoutes(router)

	routes.FoodRoutes(router)
	routes.InvoiceRoutes(router)
	routes.MenuRoutes(router)
	routes.OrderRoutes(router)
	routes.OrderItemsRoutes(router)
	routes.TableRoutes(router)

	router.Run(":" + port)
}
