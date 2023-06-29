package application

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gonozov0/go-musthave-shortener/internal/shortener/domain"
	"github.com/gonozov0/go-musthave-shortener/internal/shortener/infrastructure"
)

type ShortenRequest struct {
	URL string `json:"url"`
}

func ShortenHandler(repository infrastructure.Repository, host string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req ShortenRequest
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		shortURL, err := domain.Shorten(req.URL, repository, host)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"short_url": shortURL})
	}
}

func ResolveHandler(repository infrastructure.Repository) gin.HandlerFunc {
	return func(c *gin.Context) {
		shortURL := c.Param("shortURL")
		url, err := domain.Resolve(shortURL, repository)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		c.Redirect(http.StatusMovedPermanently, url)
	}
}
