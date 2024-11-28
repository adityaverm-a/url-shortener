package router

import (
	"url-shortener/containers"
	"url-shortener/controllers"

	"github.com/gin-gonic/gin"
)

var urlShortenerController controllers.URLShortenerController

// InjectURLShortenerRoutes is defined to set up the routes and inject the dependencies for the router to work correctly.
func InjectURLShortenerRoutes(router *gin.Engine) {
	setupRoutes(router)
}

// setupRoutes is defined to set up the router's endpoints.
func setupRoutes(r *gin.Engine) {

	// create a new instance of the container
	container := containers.NewContainer()
	urlShortenerController = container.InjectURLShortenerController()

	// a GET request to /order/:id will fetch an order
	r.GET("/:short_url", func(c *gin.Context) {
		urlShortenerController.ResolveURL(c)
	})

	v1 := r.Group("/v1")

	v1.GET("/urls", func(c *gin.Context) {
		urlShortenerController.GetAllShortURLs(c)
	})

	// a POST request to /order will create an order
	v1.POST("/shorten", func(c *gin.Context) {
		urlShortenerController.ShortenURL(c)
	})

}
