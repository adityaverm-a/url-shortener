package containers

import (
	"url-shortener/controllers"
	"url-shortener/data/repositories"
	"url-shortener/domain/services"
)

// Container is...
type Container interface {
	InjectUrlShortenerController() controllers.UrlShortenerController
}

// NewContainer creates a new Container instance
func NewContainer() Container {
	return &container{}
}

type container struct{}

// InjectUrlShortenerController injects an instance of the UrlShortenerController
func (c *container) InjectUrlShortenerController() controllers.UrlShortenerController {
	urlShortenerService, err := InjectUrlShortenerService()
	if err != nil {
		panic(1)
	}

	return controllers.NewUrlShortenerController(urlShortenerService)
}

// InjectUrlShortenerService injects an instance of the UrlShortenerService
func InjectUrlShortenerService() (services.UrlShortenerService, error) {

	urlShortenerRepository := repositories.NewUrlShortenerRepository()

	urlShortenerService := services.NewUrlShortenerService(urlShortenerRepository)

	return urlShortenerService, nil
}
