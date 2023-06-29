package domain

import (
	"github.com/gonozov0/go-musthave-shortener/internal/shortener/infrastructure"
)

func Resolve(shortURL string, repository infrastructure.Repository) (string, error) {
	return repository.Load(shortURL)
}
