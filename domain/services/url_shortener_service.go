package services

import (
	"errors"
	"time"
	"url-shortener/domain/entities"
	"url-shortener/domain/repositories"

	"golang.org/x/exp/rand"
)

// URLShortenerService is
type URLShortenerService interface {
	GetAll() map[string]entities.URL
	Shorten(input entities.CreateShortURLInput) (string, error)
	Resolve(shortURL string) (string, error)
}

// NewURLShortenerService is a factory function that creates and returns a new instance of urlShortenerService.
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

// GetAll returns all URLs stored in the repository.
func (service *urlShortenerService) GetAll() map[string]entities.URL {
	return service.repo.GetAll()
}

// Shorten handles the URL shortening process. It either generates a new short URL or uses a provided custom short URL.
// If a custom short URL is provided, it checks for uniqueness. If not, a new short URL is generated and saved.
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

// handleCustomShortURL checks if the custom short URL is available. If it is, it saves the custom short URL.
func (service *urlShortenerService) handleCustomShortURL(input entities.CreateShortURLInput) (string, error) {
	existing, _ := service.repo.GetByShortURL(input.CustomShortURL)
	if existing != nil {
		return "", errors.New("custom short URL already exists")
	}

	return service.createAndSaveShortURL(input.LongURL, input.CustomShortURL, input.TTL)
}

// createAndSaveShortURL creates a new URL object with the short URL and TTL (if provided), then saves it in the repository.
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

// generateShortURL generates a random short URL using the charset and length provided during service initialization.
func (service *urlShortenerService) generateShortURL() string {
	charset := DEFAULT_CHARSET
	if service.charset != "" {
		charset = service.charset
	}

	length := DEFAULT_SHORT_URL_LENGTH
	if service.shortURLLength > 0 {
		length = service.shortURLLength
	}

	// Seed the random number generator for randomness
	rand.Seed(uint64(time.Now().UnixNano()))

	// Create a slice to store the random string
	result := make([]byte, length)

	// Generate the random short URL using the charset
	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}

	return string(result)
}

// Resolve looks up a short URL in the repository, validates its expiration, and returns the original long URL if valid.
func (service *urlShortenerService) Resolve(shortURL string) (string, error) {
	url, _ := service.repo.GetByShortURL(shortURL)
	if url == nil {
		return "", errors.New("the requested short URL does not exist")
	}

	if !url.ExpiresAt.IsZero() && time.Now().After(url.ExpiresAt) {
		return "", errors.New("short URL has expired")
	}

	_ = service.repo.IncrementAccessCount(shortURL)

	return url.LongURL, nil
}
