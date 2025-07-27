package models

import (
	"github.com/Reza-Rayan/twitter-like-app/db"
	"time"
)

type Message struct {
	ID         int64     `json:"id"`
	SenderID   int64     `json:"sender_id"`
	ReceiverID int64     `json:"receiver_id"`
	Content    string    `json:"content"`
	SentAt     time.Time `json:"sent_at"`
}

func (m *Message) Save() error {
	query := `INSERT INTO messages (sender_id, receiver_id, content, sent_at) VALUES (?, ?, ?, ?)`
	result, err := db.DB.Exec(query, m.SenderID, m.ReceiverID, m.Content, m.SentAt)
	if err != nil {
		return err
	}
	m.ID, _ = result.LastInsertId()
	return nil
}
