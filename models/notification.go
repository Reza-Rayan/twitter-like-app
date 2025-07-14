package models

import (
	"database/sql"
	"github.com/Reza-Rayan/twitter-like-app/db"
	"time"
)

type Notification struct {
	ID          int64     `json:"id"`
	RecipientID int64     `json:"recipient_id"`
	SenderID    int64     `json:"sender_id"`
	Type        string    `json:"type"`
	PostID      *int64    `json:"post_id,omitempty"`
	Message     string    `json:"message"`
	IsRead      bool      `json:"is_read"`
	CreatedAt   time.Time `json:"created_at"`
}

func (n *Notification) Save() error {
	query := `
		INSERT INTO notifications (recipient_id, sender_id, type, post_id, message)
		VALUES (?, ?, ?, ?, ?)
	`
	_, err := db.DB.Exec(query, n.RecipientID, n.SenderID, n.Type, n.PostID, n.Message)
	return err
}

func GetUserNotifications(userID int64) ([]Notification, error) {
	query := `
		SELECT id, recipient_id, sender_id, type, post_id, message, is_read, created_at
		FROM notifications WHERE recipient_id = ? ORDER BY created_at DESC
	`
	rows, err := db.DB.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notifs []Notification
	for rows.Next() {
		var n Notification
		var postID sql.NullInt64
		err := rows.Scan(&n.ID, &n.RecipientID, &n.SenderID, &n.Type, &postID, &n.Message, &n.IsRead, &n.CreatedAt)
		if err != nil {
			return nil, err
		}
		if postID.Valid {
			n.PostID = &postID.Int64
		}
		notifs = append(notifs, n)
	}
	return notifs, nil
}
