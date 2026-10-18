package main

import (
	"log"
	"os"

	"github.com/anuj-jha-work/golang-ecommerce-core/controllers"
	"github.com/anuj-jha-work/golang-ecommerce-core/database"
	"github.com/anuj-jha-work/golang-ecommerce-core/middleware"
	"github.com/anuj-jha-work/golang-ecommerce-core/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	app := controllers.NewApplication(database.ProductData(database.Client, "Products"), database.UserData(database.Client, "Users"))

	router := gin.New()
	router.Use(gin.Logger())

	routes.UserRoutes(router)
	router.Use(middleware.Authentication())

	router.GET("/addtocart", controllers.AddToCart())
	router.GET("/removeitem", controllers.RemoveItem())
	router.GET("/checkout", controllers.BuyFromCart())
	router.GET("/instantbuy", controllers.InstantBuy())

	log.Fatal(router.Run(":" + port))

}
