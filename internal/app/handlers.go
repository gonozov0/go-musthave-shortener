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
func (h *RootHandler) createShortUrl(longUrl string) string {
	shortUrl := strconv.Itoa(int(h.urlCount))
	h.urlMap[shortUrl] = longUrl
	h.urlCount++
	return shortUrl
}
func (h *RootHandler) getLongUrl(shortUrl string) (string, error) {
	longUrl, ok := h.urlMap[shortUrl]
	if !ok {
		return "", fmt.Errorf("short url doesn't exist")
	}
	return longUrl, nil
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
		shortUrl := h.createShortUrl(string(b))
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(shortUrl))
	case http.MethodGet:
		shortUrl := strings.TrimSuffix(strings.TrimPrefix(r.URL.Path, "/"), "/")
		longUrl, err := h.getLongUrl(shortUrl)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		w.Header().Set("Location", longUrl)
		w.WriteHeader(http.StatusTemporaryRedirect)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
