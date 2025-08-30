package models

import "time"

type Notification struct {
	ID          int64  `gorm:"primaryKey"`
	RecipientID int64  `gorm:"not null"`
	SenderID    int64  `gorm:"not null"`
	Type        string `gorm:"not null"`
	PostID      *int64
	Message     string
	IsRead      bool      `gorm:"default:false"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
}
