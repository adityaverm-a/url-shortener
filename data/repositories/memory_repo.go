package repositories

import (
	"sync"
	"url-shortener/domain/entities"
	"url-shortener/domain/repositories"
)

type memoryRepo struct {
	data map[string]entities.URL
	mu   sync.RWMutex
}

func NewMemoryRepo() repositories.URLShortenerRepository {
	return &memoryRepo{
		data: make(map[string]entities.URL),
	}
}

func (m *memoryRepo) GetAll() map[string]entities.URL {
	return m.data
}

func (m *memoryRepo) Save(url entities.URL) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.data[url.ShortURL] = url
	return nil
}

func (m *memoryRepo) GetByShortURL(shortURL string) (*entities.URL, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	if url, exists := m.data[shortURL]; exists {
		return &url, nil
	}
	return nil, nil
}

func (m *memoryRepo) GetByLongURL(longURL string) (*entities.URL, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	for _, url := range m.data {
		if url.LongURL == longURL {
			return &url, nil
		}
	}
	return nil, nil
}

func (m *memoryRepo) IncrementAccessCount(shortURL string) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if url, exists := m.data[shortURL]; exists {
		url.AccessCount++
		m.data[shortURL] = url
		return nil
	}
	return nil
}
