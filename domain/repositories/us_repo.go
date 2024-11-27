package repositories

import (
	"url-shortener/domain/entities"
)

// UrlShortenerRepository is...
//
//go:generate mockgen -destination=mocks/mock_us_repo.go -package=mocks url-shortener/domain/repositories UrlShortenerRepository
type UrlShortenerRepository interface {
	Save(input entities.URL) error
	GetByLongURL(url string) (*entities.URL, error)
	GetByShortURL(url string) (*entities.URL, error)
	IncrementAccessCount(shortURL string) error
}
