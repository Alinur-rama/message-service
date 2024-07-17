package handler

import (
	"message-service/config"
	"message-service/internal/service"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, cfg *config.Config) {
	messageService := service.NewMessageService(cfg)
	r.POST("/messages", messageService.SaveMessage)
	r.GET("/statistics", messageService.GetStatistics)
}
