package repository

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Message struct {
	ID      int    `db:"id" json:"id"`
	Content string `db:"content" json:"content"`
	Status  string `db:"status" json:"status"`
}

type MessageStatistics struct {
	TotalMessages     int `db:"total_messages" json:"total_messages"`
	ProcessedMessages int `db:"processed_messages" json:"processed_messages"`
}

type MessageRepository struct {
	db *sqlx.DB
}

func NewMessageRepository(postgresURL string) *MessageRepository {
	db := sqlx.MustConnect("postgres", postgresURL)
	return &MessageRepository{db: db}
}

func (r *MessageRepository) SaveMessage(msg *Message) error {
	_, err := r.db.NamedExec(`INSERT INTO messages (content, status) VALUES (:content, :status)`, msg)
	return err
}

func (r *MessageRepository) GetMessageStatistics() (*MessageStatistics, error) {
	var stats MessageStatistics
	err := r.db.Get(&stats, `
		SELECT 
			COUNT(*) AS total_messages,
			COUNT(*) FILTER (WHERE status = 'processed') AS processed_messages
		FROM messages
	`)
	return &stats, err
}
