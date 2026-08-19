package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"screenshot-search/internal/config"
)

func main() {
	cfg := config.Load()

	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	log.Printf("starting server on port %s", cfg.Port)
	if err := router.Run(":" + cfg.Port); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
