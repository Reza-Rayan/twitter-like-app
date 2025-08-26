package models

import "time"

type Notification struct {
	ID          uint   `gorm:"primaryKey"`
	RecipientID uint   `gorm:"not null"`
	SenderID    uint   `gorm:"not null"`
	Type        string `gorm:"not null"`
	PostID      *uint
	Message     string
	IsRead      bool      `gorm:"default:false"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
}
