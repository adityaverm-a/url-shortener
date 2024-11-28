package router

import (
	"net/http"
	urlShortenerRouter "url-shortener/router"

	"github.com/gin-gonic/gin"
)

// Create, this function sets up a new instance of a Gin router
func Create() http.Handler {
	r := gin.New()

	// Server Health Check Endpoint
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	urlShortenerRouter.InjectURLShortenerRoutes(r)

	return r
}
