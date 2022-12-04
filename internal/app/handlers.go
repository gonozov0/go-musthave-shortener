package app

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

type RootHandler struct {
	urlMap   map[string]string
	urlCount uint
}

func NewRootHandler() RootHandler {
	return RootHandler{
		urlMap:   map[string]string{},
		urlCount: 0,
	}
}
func (h *RootHandler) createShortURL(longURL string) string {
	shortURL := strconv.Itoa(int(h.urlCount))
	h.urlMap[shortURL] = longURL
	h.urlCount++
	return shortURL
}
func (h *RootHandler) getLongURL(shortURL string) (string, error) {
	longURL, ok := h.urlMap[shortURL]
	if !ok {
		return "", fmt.Errorf("short url doesn't exist")
	}
	return longURL, nil
}
func (h *RootHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		b, err := io.ReadAll(r.Body)
		// обрабатываем ошибку
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		shortURL := h.createShortURL(string(b))
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(shortURL))
	case http.MethodGet:
		shortURL := strings.TrimSuffix(strings.TrimPrefix(r.URL.Path, "/"), "/")
		longURL, err := h.getLongURL(shortURL)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		w.Header().Set("Location", longURL)
		w.WriteHeader(http.StatusTemporaryRedirect)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
