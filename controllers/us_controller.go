package controllers

import (
	"fmt"
	"net/http"
	"url-shortener/app/controller"
	"url-shortener/domain/entities"
	"url-shortener/domain/services"

	"github.com/gin-gonic/gin"
)

// URLShortenerController is...
type URLShortenerController interface {
	GetAllShortURLs(gCtx *gin.Context)
	ShortenURL(gCtx *gin.Context)
	ResolveURL(gCtx *gin.Context)
}

// NewURLShortenerController creates a new instance of the urlShortenerController
func NewURLShortenerController(urlShortenerService services.URLShortenerService) URLShortenerController {
	return &urlShortenerController{urlShortenerService: urlShortenerService}
}

type urlShortenerController struct {
	controller.Controller
	urlShortenerService services.URLShortenerService
}

func (c *urlShortenerController) GetAllShortURLs(gCtx *gin.Context) {
	// Get all short URLs from the service
	urls := c.urlShortenerService.GetAll()

	if len(urls) == 0 {
		gCtx.JSON(http.StatusOK, gin.H{"message": "No URLs found"})
		return
	}

	gCtx.JSON(http.StatusOK, gin.H{"urls": urls})
}

// AddToCart adds an item in cart based on inputs and returns updated order or error
func (c *urlShortenerController) ShortenURL(gCtx *gin.Context) {
	var input entities.CreateShortURLInput

	if err := gCtx.ShouldBind(&input); err != nil {
		gCtx.JSON(http.StatusBadRequest, gin.H{"code": "failed", "msg": err.Error()})
		return
	}

	err := input.Validate()
	if err != nil {
		gCtx.JSON(http.StatusBadRequest, gin.H{"code": "failed", "msg": err.Error()})
		return
	}

	url, err := c.urlShortenerService.Shorten(input)
	if err != nil {
		c.SendWithError(gCtx, err)
		return
	}

	// Dynamically get the current host
	host := gCtx.Request.Host
	protocol := "http"           // Default to HTTP; set to HTTPS if required
	if gCtx.Request.TLS != nil { // If the request is over HTTPS
		protocol = "https"
	}

	// Construct the full short URL
	shortURL := fmt.Sprintf("%s://%s/%s", protocol, host, url)

	c.Send(gCtx, gin.H{"short_url": shortURL})
}

// UpdateOrderStatus updates an order's status based on inputs and returns updated order or error
func (c *urlShortenerController) ResolveURL(gCtx *gin.Context) {
	shortURL := gCtx.Param("short_url")

	url, err := c.urlShortenerService.Resolve(shortURL)
	if err != nil {
		c.SendWithError(gCtx, err)
		return
	}

	gCtx.Redirect(http.StatusMovedPermanently, url)
}
