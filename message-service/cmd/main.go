package main

import (
	"log"
	"message-service/config"
	"message-service/internal/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()
	r := gin.Default()

	handler.RegisterRoutes(r, cfg)

	if err := r.Run(cfg.ServerAddress); err != nil {
		log.Fatalf("could not run server: %v", err)
	}
}
