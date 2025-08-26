package service

import (
	"github.com/Reza-Rayan/twitter-like-app/internal/notification"
	"github.com/Reza-Rayan/twitter-like-app/internal/notification/repository"
)

type NotifyService interface {
	Save(notification *notification.Notification) error
	GetUserNotifications(userID int64) ([]notification.Notification, error)
	MarkAsRead(notificationID int64) error
}

type notifyService struct {
	repo repository.NotifyRepository
}

func NewNotificationService(repo repository.NotifyRepository) NotifyService {
	return &notifyService{repo: repo}
}

func (s *notifyService) Save(notification *notification.Notification) error {
	return s.repo.Save(notification)
}

func (s *notifyService) GetUserNotifications(userID int64) ([]notification.Notification, error) {
	return s.repo.GetUserNotifications(userID)
}

func (s *notifyService) MarkAsRead(notificationID int64) error {
	return s.repo.MarkAsRead(notificationID)
}
