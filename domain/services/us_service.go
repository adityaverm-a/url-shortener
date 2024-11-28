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
func NewURLShortenerService(repo repositories.URLShortenerRepository, charset string, shortURLLength int) URLShortenerService {
	return &urlShortenerService{
		repo:           repo,
		charset:        charset,
		shortURLLength: shortURLLength,
	}
}

type urlShortenerService struct {
	repo           repositories.URLShortenerRepository
	charset        string
	shortURLLength int
}

func (service *urlShortenerService) GetAll() map[string]entities.URL {
	return service.repo.GetAll()
}

// The GetOrderByID method of the urlShortenerService struct utilizes the OrderRepository instance, by calling the GetByID method on it, and returns any order and errors if received.
func (service *urlShortenerService) Shorten(input entities.CreateShortURLInput) (string, error) {
	if input.CustomShortURL != "" {
		return service.handleCustomShortURL(input)
	}

	existing, _ := service.repo.GetByLongURL(input.LongURL)
	if existing != nil {
		return existing.ShortURL, nil
	}

	shortURL := service.generateShortURL()

	return service.createAndSaveShortURL(input.LongURL, shortURL, input.TTL)
}

func (service *urlShortenerService) handleCustomShortURL(input entities.CreateShortURLInput) (string, error) {
	existing, _ := service.repo.GetByShortURL(input.CustomShortURL)
	if existing != nil {
		return "", errors.New("custom short URL already exists")
	}

	return service.createAndSaveShortURL(input.LongURL, input.CustomShortURL, input.TTL)
}

func (service *urlShortenerService) createAndSaveShortURL(longURL string, shortURL string, ttl int64) (string, error) {
	url := entities.URL{
		LongURL:  longURL,
		ShortURL: shortURL,
	}

	if ttl > 0 {
		duration := time.Duration(ttl) * time.Second
		url.ExpiresAt = time.Now().Add(duration)
	}

	err := service.repo.Save(url)

	return url.ShortURL, err
}

func (service *urlShortenerService) generateShortURL() string {
	charset := DEFAULT_CHARSET
	if service.charset != "" {
		charset = service.charset
	}

	length := DEFAULT_SHORT_URL_LENGTH

	if service.shortURLLength > 0 {
		length = service.shortURLLength
	}

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
		return "", errors.New("the requested short URL does not exist")
	}

	if time.Now().After(url.ExpiresAt) {
		return "", errors.New("short URL has expired")
	}

	_ = service.repo.IncrementAccessCount(shortURL)

	return url.LongURL, nil
}
