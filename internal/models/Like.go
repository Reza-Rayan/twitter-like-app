package models

import "time"

type Like struct {
	ID        uint `gorm:"primaryKey"`
	UserID    uint `gorm:"uniqueIndex:idx_user_post"`
	PostID    uint `gorm:"uniqueIndex:idx_user_post"`
	CreatedAt time.Time
}
