package controllers

import (
	"fmt"
	"net/http"
	"url-shortener/app/controller"
	"url-shortener/domain/entities"
	"url-shortener/domain/services"

	"github.com/gin-gonic/gin"
)

// URLShortenerController is
type URLShortenerController interface {
	GetAllShortURLs(gCtx *gin.Context)
	ShortenURL(gCtx *gin.Context)
	ResolveURL(gCtx *gin.Context)
}

// NewURLShortenerController creates and returns a new URLShortenerController instance.
func NewURLShortenerController(urlShortenerService services.URLShortenerService) URLShortenerController {
	return &urlShortenerController{urlShortenerService: urlShortenerService}
}

type urlShortenerController struct {
	controller.Controller
	urlShortenerService services.URLShortenerService
}

// GetAllShortURLs retrieves all the stored short URLs and their associated long URLs.
func (c *urlShortenerController) GetAllShortURLs(gCtx *gin.Context) {
	urls := c.urlShortenerService.GetAll()

	if len(urls) == 0 {
		gCtx.JSON(http.StatusOK, gin.H{"message": "No URLs found"})
		return
	}

	gCtx.JSON(http.StatusOK, gin.H{"urls": urls})
}

// ShortenURL accepts a long URL and returns its corresponding short URL.
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

	// Dynamically get the current host and protocol (http/https)
	host := gCtx.Request.Host
	protocol := "http"
	if gCtx.Request.TLS != nil {
		protocol = "https"
	}

	shortURL := fmt.Sprintf("%s://%s/%s", protocol, host, url)

	// Send the response with the short URL
	c.Send(gCtx, gin.H{"short_url": shortURL})
}

// ResolveURL takes a short URL and redirects the user to the corresponding long URL.
func (c *urlShortenerController) ResolveURL(gCtx *gin.Context) {
	shortURL := gCtx.Param("short_url")

	url, err := c.urlShortenerService.Resolve(shortURL)
	if err != nil {
		c.SendWithError(gCtx, err)
		return
	}

	gCtx.Redirect(http.StatusMovedPermanently, url)
}
