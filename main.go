package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"sample/src/controller"
)

func main() {
	//HttpRepo := &repos.HttpRepo{}

	//loads values from .env into the system
	router := gin.Default()

	router.GET("/catalogue", controller.GetCatalogue)
	router.POST("/order", controller.PlaceOrder)
	router.PUT("/order/:id/status", controller.UpdateOrderStatus)

	log.Fatal(router.Run(":8080"))
}
