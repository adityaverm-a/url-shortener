package repositories

import (
	"sync"
	"url-shortener/domain/entities"
	"url-shortener/domain/repositories"
)

// memoryRepo implements the URLShortenerRepository interface using sync.Map for thread-safe storage.
type memoryRepo struct {
	data sync.Map // Thread-safe map for storing URLs
}

// NewMemoryRepo creates and returns a new instance of memoryRepo.
func NewMemoryRepo() repositories.URLShortenerRepository {
	return &memoryRepo{}
}

// GetAll returns all stored URLs from the memory repository.
func (m *memoryRepo) GetAll() map[string]entities.URL {
	result := make(map[string]entities.URL)

	m.data.Range(func(key, value interface{}) bool {
		if url, ok := value.(entities.URL); ok {
			result[key.(string)] = url
		}
		return true // Continue iteration
	})

	return result
}

// Save stores a URL entity in the memory repository.
func (m *memoryRepo) Save(url entities.URL) error {
	m.data.Store(url.ShortURL, url)
	return nil
}

// GetByShortURL retrieves a URL entity by its short URL.
func (m *memoryRepo) GetByShortURL(shortURL string) (*entities.URL, error) {
	if value, ok := m.data.Load(shortURL); ok {
		url := value.(entities.URL)
		return &url, nil
	}
	return nil, nil
}

// GetByLongURL retrieves a URL entity by its long URL.
func (m *memoryRepo) GetByLongURL(longURL string) (*entities.URL, error) {
	var result *entities.URL

	m.data.Range(func(key, value interface{}) bool {
		url := value.(entities.URL)
		if url.LongURL == longURL {
			result = &url
			return false
		}
		return true
	})
	return result, nil
}

// IncrementAccessCount increases the access count for the given short URL.
func (m *memoryRepo) IncrementAccessCount(shortURL string) error {
	if value, ok := m.data.Load(shortURL); ok {
		url := value.(entities.URL)

		url.AccessCount++

		m.data.Store(shortURL, url)

		return nil
	}

	return nil
}
