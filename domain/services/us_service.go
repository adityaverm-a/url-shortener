package services

import (
	"errors"
	"time"
	"url-shortener/domain/entities"
	"url-shortener/domain/repositories"

	"golang.org/x/exp/rand"
)

// URLShortenerService is...
type URLShortenerService interface {
	Shorten(longURL string) (string, error)
	Resolve(shortURL string) (string, error)
}

// The NewURLShortenerService function is a factory function that returns a new instance of the urlShortenerService
func NewURLShortenerService(repo repositories.URLShortenerRepository) URLShortenerService {
	return &urlShortenerService{repo: repo}
}

type urlShortenerService struct {
	repo repositories.URLShortenerRepository
}

// The GetOrderByID method of the urlShortenerService struct utilizes the OrderRepository instance, by calling the GetByID method on it, and returns any order and errors if received.
func (service *urlShortenerService) Shorten(longURL string) (string, error) {
	existing, _ := service.repo.GetByLongURL(longURL)
	if existing != nil {
		return existing.ShortURL, nil
	}

	shortURL := service.generateShortURL(6)

	url := entities.URL{
		LongURL:  longURL,
		ShortURL: shortURL,
	}

	err := service.repo.Save(url)

	return shortURL, err
}

func (service *urlShortenerService) generateShortURL(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	// Seed the random number generator to ensure randomness
	rand.Seed(uint64(time.Now().UnixNano()))

	// Create a slice of runes to store the random string
	result := make([]byte, length)

	// Generate the random string
	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}

	return string(result)
}

// The CreateOrder method of the urlShortenerService struct utilizes the OrderRepository instance, by calling the Create method on it, and returns the order created and errors if received.
func (service *urlShortenerService) Resolve(shortURL string) (string, error) {
	url, err := service.repo.GetByShortURL(shortURL)
	if err != nil || url == nil {
		return "", errors.New("short URL not found")
	}

	_ = service.repo.IncrementAccessCount(shortURL)

	return url.LongURL, nil
}
