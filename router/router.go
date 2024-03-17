package router

import (
	"RestAPI_GO/controller"

	"github.com/gin-gonic/gin"
)

func Router() {
	router := gin.Default()
	router.POST("/orders", controller.Create)

	router.GET("/orders", controller.List)

	router.PUT("/orders/:orderId", controller.Update)

	router.DELETE("/orders/:orderId", controller.Delete)

	if err := router.Run(); err != nil {
		panic(err)
	}
}