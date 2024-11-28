package repositories

import (
	"url-shortener/domain/entities"
)

// URLShortenerRepository is...
//
//go:generate mockgen -destination=mocks/mock_us_repo.go -package=mocks url-shortener/domain/repositories URLShortenerRepository
type URLShortenerRepository interface {
	GetAll() map[string]entities.URL
	Save(input entities.URL) error
	GetByLongURL(url string) (*entities.URL, error)
	GetByShortURL(url string) (*entities.URL, error)
	IncrementAccessCount(shortURL string) error
}
