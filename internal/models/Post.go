package models

import "time"

type Post struct {
	ID        uint   `gorm:"primaryKey"`
	Title     string `gorm:"not null"`
	Content   string `gorm:"not null"`
	CreatedAt time.Time
	UserID    uint
	User      User
	Image     *string
	Likes     []Like
}

type PostWithLikes struct {
	ID         int64     `json:"id"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	CreatedAt  time.Time `json:"created_at"`
	UserID     int64     `json:"user_id"`
	Image      *string   `json:"image,omitempty"`
	LikesCount int       `json:"likes_count"`
}
