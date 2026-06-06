package main

import (
	"net/http"

	"WebhookFlow/internal/config"
	"WebhookFlow/internal/database"

	"github.com/gin-gonic/gin"
)

func main() {

	cfg := config.Load()

	db := database.NewPostgresPool(cfg)
	defer db.Close()

	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	router.Run(":" + cfg.Port)
}
