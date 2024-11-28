package router

import (
	"url-shortener/containers"
	"url-shortener/controllers"

	"github.com/gin-gonic/gin"
)

var urlShortenerController controllers.URLShortenerController

// InjectURLShortenerRoutes sets up the necessary routes for the URL shortener service
// and injects the required dependencies to the controller.
func InjectURLShortenerRoutes(router *gin.Engine) {
	setupRoutes(router)
}

// setupRoutes configures the individual API endpoints for the URL shortener service.
func setupRoutes(r *gin.Engine) {

	container := containers.NewContainer()
	urlShortenerController = container.InjectURLShortenerController()

	// Route to resolve a short URL and redirect to the original long URL
	// GET /:short_url
	r.GET("/:short_url", func(c *gin.Context) {
		urlShortenerController.ResolveURL(c)
	})

	v1 := r.Group("/v1")

	// Route to fetch all short URLs and their associated data
	// GET /v1/urls
	v1.GET("/urls", func(c *gin.Context) {
		urlShortenerController.GetAllShortURLs(c)
	})

	// Route to shorten a long URL (optionally with a custom short URL)
	// POST /v1/shorten
	v1.POST("/shorten", func(c *gin.Context) {
		urlShortenerController.ShortenURL(c)
	})

}
