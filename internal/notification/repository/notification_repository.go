package repository

import (
	"github.com/Reza-Rayan/twitter-like-app/db"
	"github.com/Reza-Rayan/twitter-like-app/internal/notification"
)

type NotifyRepository interface {
	Save(notification *notification.Notification) error
	GetUserNotifications(userID int64) ([]notification.Notification, error)
	MarkAsRead(notificationID int64) error
}

type notifyRepo struct{}

func NewNotificationRepository() NotifyRepository {
	return &notifyRepo{}
}

func (r *notifyRepo) Save(notification *notification.Notification) error {
	return db.DB.Create(notification).Error
}

func (r *notifyRepo) GetUserNotifications(userID int64) ([]notification.Notification, error) {
	var notifs []notification.Notification
	err := db.DB.Where("recipient_id = ?", userID).Order("created_at desc").Find(&notifs).Error
	return notifs, err
}

func (r *notifyRepo) MarkAsRead(notificationID int64) error {
	return db.DB.Model(&notification.Notification{}).Where("id = ?", notificationID).Update("is_read", true).Error
}
