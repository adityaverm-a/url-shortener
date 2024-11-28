package containers

import (
	"url-shortener/config"
	"url-shortener/controllers"
	"url-shortener/data/repositories"
	"url-shortener/domain/services"
)

// Container is...
type Container interface {
	InjectURLShortenerController() controllers.URLShortenerController
}

// NewContainer creates a new Container instance
func NewContainer() Container {
	return &container{}
}

type container struct{}

// InjectURLShortenerController injects an instance of the URLShortenerController
func (c *container) InjectURLShortenerController() controllers.URLShortenerController {
	urlShortenerService, err := InjectURLShortenerService()
	if err != nil {
		panic(1)
	}

	return controllers.NewURLShortenerController(urlShortenerService)
}

// InjectURLShortenerService injects an instance of the URLShortenerService
func InjectURLShortenerService() (services.URLShortenerService, error) {

	urlShortenerRepository := repositories.NewMemoryRepo()

	urlShortenerService := services.NewURLShortenerService(urlShortenerRepository, config.Config.Charset, config.Config.ShortURLLength)

	return urlShortenerService, nil
}
