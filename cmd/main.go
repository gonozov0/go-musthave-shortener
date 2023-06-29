package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gonozov0/go-musthave-shortener/config"
	"github.com/gonozov0/go-musthave-shortener/internal/shortener"
	"log"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	router := gin.Default()
	shortener.SetupRoutes(router, cfg.App.Host)

	err = router.Run(fmt.Sprintf(":%s", cfg.App.Port))
	if err != nil {
		log.Fatalf("failed to run router: %v", err)
	}
}
