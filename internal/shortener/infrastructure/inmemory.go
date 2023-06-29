package infrastructure

import (
	"errors"
	"sync"
)

type inMemoryRepository struct {
	sync.RWMutex
	urls map[string]string
}

func NewInMemoryRepository() Repository {
	return &inMemoryRepository{
		urls: make(map[string]string),
	}
}

func (r *inMemoryRepository) Save(url string, shortURL string) error {
	r.Lock()
	defer r.Unlock()

	if _, exists := r.urls[shortURL]; exists {
		return errors.New("short URL already exists")
	}

	r.urls[shortURL] = url
	return nil
}

func (r *inMemoryRepository) Load(shortURL string) (string, error) {
	r.RLock()
	defer r.RUnlock()

	if url, exists := r.urls[shortURL]; exists {
		return url, nil
	}

	return "", errors.New("short URL not found")
}
