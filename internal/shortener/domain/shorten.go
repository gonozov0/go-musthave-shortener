package domain

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"github.com/gonozov0/go-musthave-shortener/internal/shortener/infrastructure"
)

func Shorten(url string, repository infrastructure.Repository, host string) (string, error) {
	hash := sha1.Sum([]byte(url))
	shortURL := base64.URLEncoding.EncodeToString(hash[:])[:8]

	if err := repository.Save(url, shortURL); err != nil {
		return "", err
	}

	return fmt.Sprintf("%s/%s", host, shortURL), nil
}
