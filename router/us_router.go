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

	// a POST request to /order will create an order
	r.POST("/shorten", func(c *gin.Context) {
		urlShortenerController.ShortenURL(c)
	})

	// a GET request to /order/:id will fetch an order
	r.GET("/:shortURL", func(c *gin.Context) {
		urlShortenerController.ResolveURL(c)
	})
}
