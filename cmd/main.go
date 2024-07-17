package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"message-service/config"
	"message-service/internal/handler"
)

func main() {
	cfg := config.LoadConfig()
	r := gin.Default()

	handler.RegisterRoutes(r, cfg)

	if err := r.Run(cfg.ServerAddress); err != nil {
		log.Fatalf("could not run server: %v", err)
	}
}
