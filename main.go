package main

import (
	"log"
	"log/slog"
	"os"

	"github.com/amit8889/ecommerce-golang/controllers"
	"github.com/amit8889/ecommerce-golang/database"
	"github.com/amit8889/ecommerce-golang/middleware"
	"github.com/amit8889/ecommerce-golang/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	slog.Info("====main function start execution==>")
	port := os.Getenv("PORT")
	if port == "" {
		slog.Info("PORT environment variable is not set")
		port = "8000"
	}

	app := controllers.NewApplication(
		database.ProductData(database.Client, "Products"),
		database.UserData(database.Client, "Users"),
	)
	router := gin.New()
	router.Use(gin.Logger())
	routers.UserRouters(router)

	router.Use(middleware.Authenticate())

	router.GET("/addcart", app.AddToCart())
	router.GET("/removecart", app.RemoveFromCart())

	log.Fatal(router.Run((":" + port)))

}
