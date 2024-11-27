package router

import (
	"url-shortener/containers"
	"url-shortener/controllers"

	"github.com/gin-gonic/gin"
)

var urlShortenerController controllers.UrlShortenerController

// InjectUrlShortenerRoutes is defined to set up the routes and inject the dependencies for the router to work correctly.
func InjectUrlShortenerRoutes(router *gin.RouterGroup) {
	setupRoutes(router)
}

// setupRoutes is defined to set up the router's endpoints.
func setupRoutes(r *gin.RouterGroup) {

	// create a new instance of the container
	container := containers.NewContainer()
	urlShortenerController = container.InjectUrlShortenerController()

	// // a GET request to /orders will fetch  all orders
	// r.GET("/orders", func(c *gin.Context) {
	// 	orderController.GetOrders(c)
	// })

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// // a POST request to /order will create an order
	// r.POST("/order", func(c *gin.Context) {
	// 	orderController.CreateOrder(c)
	// })

	// // a GET request to /order/:id will fetch an order
	// r.GET("/order/:id", func(c *gin.Context) {
	// 	orderController.GetOrder(c)
	// })

	// // a POST request to /order/change-status will update status for that order
	// r.POST("/order/change-status", func(c *gin.Context) {
	// 	orderController.UpdateOrderStatus(c)
	// })
}
