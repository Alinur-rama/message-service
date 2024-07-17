package service

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"message-service/config"
	"message-service/internal/repository"

	"github.com/gin-gonic/gin"
	"github.com/segmentio/kafka-go"
)

type MessageService struct {
	repo        *repository.MessageRepository
	kafkaWriter *kafka.Writer
}

func NewMessageService(cfg *config.Config) *MessageService {
	fmt.Println(cfg.PostgresURL)
	repo := repository.NewMessageRepository(cfg.PostgresURL)
	kafkaWriter := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{cfg.KafkaBroker},
		Topic:   cfg.KafkaTopic,
	})

	return &MessageService{repo: repo, kafkaWriter: kafkaWriter}
}

func (s *MessageService) SaveMessage(c *gin.Context) {
	var msg repository.Message
	if err := c.BindJSON(&msg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	msg.Status = "processed"
	if err := s.repo.SaveMessage(&msg); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	messageBytes, err := json.Marshal(msg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = s.kafkaWriter.WriteMessages(context.Background(), kafka.Message{
		Value: messageBytes,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "message saved"})
}

func (s *MessageService) GetStatistics(c *gin.Context) {
	stats, err := s.repo.GetMessageStatistics()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, stats)
}
