package repositories

import (
	"sync"
	"url-shortener/domain/entities"
	"url-shortener/domain/repositories"
)

type memoryRepo struct {
	data sync.Map
}

func NewMemoryRepo() repositories.URLShortenerRepository {
	return &memoryRepo{}
}

// GetAll returns all stored URLs from the map
func (m *memoryRepo) GetAll() map[string]entities.URL {
	result := make(map[string]entities.URL)
	m.data.Range(func(key, value interface{}) bool {
		if url, ok := value.(entities.URL); ok {
			result[key.(string)] = url
		}
		return true
	})
	return result
}

func (m *memoryRepo) Save(url entities.URL) error {
	m.data.Store(url.ShortURL, url)
	return nil
}

func (m *memoryRepo) GetByShortURL(shortURL string) (*entities.URL, error) {
	if value, ok := m.data.Load(shortURL); ok {
		url := value.(entities.URL)
		return &url, nil
	}
	return nil, nil
}

func (m *memoryRepo) GetByLongURL(longURL string) (*entities.URL, error) {
	var result *entities.URL
	m.data.Range(func(key, value interface{}) bool {
		url := value.(entities.URL)
		if url.LongURL == longURL {
			result = &url
			return false // Stop iteration once the match is found
		}
		return true
	})
	return result, nil
}

func (m *memoryRepo) IncrementAccessCount(shortURL string) error {
	if value, ok := m.data.Load(shortURL); ok {
		url := value.(entities.URL)
		url.AccessCount++
		m.data.Store(shortURL, url)
		return nil
	}
	return nil
}
