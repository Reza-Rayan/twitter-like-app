package notification

import "time"

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
