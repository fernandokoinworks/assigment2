package main

import (
	"ass2/controller/order_controller"
	"ass2/db"
	"fmt"

	"github.com/gin-gonic/gin"
)

func init() {
	fmt.Println("init function")
	db.InitializeDB()

}

func main() {
	route := gin.Default()
	orderRoute := route.Group("/orders")
	{
		orderRoute.POST("/", order_controller.CreateOrder)
		orderRoute.GET("/", order_controller.GetOrder)
		orderRoute.PUT("/:id", order_controller.UpdateOrder)
		orderRoute.DELETE("/:id", order_controller.DeleteOrder)

	}
	route.Run(":8080")

}
