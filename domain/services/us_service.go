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
	GetAll() map[string]entities.URL
	Shorten(input entities.CreateShortURLInput) (string, error)
	Resolve(shortURL string) (string, error)
}

// The NewURLShortenerService function is a factory function that returns a new instance of the urlShortenerService
func NewURLShortenerService(repo repositories.URLShortenerRepository) URLShortenerService {
	return &urlShortenerService{repo: repo}
}

type urlShortenerService struct {
	repo repositories.URLShortenerRepository
}

func (service *urlShortenerService) GetAll() map[string]entities.URL {
	return service.repo.GetAll()
}

// The GetOrderByID method of the urlShortenerService struct utilizes the OrderRepository instance, by calling the GetByID method on it, and returns any order and errors if received.
func (service *urlShortenerService) Shorten(input entities.CreateShortURLInput) (string, error) {
	if input.CustomShortURL != "" {
		existing, _ := service.repo.GetByShortURL(input.CustomShortURL)
		if existing != nil {
			return "", errors.New("custom short URL already exists")
		}

		url := entities.URL{
			LongURL:  input.LongURL,
			ShortURL: input.CustomShortURL,
		}

		err := service.repo.Save(url)

		return input.CustomShortURL, err
	}

	existing, _ := service.repo.GetByLongURL(input.LongURL)
	if existing != nil {
		return existing.ShortURL, nil
	}

	shortURL := service.generateShortURL(6)

	url := entities.URL{
		LongURL:  input.LongURL,
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
	url, _ := service.repo.GetByShortURL(shortURL)
	if url == nil {
		return "", errors.New("The requested short URL does not exist.")
	}

	_ = service.repo.IncrementAccessCount(shortURL)

	return url.LongURL, nil
}
