package controllers

import (
	"net/http"
	"url-shortener/app/controller"
	"url-shortener/domain/entities"
	"url-shortener/domain/services"

	"github.com/gin-gonic/gin"
)

// URLShortenerController is...
type URLShortenerController interface {
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

	url, err := c.urlShortenerService.Shorten(input.LongURL)
	if err != nil {
		c.SendWithError(gCtx, err)
		return
	}

	c.Send(gCtx, gin.H{"short_url": url})
}

// UpdateOrderStatus updates an order's status based on inputs and returns updated order or error
func (c *urlShortenerController) ResolveURL(gCtx *gin.Context) {
	shortURL := gCtx.Param("shortURL")

	url, err := c.urlShortenerService.Resolve(shortURL)
	if err != nil {
		c.SendWithError(gCtx, err)
		return
	}

	gCtx.Redirect(http.StatusMovedPermanently, url)
}
