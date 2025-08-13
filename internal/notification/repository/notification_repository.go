package repository

import (
	"database/sql"
	"github.com/Reza-Rayan/twitter-like-app/db"
	"github.com/Reza-Rayan/twitter-like-app/internal/notification"
)

type NotifyRepository interface {
	Save(notification *notification.Notification) error
	GetUserNotifications(userID int64) ([]notification.Notification, error)
}

type notificationRepo struct {
	db *sql.DB
}

func NewNotificationRepository(db *sql.DB) NotifyRepository {
	return &notificationRepo{db: db}
}

// Save -> POST method
func (r *notificationRepo) Save(notification *notification.Notification) error {
	query := `
		INSERT INTO notifications (recipient_id, sender_id, type, post_id, message)
		VALUES (?, ?, ?, ?, ?)
	`
	_, err := r.db.Exec(query, notification.RecipientID, notification.SenderID, notification.Type, notification.PostID, notification.Message)
	return err
}

// GetUserNotifications -> GET method
func (r *notificationRepo) GetUserNotifications(userID int64) ([]notification.Notification, error) {
	query := `
		SELECT id, recipient_id, sender_id, type, post_id, message, is_read, created_at
		FROM notifications WHERE recipient_id = ? ORDER BY created_at DESC
	`
	rows, err := db.DB.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notifs []notification.Notification
	for rows.Next() {
		var n notification.Notification
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
