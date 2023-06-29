package shortener

import (
	"github.com/gin-gonic/gin"
	"github.com/gonozov0/go-musthave-shortener/internal/shortener/application"
	"github.com/gonozov0/go-musthave-shortener/internal/shortener/infrastructure"
)

func SetupRoutes(router *gin.Engine, host string) {
	shortener := infrastructure.NewInMemoryRepository()

	router.POST("/shorten", application.ShortenHandler(shortener, host))
	router.GET("/:shortURL", application.ResolveHandler(shortener))
}
