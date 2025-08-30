package models

import "time"

type Like struct {
	ID        int64 `gorm:"primaryKey"`
	UserID    int64 `gorm:"uniqueIndex:idx_user_post"`
	PostID    int64 `gorm:"uniqueIndex:idx_user_post"`
	CreatedAt time.Time
}
